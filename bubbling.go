package main

import "fmt"

var arr  = []int{12,14,78,56,98,34,7,4,1,1}

//冒泡
func main()  {
	for i := range arr {
		for j := range arr {
			if arr[i] < arr[j] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
	fmt.Println(arr)
}
