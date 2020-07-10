package HelloWorld

import (
	"os"
	"testing"
)


func TestGreet(t *testing.T) {
	// The buffer type from the bytes package implements the Writer interface.
	// So we'll use it in our test to send in as our Writer and then we can check what was written to it after we invoke Greet
	Greet(os.Stdout, "Ilice")
	//buffer := bytes.Buffer{}
	//Greet(&buffer, "Str1")
	//
	//got := buffer.String()
	//want := "Hello, Str1"
	//
	//if got != want {
	//	t.Errorf("got %q want %q", got, want)
}
