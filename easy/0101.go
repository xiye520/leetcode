package easy

/** 101. 对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。

https://leetcode.cn/problems/symmetric-tree/description/?envType=study-plan-v2&envId=top-interview-150
*/

func isSymmetricTree(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {

	if right == nil && left == nil {
		return true
	}
	if right == nil || left == nil {
		return false
	}

	return right.Val == left.Val && check(right.Left, left.Right) && check(right.Right, left.Left)
}

func isSymmetric2(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}
