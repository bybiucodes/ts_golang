package HelloWorld

// 常量可以提高应用程序的性能，因为它可以避免每次调用Hello时创建“ Hello”字符串实例。
// 需要明确的是，在这个示例中，性能提升微不足道！但是值得考虑的是创建常量以捕获值的含义，有时还可以提高性能。
const spanish = "Spanish"
const french = "Franch"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	/*
	 * 您有许多if语句检查特定值时，通常会改用switch语句。如果我们希望以后再添加更多语言支持，则可以使用switch重构代码以使其更易于阅读和扩展。
	 *
	 */
	//if language == spanish {
	//	return spanishHelloPrefix + name
	//}

	//if language == french {
	//	return frenchHelloPrefix + name
	//}
	return greetingPrefix(language) + name
}

// 最简单的重构是将某些功能提取到另一个功能中。
/*
 * 一些新概念： 在函数签名中，我们指定了一个命名返回值（prefix string）。
 * 这将在函数中创建一个称为prefix的变量。 将为其分配“零”值。这取决于类型，例如int为0，而字符串为“”。
 * 您可以通过调用`return`而不是`return prefix`来返回它设置的任何内容。
 * 这将显示在您的函数的“Go Doc”中，这样​​可以使您的代码意图更加清晰。
 * 如果其他case语句都不匹配，则将分支到switch case中的default。 函数名称以小写字母开头。
 * 在Go语言中，公共功能以大写字母开头，私人功能以小写字母开头。我们不希望将算法的内部知识公开，因此我们将此函数设为私有。
 */
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
