package stack

func backspaceCompare(s string, t string) bool {
	sStack := runeStack(s)
	tStack := runeStack(t)
	if len(sStack) != len(tStack) {
		return false
	}
	for i := 0; i < len(sStack); i++ {
		if sStack[i] != tStack[i] {
			return false
		}
	}
	return true
}

func runeStack(str string) []rune {
	stack := []rune{}
	for i := 0; i < len(str); i++ {
		if str[i] == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, rune(str[i]))
		}
	}
	return stack
}
