package main

import "fmt"

func get() []byte {
	raw := make([]byte, 1000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	return raw[:3]
}

func main() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0])

	sli1 := make([]int, 5, 10)
	sli2 := sli1[:3]
	fmt.Println(len(sli2), cap(sli2))
}
