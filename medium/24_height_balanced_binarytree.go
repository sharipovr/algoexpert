// ✅ Проверка, является ли бинарное дерево сбалансированным по высоте
//
// 🧠 Интуиция:
// Высотно сбалансированное дерево — это дерево, где у каждого узла
// разница между высотами левого и правого поддеревьев не больше 1.
// Мы используем post-order DFS (обход снизу вверх):
// 1. Вычисляем высоту левого и правого поддеревьев рекурсивно
// 2. Проверяем, сбалансирован ли текущий узел
// 3. Если хотя бы один несбалансирован — всё дерево несбалансировано
//
// 📝 Ручной прогон:
//         1
//        / \
//       2   3
//      /     \
//     4       6
//    / \
//   7   8
//
// - узлы 7 и 8 → высота 1
// - узел 4 → высота 2, сбалансирован (1 vs 1)
// - узел 2 → высота 3, правое поддерево пусто → diff = 3 - 0 = 3 ❌
// - но если у 2 будет правый потомок — может стать сбалансированным
//
// ⏱ Сложность:
// Время: O(n), каждый узел посещается 1 раз
// Память: O(h), стек вызовов по глубине дерева

package main

import (
	"fmt"
	"math"
)

// BinaryTree — структура узла бинарного дерева
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// ✅ Главная функция проверки баланса дерева
func (tree *BinaryTree) IsBalanced() bool {
	// Запускаем рекурсивную проверку от корня
	// Возвращается высота и булево значение
	_, isBalanced := checkHeight(tree)
	return isBalanced
}

// 🔧 Рекурсивная функция
// Возвращает: (высота поддерева, true/false — сбалансировано ли оно)
func checkHeight(node *BinaryTree) (int, bool) {
	// Базовый случай: пустое дерево
	if node == nil {
		return 0, true // Высота 0, сбалансировано
	}

	// Рекурсивно проверяем левое поддерево
	leftHeight, leftBalanced := checkHeight(node.Left)
	// Рекурсивно проверяем правое поддерево
	rightHeight, rightBalanced := checkHeight(node.Right)

	// Если хотя бы одно из поддеревьев несбалансировано — сразу выходим
	if !leftBalanced || !rightBalanced {
		return 0, false
	}

	// Проверяем баланс текущего узла по разнице высот
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return 0, false // Нарушен баланс
	}

	// Всё ок — возвращаем новую высоту и true
	return 1 + max(leftHeight, rightHeight), true
}

// Вспомогательная функция для вычисления максимума
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Сборка тестового дерева из условия:
	//         1
	//        / \
	//       2   3
	//      /     \
	//     4       6
	//    / \
	//   7   8

	tree := &BinaryTree{Value: 1}
	tree.Left = &BinaryTree{Value: 2}
	tree.Right = &BinaryTree{Value: 3}
	tree.Left.Left = &BinaryTree{Value: 4}
	tree.Left.Left.Left = &BinaryTree{Value: 7}
	tree.Left.Left.Right = &BinaryTree{Value: 8}
	tree.Right.Right = &BinaryTree{Value: 6}

	// Проверка баланса
	fmt.Println("Is Balanced:", tree.IsBalanced()) // true
}
