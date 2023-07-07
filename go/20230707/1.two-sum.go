/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// deep copy
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] == target {
			// find index
			for k := 0; k < len(numsCopy); k++ {
				if numsCopy[k] == nums[i] {
					i = k
					break
				}
			}
			for k := len(numsCopy) - 1; k >= 0; k-- {
				if numsCopy[k] == nums[j] {
					j = k
					break
				}
			}
			return []int{i, j}

		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

// @lc code=end
