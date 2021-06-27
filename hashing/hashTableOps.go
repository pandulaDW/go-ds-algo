package hashing

import "github.com/pandulaDW/go-ds-algo/linkedList"

// InsertToTable will insert a new element to the hash table
func (hTable *hashTable) InsertToTable(item interface{}) {
	var index int
	if val, ok := item.(*DataWithNumID); ok {
		index = hTable.hashFunc(val.ID)
	} else {
		index = hTable.hashFunc(item.(*DataWithStringID).ID)
	}

	// create new linked list if empty
	if hTable.hashArr[index] == nil {
		hTable.hashArr[index] = linkedList.NewLinkedList()
	}

	currentElemIndex := hTable.hashArr[index].FindIndex(func(val interface{}) bool {
		if _, ok := item.(*DataWithNumID); ok {
			return val.(*DataWithNumID).ID == item.(*DataWithNumID).ID
		}
		return val.(*DataWithStringID).ID == item.(*DataWithStringID).ID
	})

	if currentElemIndex == -1 {
		hTable.hashArr[index].InsertAtEnd(item)
	} else {
		hTable.hashArr[index].InsertAtPosition(uint(currentElemIndex), item)
	}
}

// DeleteFromTable will delete the given element from the hash table
func (hTable *hashTable) DeleteFromTable(id interface{}) {
	index := hTable.hashFunc(id)

	if hTable.hashArr[index] == nil {
		return
	}

	currentElemIndex := hTable.hashArr[index].FindIndex(func(val interface{}) bool {
		if _, ok := id.(int); ok {
			return val.(*DataWithNumID).ID == id
		}
		return val.(*DataWithStringID).ID == id
	})

	if currentElemIndex == -1 {
		return
	}

	hTable.hashArr[index].RemoveUsingIndex(uint(currentElemIndex))
}

// SearchHashTable will return the element searched using the given item ID
func (hTable *hashTable) SearchHashTable(id interface{}) interface{} {
	hashIndex := hTable.hashFunc(id)
	return hTable.hashArr[hashIndex].Find(func(val interface{}) bool {
		if _, ok := id.(int); ok {
			return val.(*DataWithNumID).ID == id
		}
		return val.(*DataWithStringID).ID == id
	})
}
