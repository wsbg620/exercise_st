package main

import "fmt"

func main() {
	data := []int{1, 2, 3}

	for _, v := range data {
		v *= 10
	}

	for i, _ := range data {
		data[i] = data[i] * 10
	}
	fmt.Println("data:", data)
}
