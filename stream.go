package memguard

import (
	"container/list"
	"os"
	"sync"
)

var (
	// StreamChunkSize is the maximum amount of data that is locked into memory at a time.
	// If you get error allocating memory, increase your system's mlock limits.
	// Use 'ulimit -l' to see mlock limit on unix systems.
	StreamChunkSize = c
	c               = os.Getpagesize() * 4
)

type queue struct {
	*list.List
}

// add data to back of queue
func (q *queue) join(e *Enclave) {
	_ = "STUB: not implemented"

	// add data to front of queue
	return
}

func (q *queue) push(e *Enclave) {
	_ = "STUB: not implemented"

	// pop data off front of queue
	// returns nil if queue is empty
	return
}

func (q *queue) pop() *Enclave {
	_ = "STUB: not implemented"
	// get element at front of queue
	return nil
}

// no data

// success => remove value
// unwrap and return (potential panic)

/*
Stream is an in-memory encrypted container implementing the reader and writer interfaces.

It is most useful when you need to store lots of data in memory and are able to work on it in chunks.
*/
type Stream struct {
	sync.Mutex
	*queue
}

// NewStream initialises a new empty Stream object.
func NewStream() *Stream { _ = "STUB: not implemented"; return nil }

/*
Write encrypts and writes some given data to a Stream object.

The data is broken down into chunks and added to the stream in order. The last thing to be written to the stream is the last thing that will be read back.
*/
func (s *Stream) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

/*
Read decrypts and places some data from a Stream object into a provided buffer.

If there is no data, the call will return an io.EOF error. If the caller provides a buffer
that is too small to hold the next chunk of data, the remaining bytes are re-encrypted and
added to the front of the queue to be returned in the next call.

To be performant, have
*/
func (s *Stream) Read(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Grab the next chunk of data from the stream.
		nil
}

// Copy the contents into the given buffer.

// Check if there is data left over.

// Re-encrypt it and push onto the front of the list.

// Not enough data or perfect amount of data.
// Either way we copied the entire buffer.

// Size returns the number of bytes of data currently stored within a Stream object.
func (s *Stream) Size() int { _ = "STUB: not implemented"; return 0 }

// Next grabs the next chunk of data from the Stream and returns it decrypted inside a LockedBuffer. Any error from the stream is forwarded.
func (s *Stream) Next() (*LockedBuffer, error) { _ = "STUB: not implemented"; return nil, nil }

// does not acquire mutex lock
func (s *Stream) next() (*LockedBuffer, error) {
	_ = "STUB: not implemented"
	// Pop data from the front of the list.
	return nil, nil
}

// Decrypt the data into a guarded allocation.

// Flush reads all of the data from a Stream and returns it inside a LockedBuffer. If an error is encountered before all the data could be read, it is returned along with any data read up until that point.
func (s *Stream) Flush() (*LockedBuffer, error) { _ = "STUB: not implemented"; return nil, nil }
