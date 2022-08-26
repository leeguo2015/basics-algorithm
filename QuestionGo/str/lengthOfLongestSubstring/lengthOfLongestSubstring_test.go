package lengthOfLongestSubstring

import (
	"testing"
)

//func TestLengthOfLongestSubstring(t *testing.T){
//	got :=  LengthOfLongestSubstring(" ")
//	want := 1   // 期望的结果
//
//	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
//		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
//	}
//}

func BenchmarkLengthOfLongestSubstring(t *testing.B) {

	for i := 0; i < 100; i++ {
		LengthOfLongestSubstring("annnadfafadfewfrreg")
	}

}
