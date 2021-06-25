package hashing

import (
	"fmt"
	"github.com/pandulaDW/go-ds-algo/linkedList"
	"strings"
)

// Data represents the type of value storing in the array. ID field is the only necessity
type Data struct {
	ID      int
	Content interface{}
}

// hashTable will create a hash array based on the given input array
type hashTable struct {
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
	hTable.hashArr = make([]*linkedList.LinkedList, 10)

	for _, item := range arr {
		hTable.InsertToTable(item)
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
