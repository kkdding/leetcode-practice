//给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 20
//
//
// Related Topics 数组 矩阵 模拟 👍 1402 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	left, right, up, down := 0, n-1, 0, n-1

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	flag := n * n
	for num <= flag {
		for p := left; p < right; p++ {
			matrix[up][p] = num
			num++
		}

		for p := up; p < down; p++ {
			matrix[p][right] = num
			num++
		}

		for p := right; p > left; p-- {
			matrix[down][p] = num
			num++
		}

		for p := down; p > up; p-- {
			matrix[p][left] = num
			num++
		}
		left++
		right--
		up++
		down--

		if left == right {
			matrix[left][up] = num
			num++
		}
	}
	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)
