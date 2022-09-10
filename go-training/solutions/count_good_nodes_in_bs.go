package solutions

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	count := 0
	countGN(root, &count, math.MinInt32)

	return count
}

func countGN(root *TreeNode, count *int, maxValue int) {
	if root == nil {
		return
	}

	if v := root.Val; v >= maxValue {
		*count++
		maxValue = v
	}

	countGN(root.Left, count, maxValue)
	countGN(root.Right, count, maxValue)
}
