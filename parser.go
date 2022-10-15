package plato

import "log"

/*
	第二步，实现解析器
	我们将第一步中生成的 token 数组转换成 AST

	[{ type: 'paren', value: '('}, ...] ==> { type: 'Program', body: [...] }
*/

// 刻画 AST 的 node 结构体
type node struct {
	kind  string
	value string
	name  string

	call       *node
	expression *node
	body       []node
	params     []node
	arguments  *[]node
	context    *[]node
}

type ast node

// pc 用于计数
var pc int

// pt 用于存储 token 切片
var pt []token

func parser(tokens []token) ast {
	pc, pt = 0, tokens

	// 创建一个 ast，其根节点为 Program
	ast := ast{
		kind: "Program",
		body: []node{},
	}

	// 当前计数器还没有到达 token 结尾，则继续添加节点
	for pc < len(pt) {
		// walk 内部会进行递归操作，然后返回一个节点
		ast.body = append(ast.body, walk())
	}

	return ast
}

// walk 函数返回一个 ast 节点，利用递归来完成嵌套节点的获取
func walk() node {
	// 获取 pc 当前指向的 token
	token := pt[pc]

	// 如果当前 token 类型为 number，则计数器 +1
	// 并返回一个 NumberLiteral 节点，设置其值为当前 token 对应的 value
	if token.kind == "number" {
		pc++

		return node{
			kind:  "NumberLiteral",
			value: token.value,
		}
	}

	// 如果当前 token 类型为括号且是左括号
	if token.kind == "paren" && token.value == "(" {
		// 跳过，往后查找括号后面的表达式
		pc++
		token = pt[pc]
		// 新建一个 CallExpression 节点
		// 该节点的 name 为当前 token 的 value，并初始化参数切片
		newNode := node{
			kind:   "CallExpression",
			name:   token.value,
			params: []node{},
		}
		// 继续检查下一个表达式
		pc++
		token = pt[pc]

		// 循环遍历寻找 CallExpression 节点的所有参数并添加到 params 中，直到遇到对应的右括号
		// 即第一部分判断 token，后面的部分判断嵌套的表达式情况
		for token.kind != "paren" || (token.kind == "paren" && token.value != ")") {
			// 这里再次使用 walk 获取节点
			newNode.params = append(newNode.params, walk())
			token = pt[pc]
		}
		pc++

		return newNode
	}
	log.Fatal(token.kind)

	return node{}
}
