//给定一个二维数组 array，请返回「螺旋遍历」该数组的结果。 
//
// 螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，直到提取完所有元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：array = [[1,2,3],[8,9,4],[7,6,5]]
//输出：[1,2,3,4,5,6,7,8,9]
// 
//
// 示例 2： 
//
// 
//输入：array  = [[1,2,3,4],[12,13,14,5],[11,16,15,6],[10,9,8,7]]
//输出：[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16]
// 
//
// 
//
// 限制： 
//
// 
// 0 <= array.length <= 100 
// 0 <= array[i].length <= 100 
// 
//
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/ 
//
// 
//
// Related Topics 数组 矩阵 模拟 👍 609 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func spiralArray(array [][]int) []int {

    if len(array) == 0 {
        return nil
    }
    if len(array) == 1 && len(array[0]) == 1 {
        return []int{1}
    }

    up, down, left, right := 0, len(array)-1, 0, len(array[0])-1
    res := make([]int, (down+1)*(right+1))

    i := 0
    flag := len(array) * len(array[0])
    for i < flag {

        for p := left; p < right; p++ {
            res[i] = array[up][p]
            i++
            if i == flag {
                return res
            }
        }

        for p := up; p < down; p++ {
            res[i] = array[p][right]
            i++
            if i == flag {
                return res
            }
        }

        for p := right; p > left; p-- {
            res[i] = array[down][p]
            i++
            if i == flag {
                return res
            }
        }

        for p := down; p > up; p-- {
            res[i] = array[p][left]
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
            res[i] = array[up][left]
            i++
        }
    }
    return res

}

//leetcode submit region end(Prohibit modification and deletion)
