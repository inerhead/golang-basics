package myPackage

import (
	"fmt"
	"strings"
)

func ShowPalindrome() {
	word := "Ama"

	is := isPalindrome(word)
	fmt.Printf("\n\n")
	fmt.Println("STRING")
	fmt.Printf("%q is Palindrome %v", word, is)
	fmt.Printf("\n\n")
}

func isPalindrome(word string) bool {
	toLowerCase := strings.ToLower(word)
	var invert string
	for i := len(toLowerCase) - 1; i >= 0; i-- {
		invert += string(toLowerCase[i])
	}

	if invert == toLowerCase {
		return true
	}

	return false

}
