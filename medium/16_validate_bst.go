package main

import "fmt"

//type BST struct {
//	Value int
//	Left  *BST
//	Right *BST
//}

// Главная функция проверки валидности BST
func validateBST(tree *BST) bool {
	// Интуиция и план:
	// Чтобы дерево было валидным BST, для каждого узла должно выполняться:
	// Всё в левом поддереве < значение узла
	// Всё в правом поддереве ≥ значение узла
	// И это должно выполняться на всех уровнях, не только между родителем и детьми.
	// Решение — рекурсивная проверка с диапазоном:
	// Мы будем передавать допустимый диапазон значений для каждого узла:
	// корень: от -∞ до +∞
	// левое поддерево: всё должно быть < текущего значения
	// правое поддерево: всё должно быть ≥ текущего значения
	// Время: O(n) — каждый узел обходим 1 раз
	// Память: O(h) — глубина стека (глубина дерева)

	// Начальный диапазон — от минимального до максимального int
	return validateHelper(tree, -1<<63, 1<<63-1)
}

// Рекурсивная проверка с допустимым диапазоном (min, max)
func validateHelper(node *BST, min, max int) bool {
	if node == nil {
		// Пустой узел всегда валиден
		return true
	}

	if node.Value < min || node.Value >= max {
		// Нарушено условие BST: значение вне допустимого диапазона
		return false
	}

	// Проверяем левое поддерево:
	// всё должно быть строго МЕНЬШЕ текущего значения
	isLeftValid := validateHelper(node.Left, min, node.Value)

	// Проверяем правое поддерево:
	// всё должно быть БОЛЬШЕ ИЛИ РАВНО текущему значению
	isRightValid := validateHelper(node.Right, node.Value, max)

	return isLeftValid && isRightValid
}

// Пример дерева для теста
func main() {
	tree := &BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left:  &BST{Value: 1},
			},
			Right: &BST{Value: 5},
		},
		Right: &BST{
			Value: 15,
			Left: &BST{
				Value: 13,
				Right: &BST{Value: 14},
			},
			Right: &BST{Value: 22},
		},
	}

	fmt.Println(validateBST(tree)) // true
}
