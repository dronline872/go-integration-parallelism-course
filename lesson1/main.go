package main

import "fmt"

// Функция для слияния двух отсортированных массивов
func mergeArrays(arr1, arr2 []int) []int {
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[k] = arr1[i]
			i++
		} else {
			merged[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		merged[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		merged[k] = arr2[j]
		j++
		k++
	}
	return merged
}

// Функция для сортировки массива пузырьком
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Обмен элементов
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Задание 1: Слияние отсортированных массивов
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8, 9}
	merged := mergeArrays(arr1, arr2)
	fmt.Println("Слияние отсортированных массивов:", merged)

	// Задание 2: Сортировка пузырьком
	arr := []int{6, 5, 3, 1, 8, 7}
	fmt.Println("Исходный массив для сортировки пузырьком:", arr)
	bubbleSort(arr)
	fmt.Println("Отсортированный массив пузырьком:", arr)
}
