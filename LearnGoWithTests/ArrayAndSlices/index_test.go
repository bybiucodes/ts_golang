package HelloWorld

import (
	"reflect"
	"testing"
)

func assertCorrectMessage(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		assertCorrectMessage(t, sum(numbers), 15)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{2, 3, 5}
		assertCorrectMessage(t, sum(numbers), 10) // numbers -> [4] | cannot use numbers (type [4]int) as type [5]int in argument to sum
	})
}

func sum(numbers []int) int {
	// range允许您遍历数组。每次调用它都会返回两个值，即索引和值。我们选择通过使用_空白标识符来忽略索引值。
	// 数组的一个有趣特性是大小以其类型编码。如果尝试将[4] int传递给期望[5] int的函数，则它将无法编译。它们是不同的类型，因此与尝试将字符串传递给需要int的函数相同。
	// 可能会认为数组的长度固定很麻烦，而且大多数时候您可能不会使用它们！
	var addVal int
	for _, num := range numbers {
		addVal += num
	}
	//for i := 0; i < 5; i++ {
	//	res += numbers[i]
	//}
	return addVal
}

func TestSumAll(t *testing.T) {
	checkSums(t, sumAll([]int{1, 2}, []int{0, 9}), []int{3, 9})
}

func sumAll(numbersToSum ...[]int) []int {
	//lenthOfNumbers := len(numbersToSum)
	//sums := make([]int, lenthOfNumbers)
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, sum(numbers))
	}

	//for i, numbers := range numbersToSum {
	//	sums[i] = sum(numbers)
	//}
	return sums
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	checkSums(t, got, want)
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		}else{
			tail := numbers[1:] // "take from 1 to the end"
			sums = append(sums, sum(tail))
		}
	}
	return sums

}
