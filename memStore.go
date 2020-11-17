package main

import "errors"

// MemStore is an in memory data storage that implements IO reader and writer
// interfaces
type MemStore struct {
	buffer      []byte
	bufferSize  int
	writeOffset int
	readOffset  int
}

type (
	// Reader is as same as IO Reader
	Reader interface {
		Read(b []byte) (n int, err error)
	}

	// Writer is as same as IO Writer
	Writer interface {
		Write(b []byte) (n int, err error)
	}
)

// CreateStore creates a new MemStore using the given buffer size
func CreateStore(size int) MemStore {
	return MemStore{make([]byte, size), size, 0, 0}
}

// Write, write the given byte slice into the mem store buffer
func (store *MemStore) Write(b []byte) (n int, err error) {
	// return error if the buffer is full
	if store.writeOffset == store.bufferSize {
		return 0, errors.New("End of buffer error")
	}

	// copy the data to the buffer, if buffer exceeds, copy until the end
	var nCopied int
	if len(b)+store.writeOffset > store.bufferSize {
		nCopied = copy(store.buffer[store.writeOffset:store.bufferSize], b)
	} else {
		nCopied = copy(store.buffer[store.writeOffset:len(b)+store.writeOffset], b)
	}

	// move the write offset
	store.writeOffset += nCopied

	return nCopied, nil
}

// Read, copies the data from mem store buffer into the given byte slice
func (store *MemStore) Read(b []byte) (n int, err error) {
	// return error if the readOffset is at the write offset
	if store.readOffset == store.writeOffset {
		return 0, errors.New("End of buffer error")
	}

	// copy the data to the byte slice, if the elements required is higher than
	// the current size of the buffer, copy all the elements
	var nCopied int
	if len(b)+store.readOffset > store.writeOffset {
		nCopied = copy(b, store.buffer[store.readOffset:store.writeOffset])
	} else {
		nCopied = copy(b, store.buffer[store.readOffset:len(b)+store.readOffset])
	}

	// move the reader offset
	store.readOffset += nCopied

	return nCopied, nil
}

// Clear will clear out of the buffer from the store
func (store *MemStore) Clear() {
	// assign the buffer to a new array, so the previous one will be GCd
	// and reset the offsets
	store.buffer = make([]byte, store.bufferSize)
	store.readOffset, store.writeOffset = 0, 0
}

// String will return a string representation of the store's buffer
func (store *MemStore) String() string {
	return string(store.buffer)
}
