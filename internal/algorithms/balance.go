package algorithms

import (
	"math"

	"github.com/aeonva1ues/avltree/internal/entity"
)

func Rotate(vertex *entity.TreeNode) *entity.TreeNode {
	if math.Abs(float64(vertex.Left.GetHeight() - vertex.Right.GetHeight())) < 2 {
		// Вращать не надо, поддерево с вершиной vertex и так сбалансировано.
		return vertex
 	}
	if vertex.Left.GetHeight()-vertex.Right.GetHeight() == -2 {
		// Нам нужны левые вращения.
		b := vertex.Right
		R := b.Right
		C := b.Left
		if C.GetHeight() <= R.GetHeight() {
			// Нужно малое левое вращение.
			return smallLeftRotation(vertex)
		} else {
			// Нужно большое левое вращение.
			return bigLeftRotation(vertex)
		}
	}
	if vertex.Left.GetHeight()-vertex.Right.GetHeight() == 2 {
		// Нам нужны правые вращения.
		b := vertex.Left
		C := b.Right
		L := b.Left
		if C.GetHeight() <= L.GetHeight() {
			// Нужно малое правое вращение.
			return smallRightRotation(vertex)
		} else {
				// Нужно большое правое вращение.
				return bigRightRotation(vertex)
		}
	}
	return nil
}

func bigLeftRotation(vertex *entity.TreeNode) *entity.TreeNode {
	// Правым ребёнком становится новый корень правого поддерева.
	vertex.Right = smallRightRotation(vertex.Right)
	// Возвращаем новый корень поддерева.
 	return smallLeftRotation(vertex)
}

func smallLeftRotation(a *entity.TreeNode) *entity.TreeNode {
	// Задаём обозначения.
	b := a.Right
	C := b.Left

	// Переусыновляем вершины.
	a.Right = C
	b.Left = a

	// Корректируем высоты в зависимости от того, равны ли высоты C и R.
	if C.GetHeight() == a.GetHeight() {
		a.Height -= 1
		b.Height += 1
	} else {
		a.Height -= 2
	}
	return b
}

func smallRightRotation(a *entity.TreeNode) *entity.TreeNode {
	// Задаём обозначения.
	b := a.Left
	C := b.Right

	a.Left = C
	b.Right = a

	// Корректируем высоты в зависимости от того, равны ли высоты C и R.
	if C.GetHeight() == a.GetHeight() {
		a.Height += 1
		b.Height -= 1
	} else {
		b.Height -= 2
	}
	return b
}

func bigRightRotation(vertex *entity.TreeNode) *entity.TreeNode {
	// Правым ребёнком становится новый корень правого поддерева.
	vertex.Left = smallLeftRotation(vertex.Left)
	// Возвращаем новый корень поддерева.
	return smallRightRotation(vertex)
}