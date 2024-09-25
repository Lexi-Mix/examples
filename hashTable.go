package hashTable

import (
	"fmt"
)

/*
Example hash table
*/

type DATA map[string]any

var instance DATA

func NewHashTable() *DATA {
	if instance == nil {
		instance = make(DATA)
	}
	return &instance
}

func (d DATA) LOOKUP(key string) *any {
	if key == "" {
		return nil
	}
	if v, ok := d[key]; ok {
		return &v
	}
	return nil
}
func (d DATA) ADD(key string, in any) bool {
	if d.LOOKUP(key) != nil {
		return false
	}
	d[key] = in
	return true
}
func (d DATA) DELETE(key string) bool {
	if d.LOOKUP(key) != nil {
		delete(d, key)
		return true
	}
	return false
}
func (d DATA) CHANGE(key string, in any) bool {
	d[key] = in
	return true
}
func (d DATA) PRINT() {
	for _, v := range d {
		fmt.Printf("%v\n", v)
	}
}
