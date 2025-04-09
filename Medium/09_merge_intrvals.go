package main

import (
	"fmt"
	"sort"
)

// mergeIntervals takes a slice of intervals, merges overlapping intervals, and returns the merged intervals.
func mergeIntervals(arr [][]int) [][]int {

	// Интуиция:
	// Сначала сортируем интервалы по началу,
	// затем идем по списку и сливаем каждый с предыдущим, если они перекрываются.
	// Перекрытие происходит, если current.start <= previous.end.
	// Edge-кейсы включают: один интервал (ничего не сливается),
	// все интервалы не перекрываются (вернем как есть),
	// все перекрываются (объединяются в один).
	// Сложность: O(n log n) из-за сортировки и O(n) для одного прохода —
	// итого O(n log n) по времени и O(n) по памяти.

	// Если пустой массив - возвращаем пустой результат
	if len(arr) == 0 {
		return [][]int{}
	}

	// Сортируем по началу
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	// Начинаем скан с первого интервала
	merged := [][]int{arr[0]}

	// Гоним по остальным
	for i := 1; i < len(arr); i++ {
		last := merged[len(merged)-1] // последний(предыдущий) в отсортированном
		curr := arr[i]

		// Если пересекаются - объединим (текущий не добавляем!)
		if last[1] >= curr[0] {
			// Увеличиваем конец предыдущего если надо
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			// иначе просто добавим текущий как есть
			merged = append(merged, curr)
		}
	}
	return merged
}

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {4, 7}, {6, 8}, {9, 10}}

	fmt.Println(mergeIntervals(intervals))
}
