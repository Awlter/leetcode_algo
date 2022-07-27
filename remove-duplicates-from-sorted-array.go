package main

import "fmt"

func main() {
	a := []int{1,1,2}
	fmt.Println(removeDuplicates(a))
}

func removeDuplicates(nums []int) int {
	n := 1
	for i, num := range nums {
			if i == 0 {
					continue
			}
			if num != nums[n-1] {
					nums[n] = num
					n++
			}
	}
	fmt.Println(nums)
	return n
}
