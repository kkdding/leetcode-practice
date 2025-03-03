//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2647 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func findKthLargest(nums []int, k int) int {
	quickSort(nums)
	return nums[len(nums)-k]
}

func countSort(nums []int) {
	bucket := make([]int, len(nums)+1)
	for _, v := range nums {
		bucket[v]++
	}

	p := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			nums[p] = i
			bucket[i]--
			p++
		}
	}
}

func quickSort(nums []int) {
	n := len(nums)
	if n == 1 || n == 0 {
		return
	}
	ptr := 0
	for i, num := range nums[:n-1] {
		if num < nums[n-1] {
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
	nums[n-1], nums[ptr] = nums[ptr], nums[n-1]
	quickSort(nums[:ptr])
	quickSort(nums[ptr+1:])
}

//leetcode submit region end(Prohibit modification and deletion)
