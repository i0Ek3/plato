package plato

import (
	"log"
	"strings"
)

/*
	第五步，实现代码生成器
	递归调用自身来打印 ast 中的每个节点，并将结果存入 string 中
*/

func generator(n node) string {
	switch n.kind {
	// 当节点类型为 Program
	case "Program":
		var result []string
		// 遍历当前节点的 body，并将遍历到的节点 v 传入 codeGenerator 自身去遍历，以获得最终的结果
		for _, v := range n.body {
			result = append(result, generator(v))
		}

		// 并为每个 result 加入换行
		return strings.Join(result, "\n")
	// 当节点类型为 ExpressionStatement，则递归便利当前节点的 expression，并在结尾加上分号
	case "ExpressionStatement":
		return generator(*n.expression) + ";"
	// 当节点类型为 CallExpression，递归遍历当前节点，以获取第一部分结果
	case "CallExpression":
		var result []string
		c := generator(*n.call)
		// 遍历当前节点的参数列表，并递归遍历节点 v，以获得结果
		for _, v := range *n.arguments {
			result = append(result, generator(v))
		}
		// 为 result 加入逗号，拼接所有字符串并加入相应的左右括号，返回结果
		res := strings.Join(result, ", ")

		return c + "(" + res + ")"
	// 当节点类型为 Identifier，直接返回当前节点的名字即可，即 Identifier 为函数操作，对应的是函数名
	case "Identifier":
		return n.name
	// 当节点类型为 NumberLiteral，即数字常量，则直接返回对应的 value 即可
	case "NumberLiteral":
		return n.value
	// 否则抛出错误并返回空串
	default:
		log.Fatal(n.kind)

		return ""
	}
}
