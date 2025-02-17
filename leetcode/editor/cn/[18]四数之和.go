//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）： 
//
// 
// 0 <= a, b, c, d < n 
// a、b、c 和 d 互不相同 
// nums[a] + nums[b] + nums[c] + nums[d] == target 
// 
//
// 你可以按 任意顺序 返回答案 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 200 
// -10⁹ <= nums[i] <= 10⁹ 
// -10⁹ <= target <= 10⁹ 
// 
//
// Related Topics 数组 双指针 排序 👍 2024 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import "sort"

func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4 {
        return nil
    }
    sort.Ints(nums)
    res := make([][]int, 0)
    for i := 0; i < len(nums)-3; i++ {
        n1 := nums[i]
        if i > 0 && n1 == nums[i-1] {
            continue
        }
        for j := i + 1; j < len(nums)-2; j++ {
            n2 := nums[j]
            if j > i+1 && n2 == nums[j-1] {
                continue
            }
            l := j + 1
            r := len(nums) - 1
            for l < r {
                n3 := nums[l]
                n4 := nums[r]
                sum := n1 + n2 + n3 + n4
                if sum < target {
                    l++
                } else if sum > target {
                    r--
                } else {
                    res = append(res, []int{n1, n2, n3, n4})
                    for l < r && n3 == nums[l+1] {
                        l++
                    }
                    for l < r && n4 == nums[r-1] {
                        r--
                    }
                    r--
                    l++
                }
            }
        }
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
