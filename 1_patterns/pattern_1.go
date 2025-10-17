package main

import "fmt"

func patternOne(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	patternOne(5)
}
