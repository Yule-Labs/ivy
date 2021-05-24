package storage

import (
	"errors"
	"sync"
	"time"
)

var ErrorItemNotFound = errors.New("item not found")

var once sync.Once
var storage Storage

func GetStorage() Storage {
	return storage
}

type Storage interface {
	Set(key string, value *Item)
	Get(key string) (item *Item, err error)
}

type defaultStorage struct {
	// Sync locker
	mu sync.RWMutex

	// All items
	items Items

	// Expiration keys
	tempItems *TempItems
}

func (s *defaultStorage) Set(key string, value *Item) {
	now := time.Now().UnixNano()
	s.mu.Lock()
	defer s.mu.Unlock()
	ttl := value.ttl.UnixNano()
	if ttl > 0 {
		if ttl < now {
			return
		}
		go s.tempItems.Append(value)
	}
	s.items[key] = value
}

func (s *defaultStorage) Get(key string) (*Item, error) {
	s.mu.RLock()
	item, ok := s.items[key]
	s.mu.RUnlock()
	if ok {
		return item, nil
	}
	return item, ErrorItemNotFound
}

func newDefaultStorage() Storage {
	return &defaultStorage{
		items:     newItems(),
		tempItems: newTempItems(),
	}
}

func InitStorage() {
	once.Do(func() {
		storage = newDefaultStorage()
	})
}