package tree

type Tree[T comparable] struct {
	Root *Node[T]
}

type Node[T comparable] struct {
	Value    T
	Children []*Node[T]
}

func (t *Tree[T]) Add(v T, parent T) {
	if t.Root == nil {
		t.Root = &Node[T]{Value: v}
		return
	}
	t.Root.Add(v, parent)
}

func (n *Node[T]) Add(v T, parent T) {
	if n.Value == parent {
		n.Children = append(n.Children, &Node[T]{Value: v})
		return
	}
	for _, child := range n.Children {
		child.Add(v, parent)
	}
}
