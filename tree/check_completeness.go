package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	var queue []*TreeNode
	var end bool
	if root.Val != 0 {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		var values []*TreeNode
		for i := 0; i < len(queue); i++ {
			if queue[i] != nil {
				values = append(values, queue[i].Left)
				values = append(values, queue[i].Right)
			} else {
				end = true
			}
			if end && queue[i] != nil {
				return false
			}

		}
		queue = values
	}
	return true
}
