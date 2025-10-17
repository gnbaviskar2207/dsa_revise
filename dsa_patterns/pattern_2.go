package dsapatterns

import "fmt"

func ComputePatternsTwo(n int) {
	for i := 0; i < n; i++ {
		num := 1
		for j := 0; j <= i; j++ {
			fmt.Print(num)
			num += 1
		}
		fmt.Println()
	}
}

func ComputePatternsThree(n int) {
	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(num)

		}
		num += 1
		fmt.Println()
	}
}

func ComputePatternsFour(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func ComputePatternsFive(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i+1; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func ComputePatternsSix(n int) {
	for row := 1; row <= n; row++ {
		space := n - row
		for i := 1; i <= space; i++ {
			fmt.Print(" ")
		}

		for i := 1; i <= row*2-1; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func ComputePatternsSeven(n int) {
	for row := 0; row <= n; row++ {
		for i := 0; i < row; i++ {
			fmt.Print(" ")
		}

		stars := ((n - row) * 2) - 1
		for i := 0; i < stars; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func patternOne(rows int, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
