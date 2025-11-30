package main

import (
	"fmt"
)

func BinaryLength(num int) int {
	if num == 0 {
		return 0
	}
	return 1 + BinaryLength(num/256)
}
func main() {
	fmt.Println("check: ", BinaryLength(1234567890123456789))
}
