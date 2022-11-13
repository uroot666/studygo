package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("babcbef", "b")
	want := []string{"", "a", "c", "ef"}

	if !reflect.DeepEqual(ret, want) {
		// 测试用例失败了
		t.Errorf("want: %v but got:%v\n", want, ret)
	}
}

func Test2Split(t *testing.T) {
	ret := Split("abcef", "bc")
	want := []string{"a", "ef"}
	if !reflect.DeepEqual(ret, want) {
		t.Fatalf("want: %v but got:%v\n", want, ret)
	}
}

// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
