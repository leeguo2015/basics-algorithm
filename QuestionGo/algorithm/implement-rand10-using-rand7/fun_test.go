package implement_rand10_using_rand7

import (
	"fmt"
	"testing"
)

func TestRand10(t *testing.T) {
	l := []int{}
	m := make(map[int]int)
	for i := 0; i < 500; i++ {
		r := rand10()
		l = append(l, r)
		if _, ok := m[r]; ok {
			m[r]++
		} else {
			m[r] = 1
		}

	}
	fmt.Println(l)
	fmt.Println(m)
}
