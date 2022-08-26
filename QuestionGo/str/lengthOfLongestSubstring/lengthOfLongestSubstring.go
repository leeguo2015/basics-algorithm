package lengthOfLongestSubstring

func LengthOfLongestSubstring(s string) int {
	maxlength := 0
	runeMap := make(map[string]int, 16)
	for i := 0; i < len(s); i++ {
		mapLen := 0
		for j := i; j < len(s); j++ {
			if _, ok := runeMap[string(s[j])]; ok {
				break
			} else {
				mapLen++
				runeMap[string(s[j])] = 1
			}
			if maxlength < mapLen {
				maxlength = mapLen
			}

		}
		for s2, _ := range runeMap {
			delete(runeMap, s2)
		}
	}
	return maxlength
}
