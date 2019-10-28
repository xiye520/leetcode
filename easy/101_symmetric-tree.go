package easy

/*
101. 对称二叉树

给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:
如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/
func isSymmetric(root *TreeNode) bool {
	/*思路
	递归结束条件：

	都为空指针则返回 true
	只有一个为空则返回 false
	递归过程：

	判断两个指针当前节点值是否相等
	判断 A 的右子树与 B 的左子树是否对称
	判断 A 的左子树与 B 的右子树是否对称
	短路：

	在递归判断过程中存在短路现象，也就是做 与 操作时，如果前面的值返回 false 则后面的不再进行计算
	*/

	return isMirror(root, root)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func isMirror2(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && isMirror2(t1.Right, t2.Left) && isMirror2(t2.Left, t1.Right)
}

func isSymmetric3(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
