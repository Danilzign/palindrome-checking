package main

import "fmt"

func main() {
	fmt.Println(palindrom("level"))
}

func palindrom(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reverseString := string(runes)
	if str == reverseString {
		return "this string is a palindrom"
	}
	return "this string is not a palindrom"
}
