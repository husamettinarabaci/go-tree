package tree

import (
	"encoding/json"
)

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

func (n *Node[T]) Find(v T) *Node[T] {
	if n.Value == v {
		return n
	}
	for _, child := range n.Children {
		tmp := child.Find(v)
		if tmp != nil {
			return tmp
		}
	}
	return nil
}

func (t *Tree[T]) GetChildrenList(parent T) []T {

	if t.Root == nil {
		return []T{}
	}
	parentNode := t.Root.Find(parent)
	if parentNode == nil {
		return []T{}
	}
	return parentNode.GetChildrenToList()
}

func (n *Node[T]) GetChildrenToList() []T {
	var list []T
	list = append(list, n.Value)
	for _, child := range n.Children {
		list = append(list, child.GetChildrenToList()...)
	}
	return list
}

func (o Tree[T]) ToJson() string {
	j, _ := json.Marshal(o)
	return string(j)
}

func (o Tree[T]) ToJsonPretty() string {
	b, _ := json.MarshalIndent(o, "", "  ")
	return string(b)
}

func FromTreeJson[T comparable](s string) Tree[T] {
	var p Tree[T]
	json.Unmarshal([]byte(s), &p)
	return p
}

func (o Node[T]) ToJson() string {
	j, _ := json.Marshal(o)
	return string(j)
}

func (o Node[T]) ToJsonPretty() string {
	b, _ := json.MarshalIndent(o, "", "  ")
	return string(b)
}

func FromNodeJson[T comparable](s string) Node[T] {
	var p Node[T]
	json.Unmarshal([]byte(s), &p)
	return p
}
