package main

import "fmt"

func CountChar(str string, c rune) string {
	count := 0
	for _, char := range str {
		if char == c {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(CountChar("will", "willing"))
}
