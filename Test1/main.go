package main

import "fmt"

func main() {
	var input int
	fmt.Print("Input Number: ")
	fmt.Scan(&input)
	for i := 0; i < input; i++ {
		hasil := ((i * (i + 1)) / 2) + 1
		if i != input-1 {
			fmt.Print(hasil, "-")
		} else {
			fmt.Print(hasil)
		}
	}
}
