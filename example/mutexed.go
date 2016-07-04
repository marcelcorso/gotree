package main

import (
	"fmt"
	"sync"

	"github.com/marcelcorso/gotree"
)

func main() {
	var wg sync.WaitGroup
	mt := gotree.NewMTree()
	keys := []string{"ma", "ro", "zin", "ko", "aq", "er", "se", "ca", "pi", "ty", "ge", "me", "mo"}
	for i := 0; i < len(keys); i++ {
		wg.Add(2)
		go func(i int) {
			mt.Add(keys[i])
			fmt.Print(keys[i])
			fmt.Print(", ")
			wg.Done()
		}(i)

		go func(i int) {
			mt.Search(keys[i%2])
			wg.Done()
		}(i)
	}
	wg.Wait()
}
