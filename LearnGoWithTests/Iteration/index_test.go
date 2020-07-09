package HelloWorld

import "testing"

const repeatCount = 5

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

// 基准测试: 您会看到该代码与测试非常相似。
// 结果中的`ns/op`表示：我们的功能平均运行136纳秒（在我的计算机上）。
// 注意默认情况下，基准按顺序运行。
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
