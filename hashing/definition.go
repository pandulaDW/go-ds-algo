package hashing

import (
	"fmt"
	"github.com/pandulaDW/go-ds-algo/linkedList"
	"strings"
)

// DataWithNumID represents the type of value storing in the array.
// A numeric ID field is the only necessity
type DataWithNumID struct {
	ID      int
	Content interface{}
}

// DataWithStringID represents the type of value storing in the array.
//A string ID field is the only necessity
type DataWithStringID struct {
	ID      string
	Content interface{}
}

// hashTable will create a hash array based on the given input array
type hashTable struct {
	hashArr []*linkedList.LinkedList
}

func polynomialRollingHash(str string) int {
	p := 31
	m := int(1e9 + 9)
	powerOfP := 1
	hashVal := 0

	for _, letter := range str {
		hashVal = hashVal + (int(letter-'a'+1)*(powerOfP))%m
		powerOfP = (powerOfP * p) % m
	}

	return hashVal
}

// hashFunc is the function used to store and read from the hash table
func (hTable *hashTable) hashFunc(id interface{}) int {
	if val, ok := id.(int); ok {
		return val % 10
	}
	return polynomialRollingHash(id.(string)) % 10
}

// NewHashTableForStringID will make a new hash table based on the input array with string IDs
// and will store the ids
func NewHashTableForStringID(arr []*DataWithStringID) *hashTable {
	hTable := new(hashTable)
	hTable.hashArr = make([]*linkedList.LinkedList, 10)

	for _, item := range arr {
		hTable.InsertToTable(item)
	}

	return hTable
}

// NewHashTableForNumID will make a new hash table based on the input array with numerical IDs
// and will store the ids
func NewHashTableForNumID(arr []*DataWithNumID) *hashTable {
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
