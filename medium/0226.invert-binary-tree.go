package medium

/*
* 226. 翻转二叉树
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

https://leetcode.cn/problems/invert-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150
*/

// 方法一：递归
// 这是一道很经典的二叉树问题。显然，我们从根节点开始，递归地对树进行遍历，并从叶子节点先开始翻转。
// 如果当前遍历到的节点 root\textit{root}root 的左右两棵子树都已经翻转，
// 那么我们只需要交换两棵子树的位置，即可完成以 root\textit{root}root 为根节点的整棵子树的翻转。
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Right = left
	root.Left = right

	return root
}

// BFS 写法 用层序遍历的方式去遍历二叉树。
// 根节点先入列，然后出列，出列就 “做事”，交换它的左右子节点（左右子树）。 并让左右子节点入列，往后，这些子节点出列，也被翻转。
// 直到队列为空，就遍历完所有的节点，翻转了所有子树。
// 解决问题的代码放在节点出列时。
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		cur.Left, cur.Right = cur.Right, cur.Left

		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return root
}
