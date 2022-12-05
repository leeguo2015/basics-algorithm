package selectionSort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	got := SelectionSort([]int{4, 1, 3, 5, 6, 7, 6, 66, 77, 1, 6})
	want := []int{1, 1, 3, 4, 5, 6, 6, 6, 7, 66, 77}
	fmt.Println(got, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("error--->\ngot: %v \n want: %v\n", got, want)
	}
}
