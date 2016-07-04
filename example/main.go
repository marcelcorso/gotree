package main

import (
	"fmt"

	"github.com/marcelcorso/gotree"
)

func main() {
	t := gotree.Tree{}
	keys := []string{"ma", "ro", "zin", "ko", "aq", "er", "se", "ca", "pi", "ty", "ge", "me", "mo"}
	for i := 0; i < len(keys); i++ {
		t.Add(keys[i])
		fmt.Print(keys[i])
		fmt.Print(", ")
	}
	fmt.Println("")

	fmt.Println("in order: ")
	t.InOrder(func(key string) {
		fmt.Print(key)
		fmt.Print(" ")
	})
	fmt.Println("")

	found := t.Search("ty")
	fmt.Printf("'ty' is in the tree? %v\n", found)
}
