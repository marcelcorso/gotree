package gotree

type Tree struct {
	root *Node
}

type Node struct {
	Key   string
	Left  *Node
	Right *Node
}

func (v *Tree) Add(key string) {
	if v.root == nil {
		v.root = &Node{Key: key}
	} else {
		v._Add(v.root, key)
	}
}

func (v *Tree) _Add(n *Node, key string) {
	if n.Key > key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			v._Add(n.Right, key)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			v._Add(n.Left, key)
		}
	}
}

func (v *Tree) InOrder(visit func(key string)) {
	if v.root != nil {
		v._InOrder(v.root, visit)
	}
}

func (v *Tree) _InOrder(n *Node, visit func(key string)) {
	if n.Left != nil {
		v._InOrder(n.Left, visit)
	}
	visit(n.Key)
	if n.Right != nil {
		v._InOrder(n.Right, visit)
	}
}

func (v *Tree) Search(key string) (response string) {
	if v.root != nil {
		response = v.search(v.root, key)
	}
	return
}

func (v *Tree) search(n *Node, key string) (response string) {

	if n.Key == key {
		return "found"
	}

	side := &n.Left
	if n.Key > key {
		side = &n.Right
	}

	if *side != nil {
		return v.search(*side, key)
	}
	return
}
