package HelloWorld

import (
	"testing"
)

/*
 * 编写测试就像编写函数一样，遵守一些规则：
 * 	|_ 它需要在一个名为xxx_test.go的文件中
 * 	|_ 测试函数必须从单词Test开始
 * 	|_ 测试函数只接受一个参数 t *testing.T
 * 	|_ t *testing.T是测试框架的“hook”，所以想要失败时，可以执行像 t.Fail()
 */
/*
 * t.Errorf:
 * 	|_ t上的Errorf方法，它将打印出一条消息并使测试失败。f代表format，它允许我们构建一个字符串，并将值插入占位符值%q中。当您使测试失败时，应该清楚它是如何工作的。
 * 	|_ 你可以在fmt go文档中阅读更多关于占位符字符串的信息。对于测试来说，%q非常有用，因为它将值包装在双引号中。
 */
func TestHello(t *testing.T) {
	/*
	 * 我们已经将断言重构为一个函数。
	 * 	|_ 这减少了重复，提高了测试的可读性。
	 * 	|_ 在Go中，你可以在其他函数中声明函数，并将它们赋值给变量。然后可以像调用普通函数一样调用它们。
	 * 	|_ 我们需要通过 t *testing.T 这样我们就可以在需要的时候告诉测试代码失败。
	 * * * *
	 * 	如果我们运行我们的测试，我们应该看到它满足新的需求，我们没有意外破坏其他功能。
	 * ***
	 */
	assertCorrectMessage := func(t *testing.T, got, want string) {
		/*
		 * t.Helper(): 这将帮助其他开发人员更容易地跟踪问题。如果您仍然不理解，将其注释掉，使测试失败并观察测试输出。
		 */
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	/*
	 * 在这里，我们将介绍测试工具中的另一个工具——子测试。有时，围绕一个“事物”将测试分组，然后使用描述不同场景的子测试是有用的。
	 * 你的测试清楚地说明代码需要做什么是很重要的。
	 */
	t.Run("saying hello to peaple", func(t *testing.T) {
		got := Hello("World","")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	// When our function is called with an empty string it defaults to printing "Hello, World", rather than "Hello, ".
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("franch", "Franch")
		want := "Bonjour, franch"
		assertCorrectMessage(t, got, want)
	})

}
