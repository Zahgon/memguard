package core

import (
	"errors"
	"sync"
	"time"
)

// Interval of time between each verify & re-key cycle.
const interval = 500 * time.Millisecond

// ErrCofferExpired is returned when a function attempts to perform an operation using a secure key container that has been wiped and destroyed.
var ErrCofferExpired = errors.New("<memguard::core::ErrCofferExpired> attempted usage of destroyed key object")

/*
Coffer is a specialized container for securing highly-sensitive, 32 byte values.
*/
type Coffer struct {
	sync.Mutex

	left  *Buffer
	right *Buffer

	rand *Buffer
}

// NewCoffer is a raw constructor for the *Coffer object.
func NewCoffer() *Coffer { _ = "STUB: not implemented"; return nil }

// Init is used to reset the value stored inside a Coffer to a new random 32 byte value, overwriting the old.
func (s *Coffer) Init() error { _ = "STUB: not implemented"; return nil }

// left = left XOR hash(right)

/*
View returns a snapshot of the contents of a Coffer inside a Buffer. As usual the Buffer should be destroyed as soon as possible after use by calling the Destroy method.
*/
func (s *Coffer) View() (*Buffer, error) { _ = "STUB: not implemented"; return nil, nil }

// data = hash(right) XOR left

/*
Rekey is used to re-key a Coffer. Ideally this should be done at short, regular intervals.
*/
func (s *Coffer) Rekey() error { _ = "STUB: not implemented"; return nil }

// Hash the current right partition for later.

// new_right = current_right XOR buf32

// new_left = current_left XOR hash(current_right) XOR hash(new_right)

/*
Destroy wipes and cleans up all memory related to a Coffer object. Once this method has been called, the Coffer can no longer be used and a new one should be created instead.
*/
func (s *Coffer) Destroy() error { _ = "STUB: not implemented"; return nil }

// Destroyed returns a boolean value indicating if a Coffer has been destroyed.
func (s *Coffer) Destroyed() bool { _ = "STUB: not implemented"; return false }

func (s *Coffer) destroyed() bool { _ = "STUB: not implemented"; return false }
