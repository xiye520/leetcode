package medium

func ExampleReorderList() {
	head := generateListNode()
	printList(head)
	printList(reorderList(head))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 1 -> 5 -> 2 -> 4 -> 3
}
