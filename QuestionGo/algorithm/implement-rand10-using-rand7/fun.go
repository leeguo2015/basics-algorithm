package implement_rand10_using_rand7

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var r = time.Now().UnixNano()

func rand10() int {
	r := float64(rand7()) / 7 * 10
	if rand7()%2 == 1 {
		return int(math.Ceil(r))
	}
	return int(math.Floor(r))
}

func rand7() int {
	a := getRand()
	fmt.Println(a)
	return a
}

func getRand() int {
	return rand.Intn(7) + 1
}

func init() {
	rand.Seed(r)
}
