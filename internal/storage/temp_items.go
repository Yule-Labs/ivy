package storage

import "sync"

type TempItems struct {
	// Sync locker
	mu sync.RWMutex

	// Expiration queue
	keys map[int][]string
}

func (t *TempItems) Append(value *Item) {
	t.mu.Lock()
	defer t.mu.Unlock()

	ttl := value.ttl.Nanosecond()
	if v, ok := t.keys[ttl]; ok {
		t.keys[ttl] = append(v, value.key)
	} else {
		t.keys[ttl] = []string{
			value.key,
		}
	}
}

func newTempItems() *TempItems {
	return &TempItems{
		keys: make(map[int][]string),
	}
}
