package main

import (
	"fmt"
	"strings"
)

func main() {
	var phrases []string

	for i := 0; i < 3; i++ {
		phrases = append(phrases, "Hello")
	}
	fmt.Println(strings.Join(phrases[:], " "))
}
