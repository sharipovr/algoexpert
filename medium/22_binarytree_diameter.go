package main

import "fmt"

// Интуиция:
// Диаметр дерева — это длина самого длинного пути между двумя любыми узлами, измеренная в числе рёбер.
// Важно: путь не обязан проходить через корень.

// План (рекурсивный обход):
// Обходим дерево пост-ордером (снизу вверх)
// Для каждого узла:
// вычисляем высоту левого и правого поддерева
// потенциальный диаметр через этот узел = LeftHeight + RightHeight
// Храним максимальный найденный диаметр в переменной
// Возвращаем высоту узла = 1 + max(LeftHeight, RightHeight)

// Структура BinaryTree
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// Метод, запускающий расчёт диаметра
func BinaryTreeDiameter(tree *BinaryTree) int {
	maxDiameter := 0
	tree.dfs(&maxDiameter)
	return maxDiameter
}

// Вспомогательная функция: возвращает высоту и обновляет max диаметр
func (tree *BinaryTree) dfs(maxDiameter *int) int {
	if tree == nil {
		return 0 // базовый случай: высота пустого поддерева — 0
	}

	// Получаем высоты левого и правого поддеревьев
	leftHeight := tree.Left.dfs(maxDiameter)
	rightHeight := tree.Right.dfs(maxDiameter)

	// Проверяем длину пути через текущий узел
	currentDiameter := leftHeight + rightHeight
	if currentDiameter > *maxDiameter {
		*maxDiameter = currentDiameter
	}

	// Возвращаем высоту текущего узла
	return 1 + max(leftHeight, rightHeight)
}

// Вспомогательная функция для max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
           1
         /   \
        3     2
       /
      7
     / \
    8   4
   /     \
  9       5
           \
            6
*/

func main() {
	// Дерево из примера
	tree := &BinaryTree{Value: 1}
	tree.Left = &BinaryTree{Value: 3}
	tree.Right = &BinaryTree{Value: 2}
	tree.Left.Left = &BinaryTree{Value: 7}
	tree.Left.Left.Left = &BinaryTree{Value: 8}
	tree.Left.Left.Left.Left = &BinaryTree{Value: 9}
	tree.Left.Left.Right = &BinaryTree{Value: 4}
	tree.Left.Left.Right.Right = &BinaryTree{Value: 5}
	tree.Left.Left.Right.Right.Right = &BinaryTree{Value: 6}

	// Диаметр: 9 → 8 → 7 → 3 → 4 → 5 → 6 = 6 рёбер
	fmt.Println("Diameter:", BinaryTreeDiameter(tree)) // Output: 6
}
