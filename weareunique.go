package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	seen1 := make(map[rune]bool)
	for _, r := range str1 {
		seen1[r] = true
	}
	seen2 := make(map[rune]bool)
	for _, r := range str2 {
		seen2[r] = true
	}
	count := 0
	for r := range seen1 {
		if !seen2[r] {
			count++
		}
	}
	for r := range seen2 {
		if !seen1[r] {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
