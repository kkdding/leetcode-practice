//给定一个由 0 和 1 组成的非空二维数组 grid ，用来表示海洋岛屿地图。
//
// 一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被
//0（代表水）包围着。
//
// 找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
//
//
//
// 示例 1：
//
//
//
//
//输入: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1
//,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0
//,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//输出: 6
//解释: 对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。
//
// 示例 2：
//
//
//输入: grid = [[0,0,0,0,0,0,0,0]]
//输出: 0
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 50
// grid[i][j] is either 0 or 1
//
//
//
//
// 注意：本题与主站 695 题相同： https://leetcode-cn.com/problems/max-area-of-island/
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 101 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		area := 1
		area += dfs(i+1, j)
		area += dfs(i-1, j)
		area += dfs(i, j+1)
		area += dfs(i, j-1)
		return area
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(i, j))
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
