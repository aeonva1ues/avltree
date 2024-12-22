package algorithms

import (
	"fmt"

	"github.com/aeonva1ues/avltree/internal/entity"
)

func Delete(root *entity.TreeNode, val int) (*entity.TreeNode, bool) {
	if nil == root {
		return root, false
	}
	if root.Value > val {
		root.Left, _ = Delete(root.Left, val)
	}
	if root.Value < val {
		root.Right, _ = Delete(root.Right, val)
	}
	if root.Value == val {
		if root.Left == nil && root.Right == nil {
			root = nil
			return root, true
		}
		if root.Left == nil && root.Right != nil {
			temp := root.Right
			root = nil
			root = temp
			return root, true
		}
		if root.Left != nil && root.Right == nil {
			temp := root.Left
			root = nil
			root = temp
			return root, true
		}

		temp := root.Right
		for temp != nil && temp.Left != nil {
			temp = temp.Left
		}
		root.Value = temp.Value
		root.Right, _ = Delete(root.Right, temp.Value)
	}
	return root, true
}

func AvlDelete(root *entity.TreeNode, value int) *entity.TreeNode {
	result, deleted := Delete(root, value)
	if deleted {
		return Rotate(result)
	}
	fmt.Println("Значение не найдено!")
	return result
}