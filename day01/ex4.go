package main

import "fmt"

func main() {

	sli := make([]int, 5, 10)
	//fmt.Println(sli)

	fmt.Printf("切片cli长度和容量:%d,%d\n", len(sli), cap(sli))

	fmt.Println(sli)
	newsli := sli[:cap(sli)]
	fmt.Println(newsli)

	var x = []int{2, 3, 5, 7, 11}
	fmt.Printf("切片x长度和容量: %d,%d\n", len(x), cap(x))

	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5] // a[low:high:max]  max-low 表示容量  high-low 表示长度
	fmt.Printf("切片t长度和容量：%d,%v\n", len(t), cap(t))
	fmt.Println(t)
}
