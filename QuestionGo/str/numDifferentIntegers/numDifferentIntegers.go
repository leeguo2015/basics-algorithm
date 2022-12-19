package numDifferentIntegers

import "fmt"

/*
https://leetcode.cn/problems/number-of-different-integers-in-a-string/description/
1805. 字符串中不同整数的数目
给你一个字符串 word ，该字符串由数字和小写英文字母组成。

请你用空格替换每个不是数字的字符。例如，"a123bc34d8ef34" 将会变成 " 123  34 8  34" 。注意，剩下的这些整数为（相邻彼此至少有一个空格隔开）："123"、"34"、"8" 和 "34" 。

返回对 word 完成替换后形成的 不同 整数的数目。

只有当两个整数的 不含前导零 的十进制表示不同， 才认为这两个整数也不同。
*/

func numDifferentIntegers1(maters string) int {
	words := make(map[string]int)
	clearZero := func(s string) string {
		for i := 0; i < len(s); i++ {
			if s[i] != 48 {
				return s[i:]
			}
			if s[i] == 48 && i+1 == len(s) {
				return "0"
			}
		}
		return s
	}
	word := ""
	for i, mater := range []byte(maters) {

		if 48 <= mater && mater <= 57 {
			word = word + string(mater)
		} else {
			if word != "" {
				words[clearZero(word)] = 0
				word = ""
			}
		}
		if len(maters) == i+1 && word != "" {
			words[clearZero(word)] = 0
		}
	}
	return len(words)
}

func numDifferentIntegers(maters string) int {
	left := 0
	right := 0
	isRight := true
	words := make(map[string]int)
	for i := 0; i < len(maters); i++ {

		if 48 < maters[i] && maters[i] <= 57 {
			if isRight {
				right = i
				isRight = false
			} else {
				if right == left {
					left = i
					continue
				}
				left = i + 1
			}
			continue
		}
		if !isRight || i+1 == len(maters) {
			isRight = true
			words[maters[right:left]] = 0
		}
		fmt.Println(right, left, i)
		fmt.Printf("%#v\n", words)
	}
	fmt.Printf("%#v\n", words)
	return len(words)
}
