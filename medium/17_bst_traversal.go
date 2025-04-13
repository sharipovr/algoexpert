package main

import "fmt"

type BST17 struct {
	Value int
	Left  *BST17
	Right *BST17
}

// 🔵 In-order (Лево → Узел → Право)
// Возвращает значения в отсортированном порядке (если дерево валидное BST17)
func (tree *BST17) InOrderTraverse(array []int) []int {
	if tree != nil {
		// Сначала обходим левое поддерево
		array = tree.Left.InOrderTraverse(array)

		// Потом добавляем значение текущего узла
		// Зачем: по правилам in-order, мы обходим сначала левую часть, потом текущий узел
		array = append(array, tree.Value)

		// Затем обходим правое поддерево
		array = tree.Right.InOrderTraverse(array)
	}
	return array // Возвращаем обновлённый массив на каждом шаге
}

// 🔴 Pre-order (Узел → Лево → Право)
// Полезен, если нужно сохранить/восстановить структуру дерева
func (tree *BST17) PreOrderTraverse(array []int) []int {
	if tree != nil {
		// Сначала добавляем текущий узел
		// Зачем: pre-order — сначала текущий, потом его поддеревья
		array = append(array, tree.Value)

		// Обходим левое поддерево
		array = tree.Left.PreOrderTraverse(array)

		// Обходим правое поддерево
		array = tree.Right.PreOrderTraverse(array)
	}
	return array
}

// 🟢 Post-order (Лево → Право → Узел)
// Полезен для удаления дерева или работы "снизу вверх"
func (tree *BST17) PostOrderTraverse(array []int) []int {
	if tree != nil {
		// Сначала обходим левое поддерево
		array = tree.Left.PostOrderTraverse(array)

		// Потом обходим правое поддерево
		array = tree.Right.PostOrderTraverse(array)

		// Потом добавляем текущий узел
		// Зачем: post-order — всегда завершаем обход внизу
		array = append(array, tree.Value)
	}
	return array
}

func main() {
	tree := &BST17{
		Value: 10,
		Left: &BST17{
			Value: 5,
			Left: &BST17{
				Value: 2,
				Left:  &BST17{Value: 1},
			},
			Right: &BST17{Value: 5},
		},
		Right: &BST17{
			Value: 15,
			Right: &BST17{Value: 22},
		},
	}

	/*
	        10
	       /  \
	      5    15
	     / \     \
	    2   5     22
	   /
	  1
	*/

	fmt.Println("InOrder:   ", tree.InOrderTraverse([]int{}))   // [1 2 5 5 10 15 22]
	fmt.Println("PreOrder:  ", tree.PreOrderTraverse([]int{}))  // [10 5 2 1 5 15 22]
	fmt.Println("PostOrder: ", tree.PostOrderTraverse([]int{})) // [1 2 5 5 22 15 10]
}
