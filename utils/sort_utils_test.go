package utils

import (
	"sort"
	"testing"
)

// func TestBubblesort(t *testing.T) {
// 	testTable := map[string]struct {
// 		input    []int
// 		expected []int
// 	}{
// 		"test1": {
// 			[]int{8, 2, 4, 9, 10, 3, 1, 5, 7, 6},
// 			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
// 		},
// 		"test2": {
// 			[]int{-1, 5, -1, -47, 67, 2888, 60},
// 			[]int{-47, -1, -1, 5, 60, 67, 2888},
// 		},
// 	}

// 	for testName, test := range testTable {
// 		result := Bubblesort(test.input)
// 		if !reflect.DeepEqual(result, test.expected) {
// 			t.Errorf("TEST FAILED: %s\nExpected:\n    %v\nGot:\n	%v", testName, test.expected, result)
// 		}
// 	}
// }

func getDescendingElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func BenchmarkBubblesort(b *testing.B) {
	n := 10000
	elements := getDescendingElements(n)
	for i := 0; i < b.N; i++ {
		Bubblesort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	n := 10
	elements := getDescendingElements(n)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}
