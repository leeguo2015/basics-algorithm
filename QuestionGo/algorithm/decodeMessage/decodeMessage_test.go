package decodeMessage

import (
	"testing"
)

type Detail struct {
	want    string
	key     string
	message string
}

var TestList = []Detail{
	{want: "this is a secret", key: "the quick brown fox jumps over the lazy dog", message: "vkbs bs t suepuv"},
	{want: "the five boxing wizards jump quickly", key: "eljuxhpwnyrdgtqkviszcfmabo", message: "zwx hnfx lqantp mnoeius ycgk vcnjrdb"},
}

func Test_decodeMessage(t *testing.T) {
	for _, detail := range TestList {
		if got := decodeMessage(detail.key, detail.message); detail.want != got {
			t.Errorf("test failed want:%#v, got %#v", detail.want, got)
		}
	}
	t.Log("pass all test case")

}
