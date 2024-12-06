package main

import (
	"fmt"
	"strconv"
)

func sumBinaryDigits(n int) int {
	binaryStr := strconv.FormatInt(int64(n), 2)
	sum := 0
	for _, char := range binaryStr {
		if char == '1' {
			sum++
		}
	}
	return sum
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Printf("Sum of binary digits of %d is: %d\n", number, sumBinaryDigits(number))
}
