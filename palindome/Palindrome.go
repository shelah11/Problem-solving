package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(palindrome("hereh"))
}

func palindrome(y string) bool {

	input := strings.ToLower(y)
	inputWithoutSpaces := strings.ReplaceAll(input, " ", "")
	var res []string

	for i := len(inputWithoutSpaces) - 1; i >= 0; i-- {

		res = append(res, string(inputWithoutSpaces[i]))
	}

	result2 := strings.Join(res, "")
	if result2 == inputWithoutSpaces {
		return true
	}
	return false
}
