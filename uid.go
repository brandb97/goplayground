package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	a := [3]int(s)

	fmt.Println(s, a)
}
