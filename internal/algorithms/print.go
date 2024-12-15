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
	InOrder(root.Left)
	InOrder(root.Right)
	fmt.Printf("%d\n", root.Value)
}

func PostOrder(root *entity.TreeNode) {
 	if nil == root {
  		return
 	}
	fmt.Printf("%d\n", root.Value)
	InOrder(root.Left)
	InOrder(root.Right)
}