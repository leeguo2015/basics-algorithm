package closestCost

import "testing"

type iceMater struct {
	baseCosts    []int
	toppingCosts []int
	target       int
}

type Detail struct {
	want  int
	mater iceMater
}

var TestList = []Detail{
	{want: 10, mater: iceMater{[]int{1, 7}, []int{3, 4}, 10}},
	{want: 17, mater: iceMater{[]int{2, 3}, []int{4, 5, 100}, 18}},
	{want: 8, mater: iceMater{[]int{3, 10}, []int{2, 5}, 9}},
	{want: 10, mater: iceMater{[]int{10}, []int{1}, 1}},
}

func Test_closestCost(t *testing.T) {

	for _, detail := range TestList {
		if got := closestCost(detail.mater.baseCosts, detail.mater.toppingCosts, detail.mater.target); detail.want != got {
			t.Errorf("test failed want:%#v, got %#v, mater: %#v", detail.want, got, detail.mater)
		}
	}
	t.Log("pass all test case")

}
