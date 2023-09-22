package easy

import "math"

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var arr []int
	midOrder(root, &arr)

	ans := math.MaxInt32
	now := arr[0]
	for i := 1; i < len(arr); i++ {
		now = arr[i] - arr[i-1]
		if now < ans {
			ans = now
		}
	}
	return ans
}

func midOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	midOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	midOrder(root.Right, arr)
}

func getMinimumDifference2(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
