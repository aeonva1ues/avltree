package algorithms

import (
	"github.com/aeonva1ues/avltree/internal/entity"
)

func Delete(root *entity.TreeNode, val int) *entity.TreeNode {
	if nil == root {
		return root
	}
	if root.Value > val {
		root.Left = Delete(root.Left, val)
	}
	if root.Value < val {
		root.Right = Delete(root.Right, val)
	}
	if root.Value == val {
		if root.Left == nil && root.Right == nil {
			root = nil
			return root
		}
		if root.Left == nil && root.Right != nil {
			temp := root.Right
			root = nil
			root = temp
			return root
		}
		if root.Left != nil && root.Right == nil {
			temp := root.Left
			root = nil
			root = temp
			return root
		}

		temp := root
		for temp != nil && temp.Left != nil {
			temp = temp.Left
		}
		root.Value = temp.Value
		root.Right = Delete(root.Right, temp.Value)
	}
	return root
}
