package revers_str

import "testing"

//func TestReverseString(t *testing.T) {
//	data := []byte{'h', 'e', 'l', 'l', 'o'}
//	got := ReverseString(data)
//	want := []byte{'o', 'l', 'l', 'e', 'h'} // 期望的结果
//
//	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
//		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
//	}
//}

//go test -bench=ReverseString -benchmem
func BenchmarkReverseString(t *testing.B) {
	for i := 0; i < 100; i++ {
		ReverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	}

}
