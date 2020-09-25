package storage

import (
	"time"
)

type DataType byte

const (
	StringType DataType = 's'
	IntType             = 'i'
)

type Item struct {
	key   string
	value interface{}
	vtype DataType
	ttl   time.Time
}

func (i *Item) GetKey() string {
	return i.key
}

func (i *Item) GetValue() interface{} {
	return i.value
}

func NewItem(
	key string,
	value interface{},
	vtype DataType,
	ttl time.Time,
) *Item {
	return &Item{
		key:   key,
		value: value,
		vtype: vtype,
		ttl:   ttl,
	}
}

type Items map[string]*Item

func newItems() Items {
	return make(Items)
}
