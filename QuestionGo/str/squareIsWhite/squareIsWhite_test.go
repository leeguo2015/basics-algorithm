package squareIsWhite

import "testing"

type Detail struct {
	want  bool
	mater string
}

var TestList = []Detail{
	{want: false, mater: "a1"},
	{want: true, mater: "h3"},
}

func Test_numDifferentIntegers(t *testing.T) {

	for _, detail := range TestList {
		if got := squareIsWhite(detail.mater); detail.want != got {
			t.Errorf("test failed want:%#v, got %#v, mater: %#v", detail.want, got, detail.mater)
		}
	}
	t.Log("pass all test case")

}
