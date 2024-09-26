package majority_element

/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:

Input: nums = [3,2,3]
Output: 3

Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2

Constraints:

n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109

Follow-up: Could you solve the problem in linear time and in O(1) space?
*/

func majorityElement(nums []int) int {
	counts := make(map[int]int) // hash map to store the frequency of each element
	for _, num := range nums {
		counts[num]++
	}

	majorityElement, majorityFrequency := 0, 0

	for num, count := range counts { // iterate through the hash map to find the majority element
		if count > majorityFrequency {
			majorityElement = num
			majorityFrequency = count
		}
	}

	return majorityElement
}
