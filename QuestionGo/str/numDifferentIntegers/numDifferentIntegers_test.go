package numDifferentIntegers

import "testing"

type Detail struct {
	want  int
	mater string
}

var TestList = []Detail{
	{want: 3, mater: "a123bc34d8ef34"},
	{want: 2, mater: "leet1234code234"},
	{want: 1, mater: "a1b01c001"},
}

func Test_numDifferentIntegers(t *testing.T) {

	for _, detail := range TestList {
		if got := numDifferentIntegers(detail.mater); detail.want != got {
			t.Errorf("test failed want:%#v, got %#v, mater: %#v", detail.want, got, detail.mater)
		}
	}
	t.Log("pass all test case")

}
