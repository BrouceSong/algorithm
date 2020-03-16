package main

import (
	"fmt"
	"math"
)

var data = []int{1, 1, 4, 7, 12, 14, 34, 56, 78, 98}

func main() {
	fmt.Println(dichotomy(data, 7))
}

func dichotomy(arr []int, num int) bool {
	k := int(math.Floor(float64(len(arr)/2))) - 1
	tmp := make([]int, 0, len(arr))
	if len(arr) <= 1 && arr[0] != num {
		return false
	}
	if len(arr) <= 1 && arr[0] == num {
		return true
	}
	if arr[k] == num {
		return true
	} else if arr[k] > num {
		tmp = arr[0:k]
	} else {
		tmp = arr[k+1:]
	}
	//fmt.Println(k, arr[k], tmp)
	return dichotomy(tmp, num)
}
