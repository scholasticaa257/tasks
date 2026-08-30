package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	str := os.Args[1]
	spaceCount := 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '\t' {
			spaceCount++
		} else {
			if spaceCount > 0 && i != spaceCount {
				fmt.Print(" ")
			}
			spaceCount = 0
			fmt.Print(string(str[i]))
		}
	}

	fmt.Println()
}
