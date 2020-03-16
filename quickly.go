package main

import "fmt"

//需要排序的数组
var arr = []int{12, 14, 78, 56, 98, 34, 7, 4, 1, 1}

func main() {
	fmt.Println(quick(arr))
}

func quick(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	key := 0
	tmp := arr[key]
	arr1 := make([]int, 0, len)
	arr2 := make([]int, 0, len)
	for i := range arr {
		if i == key {
			continue
		}
		if arr[i] < tmp {
			arr1 = append(arr1, arr[i])
		} else {
			arr2 = append(arr2, arr[i])
		}
	}
	return append(append(quick(arr1), tmp), (quick(arr2))...)
}
