package algo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue []*TreeNode
	var rs [][]int
	queue = append(queue, root)

	for len(queue) != 0 {
		n := len(queue)
		list := []int{}
		var curr *TreeNode

		for i := 0; i < n; i++ {
			curr = queue[i]
			list = append(list, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		rs = append(rs, list)
		queue = queue[n:]
	}

	return rs
}
