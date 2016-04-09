package main

import (
	"fmt"
	"gopkg.in/macaron.v1"
)

func main() {
	fmt.Println("this is test.")
	m := macaron.Classic()
	m.Get("/", func() string {
		return "hello"
	})
	m.Run()
}
