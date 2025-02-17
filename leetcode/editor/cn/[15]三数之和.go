//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
//k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//
//
// 示例 3：
//
//
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics 数组 双指针 排序 👍 7165 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "sort"

func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++ {
        n1 := nums[i]
        if n1 > 0 {
            break
        }
        if i > 0 && n1 == nums[i-1] {
            continue
        }
        l, r := i+1, len(nums)-1
        for l < r {
            n2, n3 := nums[l], nums[r]
            if n1+n2+n3 == 0 {
                res = append(res, []int{n1, n2, n3})
                for l < r && nums[l] == n2 {
                    l++
                }
                for l < r && nums[r] == n3 {
                    r--
                }
            } else if n1+n2+n3 > 0 {
                r--
            } else {
                l++
            }
        }
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
