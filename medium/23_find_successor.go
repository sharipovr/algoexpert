package main

import "fmt"

//	Интуиция:
//	In-order traversal =
//	лево → узел → право
//	In-order successor — это следующий узел в этом обходе.
//	Например, in-order для дерева:
/*
        1
       / \
      2   3
     / \
    4   5
   /
  6
*/
// ➡ in-order: 6 → 4 → 2 → 5 → 1 → 3
// Successor of 5 → 1 ✅
// Возможны 2 случая:
// 1. Есть правое поддерево
// Тогда successor — это самый левый узел в правом поддереве.
// 2. Нет правого поддерева
// Тогда идём вверх по дереву по parent, пока:
// мы не перестаём быть «правым ребёнком»
// тогда наш родитель — это и есть successor

type BinaryTree struct {
	Value  int
	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

// Находит in-order successor узла
func findSuccessor(tree, node *BinaryTree) *BinaryTree {
	// Если есть правое поддерево — идём в самый левый узел правее
	if node.Right != nil {
		return getLeftmost(node.Right)
	}

	// Иначе поднимаемся вверх по дереву
	return getFirstParentOnLeft(node)
}

// Находит самого левого потомка (вспомогательная)
func getLeftmost(node *BinaryTree) *BinaryTree {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Поднимаемся вверх, пока текущий узел — это правый ребёнок
func getFirstParentOnLeft(node *BinaryTree) *BinaryTree {
	current := node
	for current.Parent != nil && current == current.Parent.Right {
		current = current.Parent
	}
	return current.Parent // может быть nil, если мы были в самом конце
}

func main() {
	// Построим вручную дерево из примера
	// tree = 1
	//       / \
	//      2   3
	//     / \
	//    4   5
	//   /
	//  6

	// нижние
	n6 := &BinaryTree{Value: 6}
	n4 := &BinaryTree{Value: 4, Left: n6}
	n6.Parent = n4

	n5 := &BinaryTree{Value: 5}
	n2 := &BinaryTree{Value: 2, Left: n4, Right: n5}
	n4.Parent = n2
	n5.Parent = n2

	n3 := &BinaryTree{Value: 3}
	n1 := &BinaryTree{Value: 1, Left: n2, Right: n3}
	n2.Parent = n1
	n3.Parent = n1

	// Найти successor для узла 5
	successor := findSuccessor(n1, n5)
	if successor != nil {
		fmt.Println("Successor:", successor.Value) // Должно быть 1
	} else {
		fmt.Println("No successor")
	}
}
