package mergeSlice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceSort(t *testing.T) {
	s1 := []int{
		1, 2, 4,
	}
	s2 := []int{
		1, 2, 4,
	}

	result := SliceSort(s1, s2)
	want := []int{
		1, 1, 2, 2, 4, 4,
	}

	if reflect.DeepEqual(result, want) {
		fmt.Println("OK...")
	}

}
