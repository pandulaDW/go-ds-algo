package hashing

// Data represents the type of value storing in the array
type Data struct {
	ID      int
	Content string
}

// HashTable will create a hash array based on the given input array
type HashTable struct {
	dataArr []*Data
	HashArr []int
}

// hashFunc is the function used to store and read from the hash table
func (hTable *HashTable) hashFunc(item int) int {
	return item % 10
}

// NewHashTable will make a new hash table based on the input array and store the
// array indices
func NewHashTable(arr []*Data) *HashTable {
	hTable := new(HashTable)
	hTable.dataArr = arr
	hTable.HashArr = make([]int, 10)

	for i, item := range arr {
		index := hTable.hashFunc(item.ID)
		hTable.HashArr[index] = i
	}

	return hTable
}

// SearchHashTable will return the element searched using the given item ID
func (hTable *HashTable) SearchHashTable(id int) *Data {
	hashIndex := hTable.hashFunc(id)
	arrId := hTable.HashArr[hashIndex]
	return hTable.dataArr[arrId]
}
