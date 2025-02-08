//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
//
// Related Topics 数组 矩阵 模拟 👍 1841 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return nil
    }
    if len(matrix) == 1 && len(matrix[0]) == 1 {
        return []int{1}
    }

    up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
    res := make([]int, (down+1)*(right+1))

    i := 0
    flag := len(matrix) * len(matrix[0])
    for i < flag {

        for p := left; p < right; p++ {
            res[i] = matrix[up][p]
            i++
            if i == flag {
                return res
            }
        }

        for p := up; p < down; p++ {
            res[i] = matrix[p][right]
            i++
            if i == flag {
                return res
            }
        }

        for p := right; p > left; p-- {
            res[i] = matrix[down][p]
            i++
            if i == flag {
                return res
            }
        }

        for p := down; p > up; p-- {
            res[i] = matrix[p][left]
            i++
            if i == flag {
                return res
            }
        }

        left++
        right--
        up++
        down--

        if up == down && left == right {
            res[i] = matrix[up][left]
            i++
        }
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
