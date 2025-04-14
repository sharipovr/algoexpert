// ✅ Проверка симметричности бинарного дерева
//
// 🧠 Интуиция:
// Дерево симметрично, если его левое и правое поддеревья — зеркальные отражения друг друга.
// Это значит:
// - Левое поддерево левого узла должно быть зеркально правому поддереву правого узла
// - Значения в симметричных местах должны быть одинаковыми
//
// 📌 Используем рекурсивную функцию, которая сравнивает:
//   isMirror(left, right)
//   - если оба nil → true
//   - если один nil, другой нет → false
//   - если значения разные → false
//   - иначе: сравниваем left.Left vs right.Right и left.Right vs right.Left
//
// 📝 Ручной прогон:
// tree =
//         1
//       /   \
//      2     2
//     / \   / \
//    3  4  4   3
//   / \    \ /  \
//  5   6   6 5
//
// Проверка:
// - 2 и 2 → ok
// - 3 и 3 → ok
// - 5 и 5, 6 и 6 → ok
// - 4 и 4 → ok
// ✅ Всё совпадает зеркально → true
//
// ⏱ Сложность:
// Время: O(n) — каждый узел проверяется один раз
// Память: O(h) — глубина дерева (стек вызовов)

package main

import "fmt"

// BinaryTree — структура узла дерева
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// ✅ Главная функция — проверяет, симметрично ли дерево
func (tree *BinaryTree) IsSymmetric() bool {
	// Проверяем, зеркальны ли левое и правое поддеревья корня
	return isMirror(tree.Left, tree.Right)
}

// 🔁 Рекурсивная функция сравнивает два поддерева зеркально
func isMirror(left, right *BinaryTree) bool {
	// Оба узла пусты → это симметрия
	if left == nil && right == nil {
		return true
	}
	// Один есть, другого нет → несимметрично
	if left == nil || right == nil {
		return false
	}
	// Значения разные → несимметрично
	if left.Value != right.Value {
		return false
	}
	// Сравниваем внешние и внутренние пары рекурсивно:
	// left.Left с right.Right и left.Right с right.Left
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func main() {
	// Пример симметричного дерева из условия
	tree := &BinaryTree{Value: 1}
	tree.Left = &BinaryTree{Value: 2}
	tree.Right = &BinaryTree{Value: 2}
	tree.Left.Left = &BinaryTree{Value: 3}
	tree.Left.Right = &BinaryTree{Value: 4}
	tree.Right.Left = &BinaryTree{Value: 4}
	tree.Right.Right = &BinaryTree{Value: 3}
	tree.Left.Left.Left = &BinaryTree{Value: 5}
	tree.Left.Left.Right = &BinaryTree{Value: 6}
	tree.Right.Right.Left = &BinaryTree{Value: 6}
	tree.Right.Right.Right = &BinaryTree{Value: 5}

	fmt.Println("Is Symmetric:", tree.IsSymmetric()) // true
}
