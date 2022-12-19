package insertionSort

/*
插入排序
*/

import (
	"reflect"
	"testing"
)

type Detail struct {
	want  []int
	mater []int
}

var TestList = []Detail{
	{want: []int{1, 1, 3, 4, 5, 6, 6, 6, 7, 66, 77}, mater: []int{4, 1, 3, 5, 6, 7, 6, 66, 77, 1, 6}},
}

func Test_SelectionSort(t *testing.T) {

	for _, detail := range TestList {
		if got := insertionSort(detail.mater); !reflect.DeepEqual(got, detail.want) {
			t.Errorf("test failed want:%#v, got %#v, mater: %#v", detail.want, got, detail.mater)
		}
	}
	t.Log("pass all test case")

}
