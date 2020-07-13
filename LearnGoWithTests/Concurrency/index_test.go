package HelloWorld

import (
	"reflect"
	"testing"
)

// 它将检查到的每个URL的映射都返回一个布尔值-正确响应为true，错误响应为false。
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}

func mockWebsiteChecker(url string) bool {
	if url == "http://abc.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google0.com",
		"http://google1.com",
		"http://google2.com",
	}
	want := map[string]bool{
		"http://google0.com": true,
		"http://google1.com": true,
		"http://google2.com": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Want %v, got %v", want, got)
	}
}
