package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	current := postorder[len(postorder)-1]
	postorder = postorder[:len(postorder)-1]
	node := TreeNode{
		Val: current,
	}

	var leftInorder []int
	var rightInorder []int
	var found bool
	for i := 0; i < len(inorder); i++ {
		if found {
			rightInorder = append(rightInorder, inorder[i])
		} else {
			if inorder[i] == current {
				found = true
			} else {
				leftInorder = append(leftInorder, inorder[i])
			}
		}
	}
	node.Left = buildTree(leftInorder, postorder[:len(leftInorder)])
	node.Right = buildTree(rightInorder, postorder[len(leftInorder):])
	return &node

}
