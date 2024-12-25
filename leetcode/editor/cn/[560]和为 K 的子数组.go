//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//
// 子数组是数组中元素的连续非空序列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,1], k = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3], k = 3
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁴
// -1000 <= nums[i] <= 1000
// -10⁷ <= k <= 10⁷
//
//
// Related Topics 数组 哈希表 前缀和 👍 2566 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func subarraySum(nums []int, k int) int {
	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = 0
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	sum := 0
	for i := 0; i < len(prefixSum); i++ {
		for j := i + 1; j < len(prefixSum); j++ {
			if prefixSum[j]-prefixSum[i] == k {
				sum++
			}
		}
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
