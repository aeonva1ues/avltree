package entity

type TreeNode struct {
	Value  int
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

func NewEmptyTreeNode() *TreeNode {
	return &TreeNode{}
}

func (t *TreeNode) GetHeight() int {
	if t == nil {
		return 0
	}
	return t.Height
}