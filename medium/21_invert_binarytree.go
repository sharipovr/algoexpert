package main

import "fmt"

// Интуиция:
// Идея простая:
// У каждого узла поменять местами left и right
// Проделать это рекурсивно или итеративно (через стек/очередь) для всех узлов дерева
// 📦 Рекурсивный план:
// Если node == nil, возвращаем nil
// Рекурсивно инвертируем левое и правое поддеревья
// Меняем местами node.left и node.right
// Возвращаем текущий узел

// Узел бинарного дерева
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// Функция инвертирует дерево и возвращает корень
func (tree *BinaryTree) invertBinaryTree() {
	if tree == nil {
		return // Если дерево пустое - ничего делать не нужно
	}

	// Рекурсивно инвертируем левое и правое поддеревья
	tree.Left.invertBinaryTree()
	tree.Right.invertBinaryTree()

	// Меняем местами
	tree.Left, tree.Right = tree.Right, tree.Left
}

func main() {
	// Строим дерево из примера
	tree := &BinaryTree{Value: 1}
	tree.Left = &BinaryTree{Value: 2}
	tree.Right = &BinaryTree{Value: 3}
	tree.Left.Left = &BinaryTree{Value: 4}
	tree.Left.Right = &BinaryTree{Value: 5}
	tree.Right.Left = &BinaryTree{Value: 6}
	tree.Right.Right = &BinaryTree{Value: 7}
	tree.Left.Left.Left = &BinaryTree{Value: 8}
	tree.Left.Left.Right = &BinaryTree{Value: 9}

	tree.invertBinaryTree()

	// Печатаем обход дерева (in-order) для проверки
	fmt.Println("InOrder после инверсии:", inOrder(tree, []int{}))
}

// In-order обход дерева (слева → узел → справа)
func inOrder(tree *BinaryTree, result []int) []int {
	if tree != nil {
		result = inOrder(tree.Left, result)
		result = append(result, tree.Value)
		result = inOrder(tree.Right, result)
	}
	return result
}

/*
        1
       / \
      2   3
     / \ / \
    4  5 6  7
   / \
  8   9

*/
