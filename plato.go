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
