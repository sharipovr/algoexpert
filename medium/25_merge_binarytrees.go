// ✅ Объединение двух бинарных деревьев в одно
//
// 🧠 Интуиция:
// Если в одной и той же позиции в двух деревьях есть узлы,
// их значения складываются. Если узел есть только в одном дереве — он переносится как есть.
// Мы можем использовать рекурсию, проходя по обоим деревьям одновременно.
//
// 📌 Правила объединения:
// - Если оба узла существуют → sum = t1.Val + t2.Val
// - Если один из узлов nil → возвращаем другой
//
// 📝 Ручной прогон:
// tree1 =         1         tree2 =         1
//               / \                     / \
//              3   2                   5   9
//             / \                       / \
//            7   4                     2   7
//                                      \   \
//                                       6   6
//
// → merged =       2
//                /   \
//               8     11
//              / \    / \
//             9   4  7   6
//
// ⏱ Сложность:
// Время: O(n), где n — общее число узлов в обоих деревьях (проходим всё)
// Память: O(h), глубина рекурсии (h — высота самого глубокого дерева)

package main

import "fmt"

// Структура бинарного дерева
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// ✅ Главная функция объединения деревьев
func MergeBinaryTrees(tree1, tree2 *BinaryTree) *BinaryTree {
	// Базовый случай: если один из узлов отсутствует — возвращаем другой
	if tree1 == nil {
		return tree2
	}
	if tree2 == nil {
		return tree1
	}

	// Создаём новый узел, складывая значения из двух узлов
	merged := &BinaryTree{Value: tree1.Value + tree2.Value}

	// Рекурсивно объединяем левое поддерево
	merged.Left = MergeBinaryTrees(tree1.Left, tree2.Left)

	// Рекурсивно объединяем правое поддерево
	merged.Right = MergeBinaryTrees(tree1.Right, tree2.Right)

	// Возвращаем корень нового объединённого дерева
	return merged
}

// Вспомогательная функция для печати дерева in-order
func PrintInOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	PrintInOrder(tree.Left)
	fmt.Print(tree.Value, " ")
	PrintInOrder(tree.Right)
}

func main() {
	// Пример из задачи
	t1 := &BinaryTree{Value: 1}
	t1.Left = &BinaryTree{Value: 3}
	t1.Right = &BinaryTree{Value: 2}
	t1.Left.Left = &BinaryTree{Value: 7}
	t1.Left.Right = &BinaryTree{Value: 4}

	t2 := &BinaryTree{Value: 1}
	t2.Left = &BinaryTree{Value: 5}
	t2.Right = &BinaryTree{Value: 9}
	t2.Left.Left = &BinaryTree{Value: 2}
	t2.Right.Left = &BinaryTree{Value: 7}
	t2.Right.Right = &BinaryTree{Value: 6}

	merged := MergeBinaryTrees(t1, t2)

	fmt.Print("Merged in-order: ")
	PrintInOrder(merged)
	fmt.Println()
}
