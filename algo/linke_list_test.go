package algo

import (
	"testing"
	"time"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := InitListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := InitListNode([]int{9, 9, 9, 9})
	res := AddTwoNumbers(l1, l2)
	ShowListNode(res)
}

func TestSwapPairs(t *testing.T) {
	node := InitListNode([]int{2, 3, 4})
	ShowListNode(node)
	newNode := SwapPairs(node)
	ShowListNode(newNode)
}

func TestIsPalindrome(t *testing.T) {
	node := InitListNode([]int{1, 2, 2, 1})
	IsPalindrome(node)
}

func TestRemoveNthFromEnd(t *testing.T) {
	node := InitListNode([]int{1, 2})
	//node := InitListNode([]int{1, 2, 3, 4, 5})
	m := RemoveNthFromEnd(node, 1)
	ShowListNode(m)
}

func TestReverseBetween(t *testing.T) {
	node := InitListNode([]int{1, 2, 3, 4, 5})
	r := ReverseBetween(node, 2, 4)
	ShowListNode(r)
}

func TestRotateRight(t *testing.T) {
	node := InitListNode([]int{1, 2, 3, 4, 5})
	r := RotateRight(node, 2)
	ShowListNode(r)
}

func TestReservoirSampling(t *testing.T) {
	node := InitListNode([]int{1, 2, 3, 4, 5})
	t.Log(ReservoirSampling(node))
	time.Sleep(500 * time.Millisecond)
	t.Log(ReservoirSampling(node))
	time.Sleep(1 * time.Second)
	t.Log(ReservoirSampling(node))
}
