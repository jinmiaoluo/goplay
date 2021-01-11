package main

import "fmt"

var parentheses map[rune]rune = map[rune]rune{
	']': '[',
	'}': '{',
	')': '(',
}

func isValid(s string) bool {
	// Special case judgment
	if len(s) == 0 {
		return true
	}

	hash := make(map[rune]int)
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			// stack in
			stack = append(stack, v)
			hash[v]++
		}
		if (v == ']') || (v == ')') || (v == '}') {
			if hash[parentheses[v]] > 0 {
				hash[parentheses[v]]--
				if stack[len(stack)-1] == parentheses[v] {
					// stack out
					stack = stack[:len(stack)-1]
				} else {
					fmt.Println("The order of parentheses are wrong")
					return false
				}
			} else {
				fmt.Printf("`%s` is redundant\n", string(v))
				return false
			}
		}
	}
	if len(stack) > 0 {
		for i := 0; i < len(stack); i++ {
			fmt.Printf("`%s` is redundant\n", string(stack[i]))
		}
		return false
	}
	return true
}

func main() {

	var str string

	// Incorrect demo 1
	str = "[music(description).mp3](https://example.com/music(description).mp3))"
	if isValid(str) {
		fmt.Printf("string: `%s` is valid!\n", str)
	} else {
		fmt.Printf("string: `%s` isn't valid!\n", str)
	}

	// Incorrect demo 2
	str = "[[music(description).mp3](https://example.com/music(description).mp3)"
	if isValid(str) {
		fmt.Printf("string: `%s` is valid!\n", str)
	} else {
		fmt.Printf("string: `%s` isn't valid!\n", str)
	}

	// Correct demo
	str = "[music(description).mp3](https://example.com/music(description).mp3)"
	if isValid(str) {
		fmt.Printf("String: `%s` is valid!\n", str)
	} else {
		fmt.Printf("String: `%s` isn't valid!\n", str)
	}
}
