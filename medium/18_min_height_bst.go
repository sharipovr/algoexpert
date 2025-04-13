package main

import "fmt"

// Интуиция и план:
// Чтобы BST18 имело минимальную высоту, дерево должно быть максимально сбалансированным — то есть:
// Корнем должен быть серединный элемент
// Левое поддерево строим из левой части массива
// Правое поддерево — из правой части массива
// Это похоже на принцип бинарного поиска 🔍 — рекурсивно делим массив на пополам и строим дерево.
// Рекурсивный план:
// Выбираем середину массива → создаём узел
// Рекурсивно создаём левое поддерево из левой части
// Рекурсивно создаём правое поддерево из правой части
// Возвращаем корень дерева

// 🔢 Пример:
// array := [1, 2, 5, 7, 10, 13, 14, 15, 22]
// Корень: 10 (середина)
// Лево: [1, 2, 5, 7] → корень 2
// Право: [13, 14, 15, 22] → корень 14
// И так далее рекурсивно...

type BST18 struct {
	Value int
	Left  *BST18
	Right *BST18
}

// Функция создания сбалансированного BST18 из отсортированного массива
func minHeightBST18(arr []int) *BST18 {
	return constructMinHeightBST18(arr, 0, len(arr)-1)
}

func constructMinHeightBST18(arr []int, start, end int) *BST18 {
	if start > end {
		return nil // базовый случай: пустой диапазон → нет узла
	}

	// Выбираем середину — это станет корнем текущего поддерева
	mid := (start + end) / 2
	node := &BST18{Value: arr[mid]}

	// Рекурсивно строим левое и правое поддеревья
	node.Left = constructMinHeightBST18(arr, start, mid-1)
	node.Right = constructMinHeightBST18(arr, mid+1, end)

	return node
}

// Функция нужна только для целей демонстрации результата, см.задачу 17
func (tree *BST18) InOrderTraverse(array []int) []int {
	if tree != nil {
		array = tree.Left.InOrderTraverse(array)
		array = append(array, tree.Value)
		array = tree.Right.InOrderTraverse(array)
	}
	return array
}

func main() {
	array := []int{1, 2, 5, 7, 10, 13, 14, 15, 22}
	root := minHeightBST18(array)

	// Проверим результат с обходом
	fmt.Println("InOrder: ", root.InOrderTraverse([]int{})) // [1 2 5 7 10 13 14 15 22]
}
