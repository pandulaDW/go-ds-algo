package hashing

import "github.com/pandulaDW/go-ds-algo/linkedList"

// InsertToTable will insert a new element to the hash table
func (hTable *hashTable) InsertToTable(item *Data) {
	index := hTable.hashFunc(item.ID)

	// create new linked list if empty
	if hTable.hashArr[index] == nil {
		hTable.hashArr[index] = linkedList.NewLinkedList()
	}

	currentElemPos := hTable.hashArr[index].FindIndex(func(val interface{}) bool {
		return val.(*Data).ID > item.ID
	})

	if currentElemPos == -1 {
		hTable.hashArr[index].InsertAtEnd(item)
	} else {
		hTable.hashArr[index].InsertAtPosition(uint(currentElemPos), item)
	}
}

// DeleteFromTable will delete the given element from the hash table
func (hTable *hashTable) DeleteFromTable(item *Data) {
	index := hTable.hashFunc(item.ID)

	if hTable.hashArr[index] == nil {
		return
	}

	currentElemPos := hTable.hashArr[index].FindIndex(func(val interface{}) bool {
		return val.(*Data).ID == item.ID
	})

	_ = currentElemPos
}

// SearchHashTable will return the element searched using the given item ID
func (hTable *hashTable) SearchHashTable(id int) interface{} {
	hashIndex := hTable.hashFunc(id)
	return hTable.hashArr[hashIndex].Find(func(val interface{}) bool {
		return val.(*Data).ID == id
	})
}
