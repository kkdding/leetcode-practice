//给你一棵二叉树的根节点，返回该树的 直径 。
//
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
//
// 两节点之间路径的 长度 由它们之间边数表示。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,4,5]
//输出：3
//解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
//
//
// 示例 2：
//
//
//输入：root = [1,2]
//输出：1
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 10⁴] 内
// -100 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 1635 👎 0

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
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var dfs func( *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		ld := dfs(node.Left) + 1
		rd := dfs(node.Right) + 1
		res = max(res,ld + rd)
		return max(ld, rd)
	}
	dfs(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
