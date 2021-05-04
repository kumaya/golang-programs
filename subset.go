package main

import (
	"fmt"
	"strings"
)

// subsequence: subset with ordering
func subsequence(inp, out string) {
	if inp == "" {
		fmt.Printf("%q\n", out)
		return
	}
	out1 := out + string(inp[0]) // consider first val from input
	out2 := out                  // do not consider first value from input
	subsequence(inp[1:], out1)
	subsequence(inp[1:], out2)
}

func subsequenceWithSpaces(inp, out string) {
	if inp == "" {
		fmt.Printf("%q\n", out)
		return
	}
	out1 := out + "_" + string(inp[0])
	out2 := out + string(inp[0])
	subsequenceWithSpaces(inp[1:], out1)
	subsequenceWithSpaces(inp[1:], out2)
}

func subsequenceWithCase(inp, out string) {
	if inp == "" {
		fmt.Printf("%q\n", out)
		return
	}
	out1 := out + strings.ToUpper(string(inp[0]))
	out2 := out + strings.ToLower(string(inp[0]))
	subsequenceWithCase(inp[1:], out1)
	subsequenceWithCase(inp[1:], out2)
}

func subsequenceWithDigits(inp, out string) {
	if inp == "" {
		fmt.Printf("%q\n", out)
		return
	}
	out1 := out + strings.ToUpper(string(inp[0]))
	out2 := out + strings.ToLower(string(inp[0]))
	if inp[0] > 57 {
		subsequenceWithDigits(inp[1:], out1)
		subsequenceWithDigits(inp[1:], out2)
	} else {
		out3 := out + string(inp[0])
		subsequenceWithDigits(inp[1:], out3)
	}
}

func balancedParenthesis(out string, open, close int) {
	if open == 0 && close == 0 {
		fmt.Printf("%q\n", out)
		return
	}
	//if open == 0 {
	//	balancedParenthesis(out+")", open, close-1)
	//} else if open < close {
	//	balancedParenthesis(out+"(", open-1, close)
	//	balancedParenthesis(out+")", open, close-1)
	//} else {
	//	balancedParenthesis(out+"(", open-1, close)
	//}
	if open != 0 {
		balancedParenthesis(out+"(", open-1, close)
	}
	if close > open {
		balancedParenthesis(out+")", open, close-1)
	}
}

func prefixes(n, ones, zeroes int, out string) {
	if n == 0 {
		fmt.Printf("%s\n", out)
		return
	}
	if ones-zeroes <= 1 {
		out1 := out + "1"
		prefixes(n-1, ones+1, zeroes, out1)
	} else {
		out1 := out + "1"
		out2 := out + "0"
		prefixes(n-1, ones+1, zeroes, out1)
		prefixes(n-1, ones, zeroes+1, out2)
	}
}

func main() {
	s := "ab"
	fmt.Println("Subsequence: ")
	subsequence(s, "")
	fmt.Println("Subsequence with spaces: ")
	subsequenceWithSpaces(s[1:], string(s[0]))
	fmt.Println("Subsequence with cases: ")
	subsequenceWithCase(s, "")
	fmt.Println("Subsequence with digits: ")
	s = "a1b1"
	subsequenceWithDigits(s, "")
	fmt.Println("Balanced parenthesis: ")
	n := 3
	balancedParenthesis("(", n-1, n)
	fmt.Println("All prefixes #1s > #0s: ")
	n = 5
	prefixes(n, 0, 0, "")
}
