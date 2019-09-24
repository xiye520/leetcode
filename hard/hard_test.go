package hard

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	arr := []int{2, 1, 5, 6, 2, 3}
	area := largestRectangleArea(arr)
	if area != 10 {
		t.Fatal("err area:", area)
	}

	arr = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area = largestRectangleArea(arr)
	if area != 49 {
		t.Fatal("err area:", area)
	}
}

func Test_largestRectangleArea1(t *testing.T) {
	arr := []int{2, 1, 5, 6, 2, 3}
	area := largestRectangleArea1(arr)
	if area != 10 {
		t.Fatal("err area:", area)
	}

	arr = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area = largestRectangleArea1(arr)
	if area != 16 {
		t.Fatal("err area:", area)
	}
}
