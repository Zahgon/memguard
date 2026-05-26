/*
	Copyright 2019 Awn Umar <awn@spacetime.dev>

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package casting

import (
	"github.com/awnumar/memguard"
)

// Secure is some generic example struct containing sensitive information.
type Secure struct {
	Key       [32]byte
	Salt      [2]uint64
	Counter   uint64
	Something bool
}

// ByteArray10 allocates and returns a region of memory represented as a fixed-size 10 byte array.
func ByteArray10() (*memguard.LockedBuffer, *[10]byte) {
	_ = "STUB: not implemented"
	// Allocate 10 bytes of memory
	return nil, nil
}

// Return the LockedBuffer along with the cast pointer

// Uint64Array4 allocates a 32 byte memory region and returns it represented as a sequence of four unsigned 64 bit integer values.
func Uint64Array4() (*memguard.LockedBuffer, *[4]uint64) {
	_ = "STUB: not implemented"
	// Allocate the correct amount of memory
	return nil, nil
}

// Return the LockedBuffer along with the cast pointer

// SecureStruct allocates a region of memory the size of a struct type and returns a pointer to that memory represented as that struct type.
func SecureStruct() (*memguard.LockedBuffer, *Secure) {
	_ = "STUB: not implemented"
	// Initialise an instance of the struct type
	return nil, nil
}

// Allocate a LockedBuffer of the correct size

// Return the LockedBuffer along with the initialised struct

// SecureStructArray allocates enough memory to hold an array of Secure structs and returns them.
func SecureStructArray() (*memguard.LockedBuffer, *[2]Secure) {
	_ = "STUB: not implemented"
	// Initialise an instance of the struct type
	return nil, nil
}

// Allocate a LockedBuffer of four times the size of the struct type

// Cast a pointer to the start of the memory into a pointer of a fixed size array of Secure structs of length four

// Return the LockedBuffer along with the array

// SecureStructSlice takes a length and returns a slice of Secure struct values of that length.
func SecureStructSlice(size int) (*memguard.LockedBuffer, []Secure) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialise an instance of the struct type

// Allocate the enough memory to store the struct values

// Construct the slice from its parameters

// Return the LockedBuffer along with the constructed slice
