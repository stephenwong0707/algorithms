package slidingwindow

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	result := []int{}

	pMap := wordCount(p)
	sMap := wordCount(s[:len(p)])
	for i := 0; i < len(s)-len(p)+1; i++ {
		match := true
		if i > 0 {
			lastRune := rune(s[i-1])
			sMap[lastRune] = sMap[lastRune] - 1
			newRune := rune(s[i+len(p)-1])
			sMap[newRune] = sMap[newRune] + 1
		}
		for k, v := range pMap {
			if sMap[k] != v {
				match = false
			}
		}
		if match {
			result = append(result, i)
		}
	}
	return result
}

func wordCount(s string) map[rune]int {
	wc := map[rune]int{}
	for _, char := range s {
		wc[char] = wc[char] + 1
	}
	return wc
}
