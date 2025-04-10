package main

import "fmt"

func majorityElement(arr []int) int {

	// Идея:
	// Идём по массиву.
	// Поддерживаем двe переменные:
	// candidate — текущий кандидат в majority элемент.
	// count — как бы "вес" кандидата.
	// Если count == 0, выбираем нового кандидата.
	// Если текущий элемент == candidate → увеличиваем count. Иначе — уменьшаем.
	// Почему это работает?
	// Потому, что majority элемент встречается чаще чем все остальные вместе взятые, и он "переживёт" все обнуления счётчика.
	// Complexity: time O(N), space O(1)
	// Edge case: один элемент в массиве, все одинаковые элементы.
	// По умному этот алгоритм называется "Boyer-Moore Majority Voting"

	candidate := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		if count == 0 {
			// Если счетчик обнулился (ну или если мы в начале)
			// берем нового кандидата
			candidate = arr[i]
			count = 1
		} else if arr[i] == candidate {
			// Уже был как кандидат (совпадает), "усиливаем" его
			count++
		} else {
			// иначе, "ослабляем" кандидата
			count--
		}
	}
	return candidate
}

func main() {
	array := []int{1, 2, 3, 2, 2, 1, 2}
	fmt.Println(majorityElement(array)) // 2
}
