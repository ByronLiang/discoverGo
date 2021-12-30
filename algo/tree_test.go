package algo

import "testing"

func TestSumNumbers(t *testing.T) {
	root := InitTreeData()
	t.Log(SumNumbers(root))
}
