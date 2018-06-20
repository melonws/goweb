package main

import "fmt"

func main() {
	nums := []int{3,5,4,2,7,9,1,6,8}

	quickSort(nums)

	fmt.Println(nums)
}


func quickSort (array []int) {

	if len(array) <= 1 {
		return
	}
	mid, i := array[0], 1
	head, tail := 0, len(array)-1
	for head < tail {
		fmt.Println(array)
		if array[i] > mid {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
		}
	}
	array[head] = mid
	quickSort(array[:head])
	quickSort(array[head+1:])
}
