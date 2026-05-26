package memguard

import (
	"github.com/awnumar/memguard/core"
)

/*
Enclave is a sealed and encrypted container for sensitive data.
*/
type Enclave struct {
	*core.Enclave
}

/*
NewEnclave seals up some data into an encrypted enclave object. The buffer is wiped after the data is copied. If the length of the buffer is zero, the function will return nil.

A LockedBuffer may alternatively be converted into an Enclave object using its Seal method. This will also have the effect of destroying the LockedBuffer.
*/
func NewEnclave(src []byte) *Enclave { _ = "STUB: not implemented"; return nil }

/*
NewEnclaveRandom generates and seals arbitrary amounts of cryptographically-secure random bytes into an encrypted enclave object. If size is not strictly positive the function will return nil.
*/
func NewEnclaveRandom(size int) *Enclave {
	_ = "STUB: not implemented"
	// todo: stream data into enclave
	return nil
}

/*
Open decrypts an Enclave object and places its contents into an immutable LockedBuffer. An error will be returned if decryption failed.
*/
func (e *Enclave) Open() (*LockedBuffer, error) { _ = "STUB: not implemented"; return nil, nil }

/*
Size returns the number of bytes of data stored within an Enclave.
*/
func (e *Enclave) Size() int { _ = "STUB: not implemented"; return 0 }
