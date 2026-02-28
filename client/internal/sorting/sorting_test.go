package sorting

import (
	"testing"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
)

// intItems builds a []*pb.DataItem slice from int32 values for concise test data.
func intItems(vals ...int32) []*pb.DataItem {
	items := make([]*pb.DataItem, len(vals))
	for i, v := range vals {
		items[i] = &pb.DataItem{IntVal: v}
	}
	return items
}

// strItems builds a []*pb.DataItem slice from string values.
func strItems(vals ...string) []*pb.DataItem {
	items := make([]*pb.DataItem, len(vals))
	for i, v := range vals {
		items[i] = &pb.DataItem{Stringval: v}
	}
	return items
}

// intVals extracts int values from DataItems for easy comparison.
func intVals(items []*pb.DataItem) []int32 {
	vals := make([]int32, len(items))
	for i, item := range items {
		vals[i] = item.GetIntVal()
	}
	return vals
}

func equal(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type sortFunc func([]*pb.DataItem, CompareFunc) []*pb.DataItem

func TestSortAlgorithms(t *testing.T) {
	algorithms := []struct {
		name string
		fn   sortFunc
	}{
		{"InsertionSort", InsertionSort},
		{"MergeSort", MergeSort},
		{"QuickSort", QuickSort},
	}

	tests := []struct {
		name    string
		input   []int32
		wantAsc []int32
	}{
		{"already sorted", []int32{1, 2, 3, 4, 5}, []int32{1, 2, 3, 4, 5}},
		{"reverse sorted", []int32{5, 4, 3, 2, 1}, []int32{1, 2, 3, 4, 5}},
		{"unsorted", []int32{3, 1, 4, 1, 5, 9, 2, 6}, []int32{1, 1, 2, 3, 4, 5, 6, 9}},
		{"single element", []int32{42}, []int32{42}},
		{"empty", []int32{}, []int32{}},
		{"all equal", []int32{7, 7, 7, 7}, []int32{7, 7, 7, 7}},
		{"duplicates", []int32{3, 1, 2, 1, 3}, []int32{1, 1, 2, 3, 3}},
		{"two elements", []int32{2, 1}, []int32{1, 2}},
		{"negative values", []int32{-3, 0, -1, 2}, []int32{-3, -1, 0, 2}},
	}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			for _, tc := range tests {
				t.Run(tc.name+" asc", func(t *testing.T) {
					items := intItems(tc.input...)
					result := algo.fn(items, ByIntValAsc)
					got := intVals(result)
					if !equal(got, tc.wantAsc) {
						t.Errorf("got %v, want %v", got, tc.wantAsc)
					}
				})

				t.Run(tc.name+" desc", func(t *testing.T) {
					items := intItems(tc.input...)
					result := algo.fn(items, ByIntValDesc)
					got := intVals(result)
					// Reverse of wantAsc
					wantDesc := make([]int32, len(tc.wantAsc))
					for i, v := range tc.wantAsc {
						wantDesc[len(tc.wantAsc)-1-i] = v
					}
					if !equal(got, wantDesc) {
						t.Errorf("got %v, want %v", got, wantDesc)
					}
				})
			}
		})
	}
}

func TestSortAlgorithms_Strings(t *testing.T) {
	algorithms := []struct {
		name string
		fn   sortFunc
	}{
		{"InsertionSort", InsertionSort},
		{"MergeSort", MergeSort},
		{"QuickSort", QuickSort},
	}

	input := []string{"banana", "apple", "cherry", "date", "apple"}
	wantAsc := []string{"apple", "apple", "banana", "cherry", "date"}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			items := strItems(input...)
			result := algo.fn(items, ByStringValAsc)
			for i, item := range result {
				if item.GetStringval() != wantAsc[i] {
					t.Errorf("index %d: got %q, want %q", i, item.GetStringval(), wantAsc[i])
				}
			}
		})
	}
}

func TestComparators(t *testing.T) {
	a := &pb.DataItem{IntVal: 1, Stringval: "alpha"}
	b := &pb.DataItem{IntVal: 2, Stringval: "beta"}
	eq := &pb.DataItem{IntVal: 1, Stringval: "alpha"}

	tests := []struct {
		name string
		fn   CompareFunc
		a, b *pb.DataItem
		want bool
	}{
		{"IntAsc a<b", ByIntValAsc, a, b, true},
		{"IntAsc a>b", ByIntValAsc, b, a, false},
		{"IntAsc a=a", ByIntValAsc, a, eq, false},

		{"IntDesc a<b", ByIntValDesc, a, b, false},
		{"IntDesc a>b", ByIntValDesc, b, a, true},
		{"IntDesc a=a", ByIntValDesc, a, eq, false},

		{"StringAsc a<b", ByStringValAsc, a, b, true},
		{"StringAsc a>b", ByStringValAsc, b, a, false},
		{"StringAsc a=a", ByStringValAsc, a, eq, false},

		{"StringDesc a<b", ByStringValDesc, a, b, false},
		{"StringDesc a>b", ByStringValDesc, b, a, true},
		{"StringDesc a=a", ByStringValDesc, a, eq, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.fn(tc.a, tc.b); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMergeSort_Stability(t *testing.T) {
	// Items with same IntVal but different StringVal to track identity.
	items := []*pb.DataItem{
		{IntVal: 2, Stringval: "second-a"},
		{IntVal: 1, Stringval: "first-a"},
		{IntVal: 2, Stringval: "second-b"},
		{IntVal: 1, Stringval: "first-b"},
		{IntVal: 3, Stringval: "third-a"},
	}

	result := MergeSort(items, ByIntValAsc)

	// Equal-key items must preserve original relative order.
	wantStrings := []string{"first-a", "first-b", "second-a", "second-b", "third-a"}
	for i, item := range result {
		if item.GetStringval() != wantStrings[i] {
			t.Errorf("index %d: got %q, want %q (stability violated)", i, item.GetStringval(), wantStrings[i])
		}
	}
}
