package slice_vs_map

import "testing"

var (
	testSet []int
)

func init() {
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, i)
	}
}

func BenchmarkMapWrite(b *testing.B) {
	aMap := make(map[int]int)
	for i := 0; i < b.N; i++ {
		for _, elem := range testSet {
			aMap[elem] = elem
		}
	}
}

func BenchmarkSliceWrite(b *testing.B) {
	var slice []int
	for i := 0; i < b.N; i++ {
		for _, elem := range testSet {
			slice = append(slice, elem)
		}
	}
}

func BenchmarkMapRead(b *testing.B) {
	aMap := make(map[int]int)
	for _, elem := range testSet {
		aMap[elem] = elem
	}
	var value int
	for i := 0; i < b.N; i++ {
		for _, elem := range testSet {
			value = value + aMap[elem]
		}
	}
}

func BenchmarkSliceRead(b *testing.B) {
	var value int
	var slice []int
	for _, elem := range testSet {
		slice = append(slice, elem)
	}
	for i := 0; i < b.N; i++ {
		for _, elem := range testSet {
			for _, item := range slice {
				if item == elem {
					value = value + elem
					break
				}
			}
		}
	}
}
