package main

import (
  "fmt"
)

type Tree struct {
  root *Node
}

func (v *Tree) Add(key string) {
  if v.root == nil {
    v.root = &Node{key: key}
  } else {
    v._Add(v.root, key)
  }
}

func (v *Tree) _Add(n *Node, key string) {
  if n.key > key {
    if n.right == nil {
      n.right = &Node{key: key}
    } else {
      v._Add(n.right, key)
    }
  } else {
    if n.left == nil {
      n.left = &Node{key: key}
    } else {
      v._Add(n.left, key)
    }
  }
}

func (v *Tree) InOrder(visit func(key string)) {
  if v.root != nil {
    v._InOrder(v.root, visit)
  }
}

func (v *Tree) _InOrder(n *Node, visit func(key string)) {
  if n.left != nil {
    v._InOrder(n.left, visit)
  }
  visit(n.key)
  if n.right != nil {
    v._InOrder(n.right, visit)
  }
}

type Node struct {
  key string
  left *Node
  right *Node
}

func main() {
  t := Tree{}
  keys := []string{"ma", "ro", "zin", "ko", "aq", "er", "se", "ca", "pi", "ty", "ge", "me", "mo"}
  for i := 0; i < len(keys); i++ {
    t.Add(keys[i])
    fmt.Print(keys[i])
    fmt.Print(", ")
  }
  fmt.Println("")

  fmt.Println("in order: ")
  t.InOrder(func (key string) {
    fmt.Print(key)
    fmt.Print(", ")
  })
  fmt.Println("")
}


