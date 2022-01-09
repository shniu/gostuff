package tree

type Treeable interface {
	PreOrder() []int
	InOrder() []int
	PostOrder() []int
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) PreOrder() []int {
	return nil
}

func (n *Node) InOrder() []int {
	return nil
}

func (n *Node) PostOrder() []int {
	return nil
}

func NewTree() (t *Node) {
	t = new(Node)
	return
}
