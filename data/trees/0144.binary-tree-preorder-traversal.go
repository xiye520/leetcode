package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/** 144. 二叉树的前序遍历

前序遍历的规则：
（1）访问根节点
（2）前序遍历左子树
（3）前序遍历右子树

https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
*/
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	preorder(root, &res)
	return res
}

func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)
}
