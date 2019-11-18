package main

import (
	"C"
	"fmt"

	"github.com/ZoneTong/yunxingningmeng/ui"
)

func main() {
	ui.Init()
}

//export printhello
func printhello() {
	fmt.Println("hello my first!")
}
