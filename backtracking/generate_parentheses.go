package backtracking

func generateParenthesis(n int) []string {
	var result []string
	var gen func(string, int)
	gen = func(str string, maxLen int) {
		if len(str) < maxLen {
			gen(str+"(", maxLen)
			gen(str+")", maxLen)
		} else {
			if validateParentheses(str) {
				result = append(result, str)
			}
		}
	}
	gen("", n*2)
	return result
}

func validateParentheses(str string) bool {
	open := 0
	for _, x := range str {
		if x == '(' {
			open++
		} else if x == ')' {
			if open > 0 {
				open--
			} else {
				return false
			}
		}
	}
	return open == 0
}
