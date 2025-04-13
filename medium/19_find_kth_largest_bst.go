package main

import "fmt"

//Интуиция и план:
//Что такое k-е наибольшее значение?
//Если отсортировать все значения по убыванию, то это будет элемент на позиции k.
//Например, если дерево содержит:
//[1, 2, 3, 5, 15, 17, 20, 22], то:
//1-е по величине: 22
//2-е по величине: 20
//3-е по величине: 17 ✅
//🪜 План:
//Способ: обход дерева в обратном порядке (Right → Node → Left)
//Это называется reverse in-order traversal
//Он обходит узлы в порядке убывания
//Считаем количество посещённых узлов
//Когда счётчик == k, сохраняем значение и выходим

// BST структура
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// Структура для отслеживания текущего состояния при обходе
type TraversalState struct {
	count  int // сколько узлов мы уже посетили
	result int // значение k-го наибольшего элемента
}

// Находит k-е по величине значение в BST
func findKthLargestValueInBST(tree *BST, k int) int {
	// Создаем структру в которой будем хранить:
	// - сколько узлов мы уже обощли (count)
	// - результат (result), когда дойдем до k-го
	state := &TraversalState{
		count:  0,  // изначально ничего не пройдено
		result: -1, // если не найдём, останется -1
	}
	// Запускаем рекурсивный обход дерева справа налево
	reverseInOrder(tree, k, state)

	// Возвращаем найденный результат
	return state.result
}

// Выполняет обратный in-order обход: сначала правый, потом текущий, потом левый
func reverseInOrder(node *BST, k int, state *TraversalState) {
	// Базовый случай: если узел пустой или мы уже нашли нужное значение — выходим
	if node == nil || state.count >= k {
		return
	}
	// Сначала идём в ПРАВОЕ поддерево — там лежат самые большие значения
	reverseInOrder(node.Right, k, state)
	// Только если мы ещё не нашли k-й элемент
	if state.count < k {
		// Увеличиваем счётчик — мы "посетили" текущий узел
		state.count++

		// Если это уже k-й узел — сохраняем его значение как результат
		if state.count == k {
			state.result = node.Value
			return // можно выйти — всё найдено
		}
	}
	// Затем идём в ЛЕВОЕ поддерево — там лежат меньшие значения
	reverseInOrder(node.Left, k, state)
}

//Сложность:
//Время: O(h + k) где h — высота дерева
//Память: O(h) из-за стека вызова (глубина дерева)

func main() {
	/*
	         15
	        /  \
	       5    20
	      / \   / \
	     2   5 17 22
	    / \
	   1   3
	*/
	tree := &BST{Value: 15}
	tree.Left = &BST{Value: 5}
	tree.Left.Left = &BST{Value: 2}
	tree.Left.Left.Left = &BST{Value: 1}
	tree.Left.Left.Right = &BST{Value: 3}
	tree.Left.Right = &BST{Value: 5}
	tree.Right = &BST{Value: 20}
	tree.Right.Left = &BST{Value: 17}
	tree.Right.Right = &BST{Value: 22}
	k := 3
	fmt.Println(findKthLargestValueInBST(tree, k)) // Output: 17
}
