package main

import "fmt"

func appendNum(num int, a []int) {
	a = append(a, num)
}

func main() {
	a := []int{1}
	appendNum(3, a)
	fmt.Println(a)
}
