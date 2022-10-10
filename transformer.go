package plato

import "log"

/*
	第三步，遍历节点
	根据第二步中得到的 ast，我们要遍历 ast 来获得所有节点
*/

// visitor 用来遍历 ast 中的不同节点
type visitor map[string]func(n *node, parent node)

// traverser 接收 ast 和 visitor 两个参数，内部调用 traverseNode 函数对节点进行遍历
func traverser(a ast, vis visitor) {
	// 第一个 ast 节点是没有父节点的，所以第二个参数为空节点
	traverseNode(node(a), node{}, vis)
}

// traverseArray 遍历 ast 中的所有节点
// 内部调用 traverseNode 函数将当前节点和父节点传递给 visitor 进行操作
// 这里说的操作其实就是将匹配到的节点添加到指定节点列表
func traverseArray(a []node, parent node, vis visitor) {
	for _, child := range a {
		traverseNode(child, parent, vis)
	}
}

func traverseNode(n node, parent node, vis visitor) {
	// 遍历 visitor，获取节点类型 kind 和对应的函数
	for k, vf := range vis {
		// 如果类型匹配，则调用该函数对当前节点和其父节点进行操作
		// 操作意义同上
		if k == n.kind {
			vf(&n, parent)
		}
	}
	switch n.kind {
	// 如果节点类型是 Program，则调用 traverseArray 对节点列表进行遍历
	case "Program":
		traverseArray(n.body, n, vis)
		break
	// 如果节点类型是 CallExpression，则调用 traverseArray 对节点参数列表进行遍历
	case "CallExpression":
		traverseArray(n.params, n, vis)
		break
	// 如果节点类型是 NumberLiteral，表明这里就没有需要遍历的节点了，直接 break 就行
	case "NumberLiteral":
		break
	// 否则抛出错误
	default:
		log.Fatal(n.kind)
	}
}

/*
	第四步，实现转换器
	将 ast 和 visitor 传入我们的 traverser 函数，然后返回一个新的 ast

  ----------------------------------------------------------------------------
    原本的 AST                        |   转换之后的 AST
  ----------------------------------------------------------------------------
    {                                |   {
      type: 'Program',               |     type: 'Program',
      body: [{                       |     body: [{
        type: 'CallExpression',      |       type: 'ExpressionStatement',
        name: 'add',                 |       expression: {
        params: [{                   |         type: 'CallExpression',
          type: 'NumberLiteral',     |         callee: {
          value: '2'                 |           type: 'Identifier',
        }, {                         |           name: 'add'
          type: 'CallExpression',    |         },
          name: 'subtract',          |         arguments: [{
          params: [{                 |           type: 'NumberLiteral',
            type: 'NumberLiteral',   |           value: '2'
            value: '4'               |         }, {
          }, {                       |           type: 'CallExpression',
            type: 'NumberLiteral',   |           callee: {
            value: '2'               |             type: 'Identifier',
          }]                         |             name: 'subtract'
        }]                           |           },
      }]                             |           arguments: [{
    }                                |             type: 'NumberLiteral',
                                     |             value: '4'
									 |           }, {
                                     |             type: 'NumberLiteral',
                                     |             value: '2'
                                     |           }]
   								     |         }]
                                     |       }
                                     |     }]
                                     |   }
  ----------------------------------------------------------------------------
*/

func transformer(a ast) ast {
	newAst := ast{
		kind: "Program",
		body: []node{},
	}
	// a.context 是 newAst.body 的一个引用，目的是为了传递当前节点到其父节点的 context 中
	a.context = &newAst.body

	traverser(a, map[string]func(n *node, parent node){
		// 对 NumberLiteral 节点进行操作，新建对应节点并添加到父节点的 context 中
		"NumberLiteral": func(n *node, parent node) {
			*parent.context = append(*parent.context, node{
				kind:  "NumberLiteral",
				value: n.value,
			})
		},
		// 对 CallExpression 节点进行操作
		"CallExpression": func(n *node, parent node) {
			newNode := node{
				kind: "CallExpression",
				call: &node{
					kind: "Identifier",
					name: n.name,
				},
				arguments: new([]node),
			}
			// 将新节点的参数列表赋值给当前节点的 context，作为表达式参数的引用
			n.context = newNode.arguments

			// 判断父节点的类型是否为 CallExpression 节点
			// 不是则新建 ExpressionStatement 节点并进行添加
			// 否则直接添加新建节点到父节点的 context 中
			if parent.kind != "CallExpression" {
				esNode := node{
					kind:       "ExpressionStatement",
					expression: &newNode,
				}
				*parent.context = append(*parent.context, esNode)
			} else {
				*parent.context = append(*parent.context, newNode)
			}
		},
	})

	return newAst
}
