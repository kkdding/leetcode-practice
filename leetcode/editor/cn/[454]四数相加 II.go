//给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
//
//
// 0 <= i, j, k, l < n
// nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
//
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
//输出：2
//解释：
//两个元组如下：
//1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1)
// + 2 = 0
//2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1)
// + 0 = 0
//
//
// 示例 2：
//
//
//输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
//输出：1
//
//
//
//
// 提示：
//
//
// n == nums1.length
// n == nums2.length
// n == nums3.length
// n == nums4.length
// 1 <= n <= 200
// -2²⁸ <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2²⁸
//
//
// Related Topics 数组 哈希表 👍 1073 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	set := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			set[n1+n2]++
		}
	}

	res := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			if _, ok := set[-n3-n4]; ok {
				res += set[-n3-n4]
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
