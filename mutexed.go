package gotree

import "sync"

// MTree is a write sync.RWMutex wrapped Tree
type MTree struct {
	t *Tree
	sync.RWMutex
}

func NewMTree() *MTree {
	mt := MTree{}
	mt.t = &Tree{}
	return &mt
}

func (t *MTree) Add(key string) {
	t.Lock()
	defer t.Unlock()
	t.t.Add(key)
}

func (t *MTree) Search(key string) string {
	t.RLock()
	defer t.RUnlock()
	return t.t.Search(key)
}
