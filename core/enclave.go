package core

import (
	"errors"
	"sync"
)

var (
	key    = &Coffer{}
	keyMtx = sync.Mutex{}
)

func getOrCreateKey() *Coffer { _ = "STUB: not implemented"; return nil }

func getKey() *Coffer { _ = "STUB: not implemented"; return nil }

// ErrNullEnclave is returned when attempting to construct an enclave of size less than one.
var ErrNullEnclave = errors.New("<memguard::core::ErrNullEnclave> enclave size must be greater than zero")

/*
Enclave is a sealed and encrypted container for sensitive data.
*/
type Enclave struct {
	ciphertext []byte
}

/*
NewEnclave is a raw constructor for the Enclave object. The given buffer is wiped after the enclave is created.
*/
func NewEnclave(buf []byte) (*Enclave, error) {
	_ = "STUB: not implemented"
	// Return an error if length < 1.
	return nil, nil
}

// Create a new Enclave.

// Get a view of the key.

// Encrypt the plaintext.

// key is not 32 bytes long

// Destroy our copy of the key.

// Wipe the given buffer.

/*
Seal consumes a given Buffer object and returns its data secured and encrypted inside an Enclave. The given Buffer is destroyed after the Enclave is created.
*/
func Seal(b *Buffer) (*Enclave, error) {
	_ = "STUB: not implemented"
	// Check if the Buffer has been destroyed.
	return nil, nil
}

// Make the buffer mutable so that we can wipe it.

// Construct the Enclave from the Buffer's data.

// Attain a read lock.

// Destroy the Buffer object.

// Return the newly created Enclave.

/*
Open decrypts an Enclave and puts the contents into a Buffer object. The given Enclave is left untouched and may be reused.

The Buffer object should be destroyed after the contents are no longer needed.
*/
func Open(e *Enclave) (*Buffer, error) {
	_ = "STUB: not implemented"
	// Allocate a secure Buffer to hold the decrypted data.
	return nil, nil
}

// ciphertext has invalid length

// Grab a view of the key.

// Decrypt the enclave into the buffer we created.

// Destroy our copy of the key.

// Return the contents of the Enclave inside a Buffer.

/*
EnclaveSize returns the number of bytes of plaintext data stored inside an Enclave.
*/
func EnclaveSize(e *Enclave) int { _ = "STUB: not implemented"; return 0 }
