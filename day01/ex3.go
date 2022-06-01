package main

import "fmt"

func deferfuncreturn1() int {
	i := 1
	defer func(result int) {
		result++
	}(i)
	return i
}

func deferfuncreturn2() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
	/*
		    该函数的return语句可以拆分成下面两行：
			result = i
			return
			而延迟函数的执行正是在return之前，即加入defer后的执行过程如下：
			result = i
			result++
			return
	*/

}

func foo() int {
	var i int
	defer func() {
		i++
	}()
	return 1
}

func main() {
	fmt.Println("deferfuncreturn1:", deferfuncreturn1())
	fmt.Println("deferfuncreturn2:", deferfuncreturn2())
	fmt.Println("foo:", foo())
}
