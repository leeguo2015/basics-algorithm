package checkXMatrix

import "testing"

type Detail struct {
	want  bool
	mater [][]int
}

var TestList = []Detail{
	{want: true, mater: [][]int{
		{2, 0, 0, 1},
		{0, 3, 1, 0},
		{0, 5, 2, 0},
		{4, 0, 0, 2},
	}},
	{want: false, mater: [][]int{
		{5, 7, 0},
		{0, 3, 1},
		{0, 5, 0},
	}},
}

func Test_checkXMatrix(t *testing.T) {

	for _, detail := range TestList {
		if got := checkXMatrix(detail.mater); detail.want != got {
			t.Errorf("test failed want:%#v, got %#v, mater: %#v", detail.want, got, detail.mater)
		}
	}
	t.Log("pass all test case")

}
