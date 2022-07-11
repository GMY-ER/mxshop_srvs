package main

import "fmt"

func main() {
	arr := []int{}
	var arr1 []string
	arr2 := make([]int, 0)
	fmt.Println(arr, arr1, arr2)
	fmt.Printf("%T,%T,%T\n", arr, arr1, arr2)
	fmt.Printf("%d,%d,%d", len(arr), len(arr1), len(arr2))
	fmt.Println(arr2 == nil)
}
