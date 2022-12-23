package finalValueAfterOperations

import "testing"

type Detail struct {
	want  int
	mater []string
}

var TestList = []Detail{
	{want: 1, mater: []string{"--X", "X++", "X++"}},
	{want: 3, mater: []string{"++X", "++X", "X++"}},
	{want: 0, mater: []string{"X++", "++X", "--X", "X--"}},
}

func Test_finalValueAfterOperations(t *testing.T) {
	for _, detail := range TestList {
		if got := finalValueAfterOperations(detail.mater); detail.want != got {
			t.Errorf("test failed want:%#v, got %#v, mater: %#v", detail.want, got, detail.mater)
		}
	}
	t.Log("pass all test case")

}
