package plato

/*
	第一步，实现 tokenizer
	利用 tokenizer 来解析一个目标语言表达式，即将该表达式转换为 token 数组
*/

type token struct {
	kind  string
	value string
}

// 将输入字符串 token 化
func tokenizer(input string) []token {
	// 先对当前输入追加换行
	input += "\n"
	// cur 用来追踪指针在当前代码中位置
	cur := 0
	// tokens 用来存储最终的 token
	var tokens []token

	// 遍历 input 中的字符
	for cur < len([]rune(input)) {
		// char 用来记录 cur 在 input 中所指向的字符
		char := string([]rune(input)[cur])

		// 如果是左括号，我们则将其 token 化，然后加入到 tokens 中
		if char == "(" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: "(",
			})
			// 向后移动 cur
			cur++
			// 然后继续下一个字符的检查
			continue
		}

		// 同上，用来检查是否是右括号
		if char == ")" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: ")",
			})
			cur++
			continue
		}

		// 检查是否是空格，空格可以区分字符是否被分割，我们直接忽略即可，向后移动 cur
		if char == " " {
			cur++
			continue
		}

		// 检查是否是数字，需要注意被空格分开的两个数字，这样就是两个 token
		if isNumber(char) {
			// value 用于累加数字结果到 char 中
			value := ""
			// 循环遍历，相当于将第一个数字 token 保存起来
			for isNumber(char) {
				value += char
				cur++
				char = string([]rune(input)[cur])
			}
			// 如果此时 char 不是数字了，我们将当前字符 token 化，添加到 tokens 中
			tokens = append(tokens, token{
				kind:  "number",
				value: value,
			})
			continue
		}

		// 检查是否为字符，即 Lisp 中函数名字的字母构成，基本操作同上面的 isNumber
		if isLetter(char) {
			value := ""
			for isLetter(char) {
				value += char
				cur++
				char = string([]rune(input)[cur])
			}
			tokens = append(tokens, token{
				kind:  "name",
				value: value,
			})
			continue
		}
		break
	}

	return tokens
}
