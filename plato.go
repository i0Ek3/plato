/**
	编译器的编译主要依赖于以下三个步骤，分别是：

	1. 解析：将原生 code 解析成一种抽象的代码表示

		1.1 词法分析：
			利用 tokenizer 或者 lexer 将 raw code 分割成 tokens
			tokens 由很小的 object 数组构成，用来描述语法，可以是数字，标签，逗号，操作符等等

		1.2 语法分析：
			将 1.1 中的 token 重新格式化为可以描述每个语法和其他语法相关联的一种表示，这也是我们常说的中间（表示）代码或者 ast
			ast 是一种深度嵌套的对象，它以一种既易于使用又能告诉我们大量信息的方式来表示代码

	例如，转换 Lisp 表达式 (add 2 (subtract 4 2)) 为 C 语言表达式 add(2, subtract(4, 2));

	经由 1.1 之后的 Token 类似于这样：

    [
      { type: 'paren',  value: '('        },
      { type: 'name',   value: 'add'      },
      { type: 'number', value: '2'        },
      { type: 'paren',  value: '('        },
      { type: 'name',   value: 'subtract' },
      { type: 'number', value: '4'        },
      { type: 'number', value: '2'        },
      { type: 'paren',  value: ')'        },
      { type: 'paren',  value: ')'        }
    ]

	经由 1.2 之后的 AST 类似于这样：

    {
      type: 'Program',
      body: [{
        type: 'CallExpression',
        name: 'add',
        params: [{
          type: 'NumberLiteral',
          value: '2'
        }, {
          type: 'CallExpression',
          name: 'subtract',
          params: [{
            type: 'NumberLiteral',
            value: '4'
          }, {
            type: 'NumberLiteral',
            value: '2'
          }]
        }]
      }]
    }


	2. 转换：将 1 中生成的抽象代码表示按照编译器想要的行为进行操作

		这一步会对 1.2 中生成的 AST 进行操作，除了做一些修改以外，还会将其转换为其他语言，那具体是如何转换 AST 的呢？

		首先，AST 中每个节点的结构其实是比较相似的，都有相同的属性或者类型，每个节点又是一个单独的 AST，根据 1.2 中的 AST，我们可以看到某个 AST 节点

		其中，NumberLiteral 节点包含有以下属性：
				{
					type: 'NumberLiteral',
					value: '2'
				}
		CallExpression 节点包含有以下属性：
				{
					type: 'CallExpression',
					name: 'subtract',
					params: [...其他嵌套节点]
				}

		当我们转换 AST 时，我们可以对节点中的元素进行添加、删除、替换等操作，我们也可以添加、删除一个节点，或者基于该节点再新建一个新的节点

		这里我们仅关注于创建一个目标语言的新节点，为了能够获取所有的 AST 节点，我们需要利用 dfs 来遍历它们，即深度优先遍历

		对于 1.2 中生成的 AST，我们的遍历顺序如下：

				1 Program
					2 CallExpression(add)
						3 NumberLiteral(2)
					4 CallExpression(subtract)
						5 NumberLiteral(4)
						6 NumberLiteral(2)

	3. 代码生成：将 2 中生成的表示转换为目标语言的代码

		简单来说，大多数编译器的代码生成就是消除转换表示中的 AST 和字符串化的代码，我们这里会使用刚才创建好的 AST 来进行转换


	基本上，一个简易的 compiler 就完成了，但这并不意味着所有的 compiler 都这么简单。
	不同的 compiler 有不同的作用，复杂度也会不同，但通用的处理步骤就是这些。
*/

package plato

/*
	第六步，实现编译器
	即将前面几个步骤组合在一起，具体步骤如下：

    	1. input  => tokenizer   => tokens
    	2. tokens => parser      => ast
    	3. ast    => transformer => newAst
    	4. newAst => generator   => result
*/

func Compiler(input string) string {
	return compiler(input)
}

func compiler(input string) string {
	tokens := tokenizer(input)
	ast := parser(tokens)
	newAst := transformer(ast)
	result := generator(node(newAst))

	return result
}
