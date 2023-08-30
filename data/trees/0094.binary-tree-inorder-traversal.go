package trees

/*** 94. 二叉树的中序遍历

中序遍历的规则：
（1）中序遍历左子树
（2）访问根节点
（3）中序遍历右子树

https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
*/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorder(root, &res)
	return res
}

func inorder(tree *TreeNode, res *[]int) {
	if tree == nil {
		return
	}
	inorder(tree.Left, res)
	*res = append(*res, tree.Val)
	inorder(tree.Right, res)
}
