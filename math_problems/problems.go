package mathproblems

import "fmt"

func CountDigits(num int) int {
	if num < 0 {
		return 0
	}

	if num == 0 {
		fmt.Println(1)
		return 1
	}

	counter := 0
	for num > 0 {
		num = num / 10
		counter += 1
	}
	fmt.Println(counter)
	return counter
}

func ReverseNum(num int) int {
	ans := 0
	for num > 0 {
		digit := num % 10
		num = num / 10
		ans = ans*10 + digit
	}
	return ans
}

func CalcGcd(num int) int {
	ans := 1
	for i := 1; i <= num/2; i++ {
		if num%i == 0 && i > ans {
			ans = i
		}
	}
	fmt.Println(ans)
	return ans
}

func CheckIfArmstrongNum(num int) {
	ans := 0
	// temp := num
	for num > 0 {
		digit := num % 10
		num = num / 10
		cube := digit * digit * digit
		ans = ans + cube
	}
	fmt.Println(ans)
}
