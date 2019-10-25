package easy

import "container/list"

/*
104. 二叉树的最大深度

给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1
	return max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	i := depth(root)
	return i
}

func depth(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	var maxDeep int = 0
	for {
		len := queue.Len()
		if len == 0 {
			break
		}
		for i := 0; i < len; i++ {
			front := queue.Front()
			node := (front.Value).(*TreeNode)
			queue.Remove(front)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		maxDeep++
	}
	return maxDeep
}

func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode       // 辅助队列
	queue = append(queue, root) // 根节点入队
	depth := 0                  // 初始化深度为0

	for len(queue) > 0 { // 当队列不为空时，循环以下操作
		size := len(queue)
		for i := 0; i < size; i++ { // 遍历当前层级中的节点
			s := queue[0]      // 获取队首元素
			queue = queue[1:]  // 队首元素出队
			if s.Left != nil { // 如果左子树不为空，左子树入队
				queue = append(queue, s.Left)
			}
			if s.Right != nil { // 如果右子树不为空，右子树入队
				queue = append(queue, s.Right)
			}
		}
		depth++ // for循环结束后这一层级节点已经访问结束，深度加+1
	}
	return depth
}
