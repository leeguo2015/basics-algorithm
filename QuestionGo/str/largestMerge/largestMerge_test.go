package largestMerge

import "testing"

type Detail struct {
	want  string
	mater []string
}

var TestList = []Detail{
	{want: "cbcabaaaaa", mater: []string{"cabaa", "bcaaa"}},
	{want: "abdcabcabcaba", mater: []string{"abcabc", "abdcaba"}},
}

func Test_finalValueAfterOperations(t *testing.T) {
	for _, detail := range TestList {
		if got := largestMerge(detail.mater[0], detail.mater[1]); detail.want != got {
			t.Errorf("test failed want:%#v, got %#v, mater: %#v", detail.want, got, detail.mater)
		}
	}
	t.Log("pass all test case")

}
