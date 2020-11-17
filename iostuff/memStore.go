package iostuff

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
	store.buffer = make([]byte, store.bufferSize)
	// reset the offsets
	store.readOffset, store.writeOffset = 0, 0
}

// Flush all the remaning elements in the buffer to the passed slice and
// return the number of elements copied. Will also clear and reset the buffer
//
// Panics if the passed slice is too small in length
func (store *MemStore) Flush(b []byte) (n int) {
	// if the buffer is fully read, just return
	if store.readOffset == store.bufferSize {
		return 0
	}

	// if the slice length is too small, panic
	if len(b) < store.bufferSize-store.readOffset {
		panic(errors.New("Slice length is too small"))
	}

	nCopied := copy(b, store.buffer[store.readOffset:])
	store.Clear()

	return nCopied
}

// String will return a string representation of the store's buffer
func (store *MemStore) String() string {
	return string(store.buffer)
}
