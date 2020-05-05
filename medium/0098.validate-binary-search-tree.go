package medium

import "math"

/*
98. 验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre int

func isValidBST2(root *TreeNode) bool {
	pre = math.MinInt64
	return InOrder(root)
}

func InOrder(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 访问左子树
	if !InOrder(root.Left) {
		return false
	}
	// 访问当前节点：如果当前节点小于等于中序遍历的前一个节点，说明不满足BST，返回 false；否则继续遍历。
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	// 访问右子树
	return InOrder(root.Right)
}

// leetcode官方
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

func isValidBST3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, right, left int) bool {
	if root == nil {
		return true
	}

	// 左子树小于当前节点，右子树大于当前节点
	if left >= root.Val || right <= root.Val {
		return false
	}

	return isBST(root.Left, left, root.Val) && isBST(root.Right, root.Val, right)
}
