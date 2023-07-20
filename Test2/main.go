package main

import (
	"fmt"
)

func main() {
	var pemain, permainan int
	fmt.Print("Jumlah pemain: ")
	fmt.Scan(&pemain)

	skor := make([]int, pemain)
	fmt.Print("Skor: ")
	for i := 0; i < pemain; i++ {
		fmt.Scan(&skor[i])
	}

	fmt.Print("Jumlah permainan: ")
	fmt.Scan(&permainan)

	skorPermainan := make([]int, permainan)
	fmt.Print("Skor permainan: ")
	for i := 0; i < permainan; i++ {
		fmt.Scan(&skorPermainan[i])
	}

	// var pemain int = 7
	// var permainan int = 4
	// skor := []int{100, 100, 50, 40, 40, 20, 10}
	// skorPermainan := []int{5, 25, 50, 120}

	var output []int

	var arrSkor []int
	for _, list := range skor {
		if !contains(arrSkor, list) {
			arrSkor = append(arrSkor, list)
		}
	}

	var arrSP []int
	for _, list := range skorPermainan {
		if !contains(arrSP, list) {
			arrSP = append(arrSP, list)
		}
	}

	for _, list := range arrSP {
		for i := len(arrSkor) - 1; i >= 0; i-- {
			if list == arrSkor[i] {
				output = append(output, i+1)
				// fmt.Println(list)
				break
			} else if list < arrSkor[i] {
				output = append(output, i+2)
				// fmt.Println(list)
				break
			}
			if i == 0 && list > arrSkor[i] {
				output = append(output, i+1)
				// fmt.Println(list)
			}
		}
	}

	// fmt.Println()
	fmt.Print("Output: ")
	for _, list := range output {
		fmt.Print(list, " ")
	}

}

func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}