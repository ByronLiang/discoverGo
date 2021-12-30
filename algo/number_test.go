package algo

import "testing"

func TestSmallerNumbersThanCurrent(t *testing.T) {
	t.Log(SmallerNumbersThanCurrent([]int{4, 4, 2, 4}))
}

func TestUniqueOccurrences(t *testing.T) {
	t.Log(UniqueOccurrences([]int{1, 2}))
}

func TestValidMountainArray(t *testing.T) {
	t.Log(ValidMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
