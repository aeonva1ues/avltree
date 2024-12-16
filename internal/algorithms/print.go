package algorithms

import (
	"fmt"

	"github.com/aeonva1ues/avltree/internal/entity"
)

func InOrder(root *entity.TreeNode) {
	if nil == root {
		return
	}
	InOrder(root.Left)
	fmt.Printf("%d\n", root.Value)
	InOrder(root.Right)
}

func PreOrder(root *entity.TreeNode) {
	if nil == root {
		return
	}
	fmt.Printf("%d\n", root.Value)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func PostOrder(root *entity.TreeNode) {
 	if nil == root {
  		return
 	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d\n", root.Value)
}