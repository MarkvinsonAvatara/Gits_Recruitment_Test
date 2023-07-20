package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if isBalanced(input) {
		fmt.Println("Output: YES")
	} else {
		fmt.Println("Output: NO")
	}
}

func isBalanced(input string) bool {
	stack := make([]int32, 0)

	for _, ch := range input {
		if isOpeningBracket(ch) {
			stack = append(stack, ch)
		} else if isClosingBracket(ch) {
			if len(stack) == 0 || !isMatchingBracket(stack[len(stack)-1], ch) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isOpeningBracket(ch int32) bool {
	return ch == '(' || ch == '[' || ch == '{'
}

func isClosingBracket(ch int32) bool {
	return ch == ')' || ch == ']' || ch == '}'
}

func isMatchingBracket(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '[':
		return close == ']'
	case '{':
		return close == '}'
	default:
		return false
	}
}
