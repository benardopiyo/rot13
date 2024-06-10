// Write a program that takes a string and displays it,
// replacing each of its letters by the letter
// 13 spots ahead in alphabetical order.
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1]

	var res string
	for _, char := range arg {
		if char >= 'a' && char <= 'z' {
			res = res + string((char-'a'+13)%26+'a')
		} else if char >= 'A' && char <= 'Z' {
			res = res + string((char-'A'+13)%26+'A')
		} else {
			res = res + string(char)
		}
	}

	for _, val := range res {
		z01.PrintRune(val)
	}
	z01.PrintRune(10)
}
