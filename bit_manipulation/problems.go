package bitmanipulation

import (
	"fmt"
)

func BitManipulationProblems() {
	convertDecimalToBinary(13)
	convertDecimalToBinary(1333)

	convertBinaryToDecimal(1101)
}

func convertDecimalToBinary(num int) {
	ans := 0
	place := 1
	decimalNum := num
	for num > 0 {
		binaryDigit := num % 2
		num = num / 2
		ans = ans + binaryDigit*place
		place = place * 10
	}
	fmt.Printf("%d ===> %d\n", decimalNum, ans)
}

func convertBinaryToDecimal(num int) {
	ans := 0
	power := 1
	bianryNum := num
	for num > 0 {
		digit := num % 10
		ans += digit * power
		num = num / 10
		power *= 2
	}
	fmt.Printf("%d ===> %d\n", bianryNum, ans)
}
