package main

import (
	"fmt"
	"ts_golang/invoke"
)

func main() {
	msg := invoke.Hello("Golang", "")
	fmt.Println(msg)

}
