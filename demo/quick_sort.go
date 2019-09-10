package demo

// 2.快排

// 左闭右闭[a,b]
func QuickSort(list []int) {
	var (
		midValue   int
		head, tail int
	)
	if len(list) <= 1 {
		return
	}
	head = 1
	midValue = list[0]
	tail = len(list) - 1
	for head < tail {
		for list[head] < midValue {
			head++
			if head >= tail {
				goto END
			}
		}
		for list[tail] > midValue {
			tail--
			if head >= tail {
				goto END
			}
		}
		list[head], list[tail] = list[tail], list[head]
		head++
		tail--
	}
END:
	if list[head] > midValue {
		list[0], list[head-1] = list[head-1], list[0]

		QuickSort(list[0 : head-1])
		QuickSort(list[head:])
	} else {
		list[0], list[head] = list[head], list[0]

		QuickSort(list[0:head])
		QuickSort(list[head+1:])
	}
}
