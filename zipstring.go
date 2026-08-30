package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	var res string
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		res += strconv.Itoa(j-i) + string(s[i])
		i = j
	}
	return res
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
