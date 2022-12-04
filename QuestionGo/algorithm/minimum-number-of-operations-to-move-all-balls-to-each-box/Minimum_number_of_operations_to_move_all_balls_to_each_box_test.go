package minimum_number_of_operations_to_move_all_balls_to_each_box

import (
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T) {
	testData := map[string][]int{
		"110":    {1, 1, 3},
		"001011": {11, 8, 5, 4, 3, 4},
	}
	for k, v := range testData {
		if got := MinOperations(k); !reflect.DeepEqual(got, v) {
			t.Errorf("Test results are not equal: want %#v got: %#v", v, got)
		}
	}
}

func BenchmarkMinOperations(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
		MinOperations("001011")
	}
}
