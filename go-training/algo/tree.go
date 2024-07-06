package algo

// https://leetcode.com/problems/binary-tree-inorder-traversal/
// Using resursive
func inorderTraversal(root *TreeNode) []int {
	rs := []int{}

	travel(root, &rs)

	return rs
}

func travel(root *TreeNode, rs *[]int) {
	if root == nil {
		return
	}

	travel(root.Left, rs)
	travel(root.Right, rs)

	*rs = append(*rs, root.Val)
}

// Using stack instead of recursive
func inorderTraversalC2(root *TreeNode) []int {
	var rs []int
	var stack []*TreeNode
	var curr *TreeNode = root

	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rs = append(rs, curr.Val)
		curr = curr.Right
	}

	return rs
}

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return ""
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	return nil
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func buildTree(preorder []int, inorder []int) *TreeNode {
	pLen := len(preorder)
	iLen := len(inorder)

	if pLen == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}

	if len(preorder) == 1 {
		return root
	}

	rootIndex := 0
	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
		}
	}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:pLen], inorder[rootIndex+1:iLen])

	return root
}

// Public function
func BuildTree(preorder []int, inorder []int) *TreeNode {
	return buildTree(preorder, inorder)
}

// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/description/
func longestZigZag(root *TreeNode) int {
	var rs int
	var zigZagMove func(node *TreeNode, direction int, currLen int)

	zigZagMove = func(node *TreeNode, direction, currLen int) {
		if node == nil {
			return
		}

		rs = max(rs, currLen)

		if direction == -1 {
			// Move right => curLen++
			zigZagMove(node.Right, 1, currLen+1)
			// Move left => currLen = 1
			zigZagMove(node.Left, -1, 1)

		} else {
			// Move left => curLen++
			zigZagMove(node.Left, -1, currLen+1)
			// Move right => currLen = 1
			zigZagMove(node.Right, 1, 1)
		}
	}

	zigZagMove(root.Left, -1, 1)
	zigZagMove(root.Right, 1, 1)

	return rs
}

// https://leetcode.com/problems/recover-binary-search-tree/description/
func recoverTree(root *TreeNode)  {
	var prev, first, second *TreeNode
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		if prev != nil {
			if first == nil && prev.Val >= node.Val {first = prev}
			if first != nil && prev.Val >= node.Val {second = node} 
		}
		prev = node
		dfs(node.Right)
	}

	dfs(root)

	first.Val, second.Val = second.Val, first.Val
}

// https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
func maxPathSum(root *TreeNode) int {
	return -1
}
