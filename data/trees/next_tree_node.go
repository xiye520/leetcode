package trees

import "fmt"

// 8.二叉树的下一个结点

type treeNode struct {
	value  int
	parent *treeNode
	left   *treeNode
	right  *treeNode
}

func CreateTree(preorder, inorder []int) (ptNode *treeNode) {

	var (
		leftChiledSize int
	)

	if len(preorder) != len(inorder) {
		fmt.Println("error")
		// debug
		panic("err1")
	}
	ptNode = &treeNode{
		parent: nil,
		left:   nil,
		right:  nil,
		value:  preorder[0],
	}
	// 递归退出条件
	if len(preorder) == 1 {
		return ptNode
	}

	leftChiledSize = GetLeftChiledSize(preorder[0], inorder)
	ptNode.left = CreateTree(preorder[1:1+leftChiledSize], inorder[0:leftChiledSize])
	ptNode.left.parent = ptNode
	ptNode.right = CreateTree(preorder[1+leftChiledSize:], inorder[leftChiledSize+1:])
	ptNode.right.parent = ptNode
	return ptNode
}

func GetLeftChiledSize(root int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if root == inorder[i] {
			return i
		}
	}
	// debug
	panic("err2")
}

// 前序遍历获取节点
func PreOrderGetNode(ptree *treeNode, target int) (res *treeNode) {
	if ptree.value == target {
		return ptree
	}
	if ptree.left != nil {
		if res = PreOrderGetNode(ptree.left, target); res != nil {
			return res
		}
	}
	if ptree.right != nil {
		if res = PreOrderGetNode(ptree.right, target); res != nil {
			return res
		}
	}

	// debug
	// fmt.Println("can not find")

	return nil
}

//获取中序遍历下的下一个节点
func InOrderNext(pt *treeNode) (res *treeNode) {
	// 有右子树 -> 取右树最左节点
	if pt.right != nil {
		circle := pt.right
		for circle.left != nil {
			circle = circle.left
		}
		return circle
	}
	// 无右子树 -> 向上找第一个左链接指向的树包含该节点的祖先节点。
	circle := pt
	for circle.parent != nil {
		if circle.parent.left == circle {
			return circle.parent
		}
		circle = circle.parent
	}

	// debug
	fmt.Println("cannot find")

	return nil
}

func topic8() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	// 生成树
	rootTree := CreateTree(preorder, inorder)

	// 前序遍历获取值为15的节点
	target := PreOrderGetNode(rootTree, 7)

	res := InOrderNext(target)
	if res == nil {
		fmt.Println("没有下一个节点")
	} else {
		fmt.Println(res.value)
	}
}
