package main

import (
	"fmt"
)

// Reverse return the reverse of the str
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(Reverse("hello,世界"))
	fmt.Println("hello,世界")
}
