package sorting

import pb "github.com/gerhardotto/animated-telegram/client/backendservice"

// MergeSort sorts items in-place using a stable merge sort (O(n log n) time, O(n) space).
// It returns the sorted slice for convenience.
func MergeSort(items []*pb.DataItem, less CompareFunc) []*pb.DataItem {
	if len(items) <= 1 {
		return items
	}
	buf := make([]*pb.DataItem, len(items))
	mergeSort(items, buf, less)
	return items
}

func mergeSort(items, buf []*pb.DataItem, less CompareFunc) {
	if len(items) <= 1 {
		return
	}
	mid := len(items) / 2
	mergeSort(items[:mid], buf[:mid], less)
	mergeSort(items[mid:], buf[mid:], less)
	mergeParts(items, mid, buf, less)
}

// mergeParts merges two sorted halves of items ([:mid] and [mid:]) in-place.
// buf must be at least len(items) in length; only buf[:mid] is used as scratch space.
func mergeParts(items []*pb.DataItem, mid int, buf []*pb.DataItem, less CompareFunc) {
	copy(buf[:mid], items[:mid])
	l, r, i := 0, mid, 0
	for l < mid && r < len(items) {
		if !less(items[r], buf[l]) {
			items[i] = buf[l]
		} else {
			items[i] = items[r]
		}
		i++
	}
	for l < mid {
		items[i] = buf[l]
		l++
		i++
	}
	// Remaining right elements are already in their correct positions in items[r:].
}
