package main

import "fmt"

// Интуиция:
// В pre-order обходе порядок:
// Текущий узел → Левое поддерево → Правое поддерево
// Первый элемент — это корень дерева
// Далее:
// Всё, что меньше корня → левое поддерево
// Всё, что больше корня → правое поддерево
// Пример:
// preOrder = [10, 4, 2, 1, 5, 17, 19, 18]
// 10 — корень
// Всё что <10 (до 17): 4, 2, 1, 5 → левое поддерево
// Всё что >10: 17, 19, 18 → правое поддерево
// 📦 План:
// Поддерживаем индекс currentIdx, указывающий на текущий элемент в массиве
// Рекурсивно строим дерево:
// текущий элемент → создаём узел
// всё что меньше → идёт влево
// всё что больше → идёт вправо
// Для каждого рекурсивного вызова задаём допустимый диапазон minValue, maxValue (как в задаче Validate BST)
type BST20 struct {
	Value int

	Left  *BST20
	Right *BST20
}

// Главная функция — принимает preorder массив и возвращает корень дерева
func ReconstructBST(preOrder []int) *BST20 {
	idx := 0
	return helper(preOrder, &idx, -1<<63, 1<<63-1)
}

// Строим поддерево в пределах допустимого диапазона (min, max)
func helper(array []int, currentIdx *int, minVal, maxVal int) *BST20 {
	// Если всё просмотрено — возвращаем nil
	if *currentIdx >= len(array) {
		return nil
	}

	currentVal := array[*currentIdx]

	// Значение должно быть в пределах [minVal, maxVal)
	if currentVal < minVal || currentVal >= maxVal {
		return nil // текущий узел не может быть размещён здесь
	}

	// Создаём узел
	*currentIdx++
	node := &BST20{Value: currentVal}

	// Рекурсивно строим левое поддерево
	node.Left = helper(array, currentIdx, minVal, currentVal)

	// Рекурсивно строим правое поддерево
	node.Right = helper(array, currentIdx, currentVal, maxVal)

	return node
}

func main() {
	values := []int{10, 4, 2, 1, 5, 17, 19, 18}
	root := ReconstructBST(values)

	// Проверка in-order обходом — должен вернуть отсортированный массив
	fmt.Println("InOrder:", inOrderTraverse(root, []int{}))
}

// Вспомогательная функция для обхода дерева (для проверки)
func inOrderTraverse(tree *BST20, array []int) []int {
	if tree != nil {
		array = inOrderTraverse(tree.Left, array)
		array = append(array, tree.Value)
		array = inOrderTraverse(tree.Right, array)
	}
	return array
}
