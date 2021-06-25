package hashing

import (
	"fmt"
	"github.com/pandulaDW/go-ds-algo/linkedList"
	"strings"
)

// Data represents the type of value storing in the array
type Data struct {
	ID      int
	Content interface{}
}

// hashTable will create a hash array based on the given input array
type hashTable struct {
	dataArr []*Data
	hashArr []*linkedList.LinkedList
}

// hashFunc is the function used to store and read from the hash table
func (hTable *hashTable) hashFunc(item int) int {
	return item % 10
}

// NewHashTable will make a new hash table based on the input array and store the
// array indices
func NewHashTable(arr []*Data) *hashTable {
	hTable := new(hashTable)
	hTable.dataArr = arr
	hTable.hashArr = make([]*linkedList.LinkedList, 10)

	for _, item := range arr {
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

	return hTable
}

func (hTable *hashTable) String() string {
	sb := strings.Builder{}
	for i, list := range hTable.hashArr {
		var row string
		if list == nil {
			row = fmt.Sprintf("%d: []\n", i)
		} else {
			row = fmt.Sprintf("%d: %s\n", i, list.String())
		}
		sb.WriteString(row)
	}
	return sb.String()
}
