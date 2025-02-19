//给定二叉树的根节点 root ，返回所有左叶子之和。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入: root = [3,9,20,null,null,15,7] 
//输出: 24 
//解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
// 
//
// 示例 2: 
//
// 
//输入: root = [1]
//输出: 0
// 
//
// 
//
// 提示: 
//
// 
// 节点数在 [1, 1000] 范围内 
// -1000 <= Node.val <= 1000 
// 
//
// 
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 758 👎 0

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

//type TreeNode struct {
//    Val   int
//    Left  *TreeNode
//    Right *TreeNode
//}

func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0
    if root.Left!=nil{
        if root.Left.Left==nil && root.Left.Right==nil{
            res += root.Left.Val
        }else{
            res += sumOfLeftLeaves(root.Left)
        }
    }
    if root.Right!=nil{
        res += sumOfLeftLeaves(root.Right)
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
