//给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历） 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：[[15,7],[9,20],[3]]
// 
//
// 示例 2： 
//
// 
//输入：root = [1]
//输出：[[1]]
// 
//
// 示例 3： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 2000] 内 
// -1000 <= Node.val <= 1000 
// 
//
// Related Topics 树 广度优先搜索 二叉树 👍 827 👎 0

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

import "slices"

//type TreeNode struct {
//    Val   int
//    Left  *TreeNode
//    Right *TreeNode
//}

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    res :=make([][]int, 0)
    curQueue := []*TreeNode{root}

    for len(curQueue) >0{
        val := make([]int, 0)
        nextQueue := make([]*TreeNode, 0)
        for _, node := range curQueue{
            if node.Left!=nil{
                nextQueue = append(nextQueue,node.Left)
            }
            if node.Right!=nil{
                nextQueue = append(nextQueue,node.Right)
            }
            val = append(val,node.Val)
        }
        res = append(res,val)
        curQueue = nextQueue
    }
    slices.Reverse(res)
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
