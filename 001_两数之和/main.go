package main

import "fmt"

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]

// 使用暴力解法
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 使用hash表
func twoSum2(nums []int, target int) []int {
	// key:数值,value:下标
	nummap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if j, ok := nummap[target-nums[i]]; ok {
			return []int{j, i}
		} else {
			nummap[nums[i]] = i
		}
	}
	return nil
}

func main() {
	nums := []int{2, 6, 11, 15, 7}
	target := 9
	result := twoSum1(nums, target)
	fmt.Println(result)
	result2 := twoSum2(nums, target)
	fmt.Println(result2)
}
