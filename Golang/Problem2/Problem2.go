package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
Write a program that reads a string from the user and performs the following:
Counts the number of characters (including spaces and punctuation).
Counts the number of words (delimited by spaces).
Counts the number of uppercase and lowercase letters.
*/
func main() {
	// var input string
	fmt.Print("Enter a string: ")
	// fmt.Scanln(&input)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		// Now you have the entire line, including spaces, stored in the 'text' variable
		fmt.Println("You entered:", text)
		countCaps(text)

	}
	// countChars(input)
	// countWords(input)
}
func countChars(input string) {
	count := 0
	for _, char := range input {
		if char != ' ' {
			count++
		}
	}
	fmt.Println("total:", len(input))
	fmt.Printf("Number of characters: %d\n", count)
}

func countWords(input string) {
	count := 0
	words := 0
	for _, char := range input {
		if char == ' ' {
			words++
		}
	}
	count = words + 1
	fmt.Printf("Number of words: %d\n", count)
}
func countCaps(input string) {
	array_input := strings.Split(input, " ")
	upper_cnt := 0
	lower_cnt := 0
	var notUpper bool
	for _, v := range array_input {
		notUpper = false
		for _, v := range v {
			if unicode.ToUpper(v) != v {
				notUpper = true
				break
			}
		}
		if notUpper {
			lower_cnt++
		} else {
			upper_cnt++
		}
	}
	fmt.Printf("Total Upper Caps: %d, Lower Caps: %d", upper_cnt, lower_cnt)
}
