package revers_str

func ReverseString(s []byte) []byte {
	l := len(s)
	h := l / 2
	for i := 0; i < l; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
		if i >= h-1 {
			break
		}
	}
	return s
}
