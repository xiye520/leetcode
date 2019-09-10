package demo

import "fmt"

// 7.重建二叉树

type treeNode2 struct {
	value int
	left  *treeNode2
	right *treeNode2
}

func CreateTree2(preorder, inorder []int) (ptNode *treeNode2) {

	var (
		leftChiledSize int
	)

	if len(preorder) != len(inorder) {
		fmt.Println("error")
		// debug
		panic("err1")
	}
	ptNode = &treeNode2{
		left:  nil,
		right: nil,
		value: preorder[0],
	}
	// 递归退出条件
	if len(preorder) == 1 {
		return ptNode
	}

	leftChiledSize = GetLeftChiledSize2(preorder[0], inorder)
	ptNode.left = CreateTree2(preorder[1:1+leftChiledSize], inorder[0:leftChiledSize])
	ptNode.right = CreateTree2(preorder[1+leftChiledSize:], inorder[leftChiledSize+1:])
	return ptNode
}

//获取左子树节点数
func GetLeftChiledSize2(root int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if root == inorder[i] {
			return i
		}
	}
	// debug
	panic("err2")
}

// 后续遍历
func PostOrderTraversal(ptree *treeNode2) {
	if ptree.left == nil && ptree.right == nil {
		fmt.Println(ptree.value)
		return
	}

	if ptree.left != nil {
		PostOrderTraversal(ptree.left)
	}
	if ptree.right != nil {
		PostOrderTraversal(ptree.right)
	}
	fmt.Println(ptree.value)
}

func topic7() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	treeRoot := CreateTree2(preorder, inorder)
	PostOrderTraversal(treeRoot)
}
