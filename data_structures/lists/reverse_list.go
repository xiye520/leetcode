package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func printNode(head *Node) {
	for head != nil {
		//fmt.Print(head.Value, "\t")
		fmt.Println(head)
		head = head.Next
	}
	fmt.Println()
}

func reverseNode(head *Node) *Node {
	//  先声明两个变量
	//  前一个节点
	var preNode *Node
	preNode = nil
	//  后一个节点
	nextNode := new(Node)
	nextNode = nil
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.Next
		//  将头节点指向前一个节点
		head.Next = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	return preNode
}

func reverseList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *Node
	cur := head
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func main() {
	node1 := new(Node)
	node1.Value = 1
	node2 := new(Node)
	node2.Value = 2
	node3 := new(Node)
	node3.Value = 3
	node4 := new(Node)
	node4.Value = 4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	printNode(node1)

	//head := reverseNode(node1)
	head := reverseList(node1)
	printNode(head)
}
