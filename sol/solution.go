package sol

import "sort"

func subsetsWithDup(nums []int) [][]int {
	// sort nums
	sort.Ints(nums)
	result := [][]int{}
	subset := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			// match the latest
			temp := make([]int, len(subset))
			copy(temp, subset)
			result = append(result, temp)
			return
		}
		// choose nums[i]
		subset = append(subset, nums[i])
		dfs(i + 1)
		subset = subset[:len(subset)-1]

		// not choose nums[i], skip the duplicate
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		dfs(i + 1)
	}
	dfs(0)
	return result
}
