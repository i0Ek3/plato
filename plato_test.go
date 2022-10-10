package plato

import (
	"reflect"
	"testing"
)

var testTokens = []token{
	token{
		kind:  "paren",
		value: "(",
	},
	token{
		kind:  "name",
		value: "add",
	},
	token{
		kind:  "number",
		value: "10",
	},
	token{
		kind:  "paren",
		value: "(",
	},
	token{
		kind:  "name",
		value: "subtract",
	},
	token{
		kind:  "number",
		value: "10",
	},
	token{
		kind:  "paren",
		value: "(",
	},
	token{
		kind:  "name",
		value: "add",
	},
	token{
		kind:  "number",
		value: "6",
	},
	token{
		kind:  "paren",
		value: "(",
	},
	token{
		kind:  "name",
		value: "subtract",
	},
	token{
		kind:  "number",
		value: "4",
	},
	token{
		kind:  "number",
		value: "2",
	},
	token{
		kind:  "paren",
		value: ")",
	},
	token{
		kind:  "paren",
		value: ")",
	},
	token{
		kind:  "paren",
		value: ")",
	},
	token{
		kind:  "paren",
		value: ")",
	},
}

var testAst = ast{
	kind: "Program",
	body: []node{
		node{
			kind: "CallExpression",
			name: "add",
			params: []node{
				node{
					kind:  "NumberLiteral",
					value: "10",
				},
				node{
					kind: "CallExpression",
					name: "subtract",
					params: []node{
						node{
							kind:  "NumberLiteral",
							value: "10",
						},
						node{
							kind: "CallExpression",
							name: "add",
							params: []node{
								node{
									kind:  "NumberLiteral",
									value: "6",
								},
								node{
									kind: "CallExpression",
									name: "subtract",
									params: []node{
										node{
											kind:  "NumberLiteral",
											value: "4",
										},
										node{
											kind:  "NumberLiteral",
											value: "2",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestTokenizer(t *testing.T) {
	result := tokenizer(TESTINPUT)
	if !reflect.DeepEqual(result, testTokens) {
		t.Error("\nExpected:", testTokens, "\nGot:", result)
	}
}

func TestParser(t *testing.T) {
	result := parser(testTokens)
	if !reflect.DeepEqual(result, testAst) {
		t.Error("\nExpected:", testAst, "\nGot:", result)
	}
}

func BenchmarkTokenizer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tokenizer(TESTINPUT)
	}
}

func BenchmarkParser(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parser(testTokens)
	}
}
