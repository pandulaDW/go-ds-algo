package hashing

// SearchHashTable will return the element searched using the given item ID
func (hTable *hashTable) SearchHashTable(id int) interface{} {
	hashIndex := hTable.hashFunc(id)
	return hTable.hashArr[hashIndex].Find(func(val interface{}) bool {
		return val.(*Data).ID == id
	})
}
