package checkPowersOfThree

import "testing"

type Detail struct {
	want  bool
	mater int
}

var TestList = []Detail{
	{want: true, mater: 91},
	{want: true, mater: 12},
	{want: false, mater: 21},
}

func Test_numDifferentIntegers(t *testing.T) {
	for _, detail := range TestList {
		if got := checkPowersOfThree(detail.mater); detail.want != got {
			t.Errorf("test failed want:%#v, got %#v, mater: %#v", detail.want, got, detail.mater)
		}
	}
	t.Log("pass all test case")

}
