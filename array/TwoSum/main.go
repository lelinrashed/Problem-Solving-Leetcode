package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4}
	target := 7
	fmt.Println("result", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	result := []int{}

	for index, value := range nums {
		if checkValueExists(hashMap, target-value) {
			result = append(result, returnKey(hashMap, target-value))
			result = append(result, index)
		}
		hashMap[index] = value
	}
	return result
}

func checkValueExists(argss map[int]int, value int) bool {
	for _, v := range argss {
		if value == v {
			return true
		}
	}
	return false
}

func returnKey(argss map[int]int, value int) (key int) {
	for i, v := range argss {
		if value == v {
			key = i
			return
		}
	}
	return
}
