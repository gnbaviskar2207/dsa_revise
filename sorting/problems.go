package sorting

import basicrecursion "github.com/gnbaviskar2207/dsa_revise/basic_recursion"

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		lowestEleIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[lowestEleIdx] {
				lowestEleIdx = j
			}
		}
		temp := arr[i]
		arr[i] = arr[lowestEleIdx]
		arr[lowestEleIdx] = temp
	}
	basicrecursion.PrintArray(arr)
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	basicrecursion.PrintArray(arr)
}

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			} else {
				break
			}
		}
	}
	basicrecursion.PrintArray(arr)
}

func MergeSort(arr []int) {
	basicrecursion.PrintArray(mergeSortAlgo(arr, 0, len(arr)-1))
}

func mergeSortAlgo(arr []int, low int, high int) []int {
	if low >= high {
		return []int{arr[low]}
	}

	mid := (low + high) / 2
	first := mergeSortAlgo(arr, low, mid)
	second := mergeSortAlgo(arr, mid+1, high)
	return mergeTwoSortedArrays(first, second)

}

func mergeTwoSortedArrays(first []int, second []int) []int {
	firstIndex := 0
	secondIndex := 0
	outputArray := []int{}
	for firstIndex < len(first) && secondIndex < len(second) {
		if first[firstIndex] < second[secondIndex] {
			outputArray = append(outputArray, first[firstIndex])
			firstIndex += 1
		} else {
			outputArray = append(outputArray, second[secondIndex])
			secondIndex += 1
		}
	}

	for firstIndex < len(first) {
		outputArray = append(outputArray, first[firstIndex])
		firstIndex += 1
	}

	for secondIndex < len(second) {
		outputArray = append(outputArray, second[secondIndex])
		secondIndex += 1
	}
	return outputArray
}

func BubbleSortUsingRecursion(arr []int, endIndex int) {

	if endIndex == 1 {
		return
	}

	for i := 0; i < endIndex-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
		}
	}
	BubbleSortUsingRecursion(arr, endIndex-1)
}

func SelectionSortUsingRecusrion(arr []int, start int) {
	// base condition
	if start >= len(arr)-1 {
		return
	}

	minIndex := start
	for i := start + 1; i < len(arr); i++ {
		if arr[i] < arr[start] {
			minIndex = i
		}
	}
	// swap
	temp := arr[start]
	arr[start] = arr[minIndex]
	arr[minIndex] = temp

	SelectionSortUsingRecusrion(arr, start+1)
}

func InsertionSortRecursive(arr []int, last int) {
	if last >= len(arr)-1 {
		return
	}

	for j := last + 1; j > 0; j-- {
		if arr[j] < arr[j-1] {
			// swap
			temp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp
		} else {
			break
		}
	}
	InsertionSortRecursive(arr, last+1)
}

func QuickSort(arr []int, low int, high int) {
	if low > high {
		return
	}

	partitionIndex := performQuickSort(arr, low, high)
	QuickSort(arr, low, partitionIndex-1)
	QuickSort(arr, partitionIndex+1, high)
}

func performQuickSort(arr []int, low int, high int) int {
	pivot := arr[low]
	i := low + 1
	j := high
	for i < j {
		for arr[i] < pivot && i < high {
			i += 1
		}

		for arr[j] > pivot && j > low {
			j -= 1
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}
