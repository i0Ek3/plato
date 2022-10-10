package plato

func isNumber(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= '0' && n <= '9' {
		return true
	}
	return false
}

func isLetter(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= 'a' && n <= 'z' {
		return true
	}
	return false
}

func isRight(char string, a, b rune) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= a && n <= b {
		return true
	}
	return false
}
