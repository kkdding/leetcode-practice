//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。 
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
// 
//
// 示例 2： 
//
// 
//
// 
//输入： heights = [2,4]
//输出： 4 
//
// 
//
// 提示： 
//
// 
// 1 <= heights.length <=10⁵ 
// 0 <= heights[i] <= 10⁴ 
// 
//
// Related Topics 栈 数组 单调栈 👍 2873 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func largestRectangleArea(heights []int) (ans int) {
    left := make([]int, len(heights))
    for i := 0; i < len(heights); i++ {
        left[i] = len(heights)
    }
    for i := 0; i < len(heights); i++ {
        for j := i + 1; j < len(heights); j++ {
            if heights[j] < heights[i] {
                left[i] = j
                break
            }
        }
    }

    right := make([]int, len(heights))
    for i := 0; i < len(heights); i++ {
        right[i] = -1
    }
    for i := len(heights) - 1; i >= 0; i-- {
        for j := i - 1; j >= 0; j-- {
            if heights[j] < heights[i] {
                right[i] = j
                break
            }
        }
    }

    area := 0
    for i := 0; i < len(heights); i++ {
        area = max(area, (left[i]-right[i]-1)*heights[i])
    }
    return area
}

//leetcode submit region end(Prohibit modification and deletion)
