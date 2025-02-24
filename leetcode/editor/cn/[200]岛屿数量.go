//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 2686 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func numIslands(grid [][]byte) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				num++
				bfs(grid, i, j)
			}
		}
	}
	return num
}

func dfs(grid [][]byte, i, j int) {
	grid[i][j] = 0
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	if j-1 >= 0 && grid[i][j-1] == 49 {
		dfs(grid, i, j-1)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
	return
}

func bfs(grid [][]byte, i, j int) {
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{i, j})
	nextQueue := make([][2]int, 0)
	for len(queue) > 0 {
        for _, v := range queue {
            i, j := v[0], v[1]
            if i-1 >= 0 && grid[i-1][j] == '1' {
                nextQueue = append(nextQueue, [][2]int{{i - 1, j}}...)
                grid[i-1][j] = '0'
            }
            if i+1 < len(grid) && grid[i+1][j] == '1' {
                nextQueue = append(nextQueue, [][2]int{{i + 1, j}}...)
                grid[i+1][j] = '0'
            }
            if j-1 >= 0 && grid[i][j-1] == '1' {
                nextQueue = append(nextQueue, [][2]int{{i, j - 1}}...)
                grid[i][j-1] = '0'
            }
            if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
                nextQueue = append(nextQueue, [][2]int{{i, j + 1}}...)
                grid[i][j+1] = '0'
            }
        }
        queue = nextQueue
        nextQueue = make([][2]int, 0)
    }
}

//leetcode submit region end(Prohibit modification and deletion)
