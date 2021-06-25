package hashing

import "github.com/pandulaDW/go-ds-algo/linkedList"

// SearchHashTable will return the element searched using the given item ID
func (hTable *hashTable) SearchHashTable(id int) *linkedList.LinkedList {
	hashIndex := hTable.hashFunc(id)
	return hTable.hashArr[hashIndex]
}
