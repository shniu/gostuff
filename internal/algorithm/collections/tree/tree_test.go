package tree

import "testing"

func TestEmptyTree(t *testing.T) {
	var t1 Treeable
	t1 = NewTree()
	t1.PreOrder()
}
