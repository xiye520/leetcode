package hard

/*
84.柱状图中最大的矩形

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例:

输入: [2,1,5,6,2,3]
输出: 10
*/
func largestRectangleArea(heights []int) int {
	var area int

	return area
}

/*
o(n) 算法，使用栈的方式
只有当前栈为空或者当前高度大于栈顶值的高度，才能入栈
如果当前值等于栈顶值的高度，不做任何处理
如果当前值小于栈顶值的高度，那么将当前值弹出站，计算弹出的值对应的最大面积，更新最大面积，如果并记录弹出的开始位置；如果栈顶的值等于当前值，不做处理；如果当前值大于栈顶值或者栈顶为空，将当前值push到栈里面，并且记录起始位置为最后弹出的值的起始位置
遍历完毕，如果栈不为空，将栈元素弹出，并以此记录每个元素对应的最大面积

*/
type recordStruct struct {
	Start int
	Value int
}

func largestRectangleArea1(heights []int) int {
	maxArea := 0
	stackArray := make([]recordStruct, len(heights))
	topPos := 0
	for i, v := range heights {
		flagPush := false
		start := i
		if topPos == 0 || v > stackArray[topPos-1].Value {
			flagPush = true
		} else {
			//equal不做任何处理
			for topPos > 0 && v < stackArray[topPos-1].Value {
				tmpArea := stackArray[topPos-1].Value * (i - stackArray[topPos-1].Start)
				if maxArea < tmpArea {
					maxArea = tmpArea
				}
				start = stackArray[topPos-1].Start
				topPos--
			}
			if topPos == 0 || v > stackArray[topPos-1].Value {
				flagPush = true
			}
		}
		if flagPush {
			stackArray[topPos] = recordStruct{
				Start: start,
				Value: v,
			}
			topPos++
		}
	}
	for topPos > 0 {
		tmpArea := stackArray[topPos-1].Value * (len(heights) - stackArray[topPos-1].Start)
		if maxArea < tmpArea {
			maxArea = tmpArea
		}
		topPos--
	}
	return maxArea
}
