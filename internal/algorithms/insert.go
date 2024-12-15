package algorithms

import "github.com/aeonva1ues/avltree/internal/entity"

func Insert(root *entity.TreeNode, value, height int) *entity.TreeNode {
	if root == nil {
		return &entity.TreeNode{Value: value, Height: height}
	}
	if value < root.Value {
		root.Left = Insert(root.Left, value, height + 1)
	} else if value >= root.Value {
		root.Right = Insert(root.Right, value, height + 1)
	}
	return root
}

func AvlInsert(root *entity.TreeNode, value, height int) *entity.TreeNode {
	result := Insert(root, value, height)
	return Rotate(result)
}