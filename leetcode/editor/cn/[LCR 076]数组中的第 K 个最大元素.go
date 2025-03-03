//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,1,5,6,4], k = 2
//输出：5
//
//
// 示例 2：
//
//
//输入：nums = [3,2,3,1,2,4,5,5,6], k = 4
//输出：4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
//
// 注意：本题与主站 215 题相同： https://leetcode-cn.com/problems/kth-largest-element-in-an-
//array/
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 104 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
    "container/heap"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func heapSort(nums []int) []int {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}

	res := make([]int, 0)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

func findKthLargest(nums []int, k int) int {
	nums = heapSort(nums)
	return nums[len(nums)-k]
}

//leetcode submit region end(Prohibit modification and deletion)
