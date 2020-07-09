package Integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	// 未定义Add时，运行，报错
	// 编写最少量的代码以运行测试并检查失败的测试输出
	// 当您有多个相同类型的参数（在我们的示例中为两个整数）而不是具有（x int，y int）时，可以将其缩短为（x，y int）。
	sum0 := Add(2, 2)
	sum1 := Add(1, 5)
	fmt.Println(sum1) // output:6
	expected := 4

	if sum0 != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, sum0)
	}
}

// Write enough code to make it pass
func Add(x, y int) int {
	//return 0 // err
	//return 4 // pass
	return x + y
}
