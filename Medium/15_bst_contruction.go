package main

import "fmt"

// Интуиция и план
// Что такое Binary Search Tree (BST)?
// Это бинарное дерево, где:
// У каждого узла есть значение Value
// Все значения в левом поддереве < Value
// Все значения в правом поддереве ≥ Value
// Что мы реализуем:
// insert(Value) — вставка нового значения в правильное место
// contains(Value) — проверка, есть ли значение в дереве
// remove(Value) — удаление первого найденного значения

// Узел BST
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// Вставка нового элемента
func (tree *BST) Insert(num int) *BST {
	if num < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: num}
		} else {
			tree.Left.Insert(num)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: num}
		} else {
			tree.Right.Insert(num)
		}
	}
	return tree
}

// Проверка наличия значения
func (tree *BST) Contains(num int) bool {
	if num < tree.Value {
		if tree.Left == nil {
			return false
		}
		return tree.Left.Contains(num)
	} else if num > tree.Value {
		if tree.Right == nil {
			return false
		}
		return tree.Right.Contains(num)
	}
	return true
}

// Удаление узла
func (tree *BST) Remove(num int) *BST {
	// Мы используем вспомогательную рекурсивную функцию,
	// чтобы отслеживать родителя узла.
	return tree.remove(num, nil)
}

// Вспомогательная рекурсивная функция удаления
func (tree *BST) remove(num int, parent *BST) *BST {
	// В начале отработаем траверсинг дерева
	if num < tree.Value {
		// Идем влево если удаляемое значение меньше текущего
		if tree.Left != nil {
			tree.Left.remove(num, tree)
		}
	} else if num > tree.Value {
		// Идем вправо если удаляемое значение больше текущего
		if tree.Right != nil {
			tree.Right.remove(num, tree)
		}
	} else {
		// Нашли узел, который нужно удалить. Нужно отработать все ситуации с ним
		// их четыре:
		// Случай 1: у узла есть **оба** потомка
		if tree.Left != nil && tree.Right != nil {
			// Находим минимальное значение в правом поддереве
			tree.Value = tree.Right.getMinValue()
			// Удаляем узел с этим минимальным значением (он теперь дублируется)
			tree.Right.remove(tree.Value, tree)

		} else if parent == nil {
			// Случай 2: удаляем **корневой узел**

			// Если у корня есть левый ребёнок — замещаем его текущим деревом
			if tree.Left != nil {
				*tree = *tree.Left
			} else if tree.Right != nil {
				*tree = *tree.Right
			} else {
				// Если у корня нет детей — ничего не делаем
				// (по условию, одиночный узел нельзя удалить)
			}

		} else if parent.Left == tree {
			// Случай 3: у узла **один потомок**, и он — левый ребёнок своего родителя
			if tree.Left != nil {
				// Заменяем узел на его левого потомка
				parent.Left = tree.Left
			} else {
				// Иначе — заменяем на правого
				parent.Left = tree.Right
			}

		} else if parent.Right == tree {
			// Случай 4: узел — правый ребёнок родителя
			if tree.Left != nil {
				// Заменяем на левого потомка
				parent.Right = tree.Left
			} else {
				// Или на правого
				parent.Right = tree.Right
			}
		}
	}

	// Возвращаем ссылку на (возможно изменённое) дерево
	return tree
}
func (tree *BST) getMinValue() int {
	// Ищем самый "левый" узел — он и есть минимальный
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
}

func main() {
	bst := &BST{Value: 10}
	bst.Insert(5).Insert(15).Insert(2).Insert(5).Insert(13).Insert(22).Insert(1).Insert(14)
	bst.Insert(12) // Вставка нового значения

	bst.Remove(10) // Удаление корня

	fmt.Println(bst.Contains(15)) // true
}
