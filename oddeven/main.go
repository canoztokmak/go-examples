package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, el := range slice {
		if el%2 == 1 {
			fmt.Println(el, "is odd")
		} else {
			fmt.Println(el, "is even")
		}
	}
}
