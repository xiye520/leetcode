package trees

/** 145. 二叉树的后序遍历

后序遍历二叉树的规则：
（1）后序遍历左子树
（2）后序遍历右子树
（3）访问根节点

https://leetcode.cn/problems/binary-tree-postorder-traversal/submissions/
*/
func postorderTraversal(root *TreeNode) []int {
	var res []int
	postorder(root, &res)
	return res
}

func postorder(tree *TreeNode, res *[]int) {
	if tree == nil {
		return
	}
	postorder(tree.Left, res)
	postorder(tree.Right, res)
	*res = append(*res, tree.Val)
}
