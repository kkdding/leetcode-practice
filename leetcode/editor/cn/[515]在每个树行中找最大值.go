//给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。 
//
// 
//
// 示例1： 
//
// 
//
// 
//输入: root = [1,3,2,5,3,null,9]
//输出: [1,3,9]
// 
//
// 示例2： 
//
// 
//输入: root = [1,2,3]
//输出: [1,3]
// 
//
// 
//
// 提示： 
//
// 
// 二叉树的节点个数的范围是 [0,10⁴] 
// 
// -2³¹ <= Node.val <= 2³¹ - 1 
// 
//
// 
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 389 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "math"

//type TreeNode struct {
//    Val   int
//    Left  *TreeNode
//    Right *TreeNode
//}

func largestValues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    res := make([]int, 0)
    curQueue := []*TreeNode{root}

    for len(curQueue) > 0 {
        nextQueue := make([]*TreeNode, 0)
        maxVal := math.MinInt
        for _, node := range curQueue {
            if node.Left != nil {
                nextQueue = append(nextQueue, node.Left)
            }
            if node.Right != nil {
                nextQueue = append(nextQueue, node.Right)
            }
            maxVal = max(maxVal, node.Val)
        }
        res = append(res, maxVal)
        curQueue = nextQueue
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
