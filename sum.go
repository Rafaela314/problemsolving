package main

/*Two Sum problem:
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	Input: nums = [2,4,5,11,15], target = 9
	Output: [1,2]
	Output: Because nums[1] + nums[2] == 9, we return [1, 2].
  Target =13
  Output =[0,3]
--exception = 1-1-
*/

func TwoSum(nums []int, target int) [2]int {

	var result [2]int

	if len(nums) > 0 {
		for k, v := range nums {
			for i := k + 1; i < len(nums); i++ {
				if nums[i]+v == target {
					result[0] = k
					result[1] = i
					break
				}
			}
		}
	}

	if result[1] == 0 {
		return [2]int{-1, -1}
	}

	return result
}
