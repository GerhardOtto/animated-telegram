package sorting

import pb "github.com/gerhardotto/animated-telegram/client/backendservice"

// InsertionSort sorts items in-place using insertion sort (O(n²) time, O(1) space).
// It returns the sorted slice for convenience.
func InsertionSort(items []*pb.DataItem, less CompareFunc) []*pb.DataItem {
	for i := 1; i < len(items); i++ {
		key := items[i]
		j := i - 1
		for j >= 0 && less(key, items[j]) {
			items[j+1] = items[j]
			j--
		}
		items[j+1] = key
	}
	return items
}
