package basicrecursion

import "fmt"

func PrintNameNTimes(counter int, num int) {
	if counter >= num {
		return
	}
	fmt.Println("mogly")
	PrintNameNTimes(counter+1, num)
}

func PrintTillN(counter int, num int) {
	if counter > num {
		return
	}
	fmt.Println(counter)
	PrintTillN(counter+1, num)
}

func PrintNToOne(num int) {
	if num < 1 {
		return
	}
	fmt.Println(num)
	PrintNToOne(num - 1)
}

func PrintTillNBacktracking(counter int, num int) {
	if counter > num {
		return
	}
	PrintTillNBacktracking(counter+1, num)
	fmt.Println(counter)
}

func OneToNBacktracking(num int) {
	if num < 1 {
		return
	}
	OneToNBacktracking(num - 1)
	fmt.Println(num)
}

func NToOneBacktracking(num int, counter int) {
	if counter > num {
		return
	}
	NToOneBacktracking(num, counter+1)
	fmt.Println(counter)
}

func ParametirizedSummation(num int, counter int, sum int) {
	if counter > num {
		fmt.Println(sum)
		return
	}
	ParametirizedSummation(num, counter+1, sum+counter)
}

func ParametirizedSummationTwo(counter int, sum int) {
	if counter < 0 {
		fmt.Println(sum)
		return
	}
	ParametirizedSummationTwo(counter-1, sum+counter)
}

func FunctionalSummation(num int) int {
	if num < 0 {
		return 0
	}
	return num + FunctionalSummation(num-1)
}

func ReverseArray(arr []int, l int, r int) {
	if l == r || r < l {
		PrintArray(arr)
		return
	}
	val := arr[l]
	arr[l] = arr[r]
	arr[r] = val
	ReverseArray(arr, l+1, r-1)
}

func IsStringPalindrome(str string, l, r int) bool {
	if l > r || l == r {
		return true
	}
	if str[l] != str[r] {
		return false
	}
	return IsStringPalindrome(str, l+1, r-1)
}

func GetFactorial(num int) int {
	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	return GetFactorial(num-1) + GetFactorial(num-2)
}

func PrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
