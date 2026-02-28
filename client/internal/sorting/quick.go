package sorting

import pb "github.com/gerhardotto/animated-telegram/client/backendservice"

// QuickSort sorts items in-place using quicksort (O(n log n) average time, O(log n) stack space).
// It returns the sorted slice for convenience.
func QuickSort(items []*pb.DataItem, less CompareFunc) []*pb.DataItem {
	quickSort(items, less)
	return items
}

func quickSort(items []*pb.DataItem, less CompareFunc) {
	if len(items) <= 1 {
		return
	}
	pivot := partition(items, less)
	quickSort(items[:pivot], less)
	quickSort(items[pivot+1:], less)
}

// partition uses the last element as pivot and returns its final index.
func partition(items []*pb.DataItem, less CompareFunc) int {
	pivot := items[len(items)-1]
	i := 0
	for j := 0; j < len(items)-1; j++ {
		if less(items[j], pivot) {
			items[i], items[j] = items[j], items[i]
			i++
		}
	}
	items[i], items[len(items)-1] = items[len(items)-1], items[i]
	return i
}
