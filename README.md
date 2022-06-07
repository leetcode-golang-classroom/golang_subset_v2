# golang_subset_v2

Given an integer array `nums` that may contain duplicates, return *all possible subsets (the power set)*.

The solution set **must not** contain duplicate subsets. Return the solution in **any order**.

## Examples

**Example 1:**

```
Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

```

**Example 2:**

```
Input: nums = [0]
Output: [[],[0]]

```

**Constraints:**

- `1 <= nums.length <= 10`
- `-10 <= nums[i] <= 10`

## 解析

給定一個會有重複值的整數陣列 nums

要求實做出一個演算法找出所有不重複的子陣列集合

這題跟 [78. Subsets](https://www.notion.so/78-Subsets-22b988f1407149f5a2eecd33094f60a7) 類似

只差在會有重複值的元素

所以要如何避免窮舉出重複值就是一個關鍵

首先觀察決策樹，如下圖：

![](https://i.imgur.com/Y2I77Wp.png)

可以發現紫色的部份是重複的

注意到如果把紫色的 subTree 去掉，可以發現從重複元素2開始可以分為 至少出現一次的2 與都不出現的2

要來程式碼的部份要來做到這件事情

首先要先做排序

然後每次遇到重複的元素就直接跳到下一個元素再繼續做

這樣就可以避免重複

而因為原本窮舉法的時間複雜度為 O($n*2^n$）

所以排序的時間複雜度 O($nlogn$) 不影響

## 程式碼
```go
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

```
## 困難點

1. 理解決策樹做窮舉的作法
2. 想出如何避免重複選取的作法

## Solve Point

- [x]  Understand what problem to solve
- [x]  Analysis Complexity