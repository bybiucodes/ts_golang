package HelloWorld

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string) {
	// Our first round of code was not easy to test because it wrote data to somewhere we couldn't control.

	// Catch the val of Buffer -> byte.Buffer
	fmt.Fprintf(w, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world.")
}