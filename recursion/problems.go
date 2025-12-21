package recursion

import "fmt"

func RecursionProblems() {
	res := make([]string, 0)
	generateBinaryStrings(4, "", &res)
	fmt.Println(res)
}

func generateBinaryStrings(n int, str string, res *[]string) {
	if n == len(str) {
		*res = append(*res, str)
		return
	}
	generateBinaryStrings(n, str+"0", res)
	generateBinaryStrings(n, str+"1", res)
}
