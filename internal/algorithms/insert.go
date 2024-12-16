package algorithms

import "github.com/aeonva1ues/avltree/internal/entity"

func Insert(root *entity.TreeNode, value int) *entity.TreeNode {
	if root == nil {
		return &entity.TreeNode{Value: value, Height: 1}
	}
	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else if value >= root.Value {
		root.Right = Insert(root.Right, value)
	}
	root.Height = 1 + max(root.Left.GetHeight(), root.Right.GetHeight())
	return root
}

func AvlInsert(root *entity.TreeNode, value int) *entity.TreeNode {
	result := Insert(root, value)
	return Rotate(result)
}