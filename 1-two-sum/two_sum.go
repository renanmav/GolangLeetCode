package two_sum

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.

Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
*/

// slow
func twoSumNaive(nums []int, target int) []int {
	sumMap := make(map[int][]int)

	for i := 0; i < len(nums); i++ { // O(n)
		for j := 0; j < len(nums); j++ { // O(n) -> O(n*n) -> O(n^2)
			if i == j {
				continue
			}
			sumResult := nums[i] + nums[j]
			sumMap[sumResult] = []int{j, i}
		}
	}

	return sumMap[target]
}

func twoSum(nums []int, target int) []int {
	numToIndexMap := make(map[int]int)

	for i, num := range nums { // O(n)
		diff := target - num

		if j, found := numToIndexMap[diff]; found {
			return []int{j, i}
		}

		numToIndexMap[num] = i
	}

	return nil
}
