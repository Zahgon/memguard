package memguard

import (
	"bytes"
	"io"

	"github.com/awnumar/memguard/core"
)

/*
LockedBuffer is a structure that holds raw sensitive data.

The number of LockedBuffers that you are able to create is limited by how much memory your system's kernel allows each process to mlock/VirtualLock. Therefore you should call Destroy on LockedBuffers that you no longer need or defer a Destroy call after creating a new LockedBuffer.
*/
type LockedBuffer struct {
	*core.Buffer
}

// Constructs a LockedBuffer object from a core.Buffer while also setting up the finalizer for it.
func newBuffer(buf *core.Buffer) *LockedBuffer { _ = "STUB: not implemented"; return nil }

// Constructs a quasi-destroyed LockedBuffer with size zero.
func newNullBuffer() *LockedBuffer { _ = "STUB: not implemented"; return nil }

/*
NewBuffer creates a mutable data container of the specified size.
*/
func NewBuffer(size int) *LockedBuffer {
	_ = "STUB: not implemented"
	// Construct a Buffer of the specified size.
	return nil
}

// Construct and return the wrapped container object.

/*
NewBufferFromBytes constructs an immutable buffer from a byte slice. The source buffer is wiped after the value has been copied over to the created container.
*/
func NewBufferFromBytes(src []byte) *LockedBuffer {
	_ = "STUB: not implemented"
	// Construct a buffer of the correct size.
	return nil
}

// Move the data over.

// Make the buffer immutable.

// Return the created Buffer object.

/*
NewBufferFromReader reads some number of bytes from an io.Reader into an immutable LockedBuffer.

An error is returned precisely when the number of bytes read is less than the requested amount. Any data read is returned in either case.
*/
func NewBufferFromReader(r io.Reader, size int) (*LockedBuffer, error) {
	_ = "STUB: not implemented"
	// Construct a buffer of the provided size.
	return nil, nil
}

// Attempt to fill it with data from the Reader.

// nothing was read

// partial read

// success

/*
NewBufferFromReaderUntil constructs an immutable buffer containing data sourced from an io.Reader object.

If an error is encountered before the delimiter value, the error will be returned along with the data read up until that point.
*/
func NewBufferFromReaderUntil(r io.Reader, delim byte) (*LockedBuffer, error) {
	_ = "STUB: not implemented"
	// Construct a buffer with a data page that fills an entire memory page.
	return nil, nil
}

// Loop over the buffer a byte at a time.

// If we have filled this buffer...

// Construct a new buffer that is a page size larger.

// Copy the data over.

// Destroy the old one and reassign its variable.

// Attempt to read a single byte.

// if we did not read a byte
// and there was no error
// try again

// if instead there was an error, we're done early
// no data read

// we managed to read a byte, check if it was the delimiter
// note that errors are ignored in this case where we got data

// if first byte was delimiter, there's no data to return

/*
NewBufferFromEntireReader reads from an io.Reader into an immutable buffer. It will continue reading until EOF.

A nil error is returned precisely when we managed to read all the way until EOF. Any data read is returned in either case.
*/
func NewBufferFromEntireReader(r io.Reader) (*LockedBuffer, error) {
	_ = "STUB: not implemented"
	// Create a buffer with a data region of one page size.
	return nil, nil
}

// Attempt to read some data from the reader.

// Nothing read but no error, try again.

// 1) so either have data and no error
// 2) or have error and no data
// 3) or both have data and have error

// Increment the read count by the number of bytes that we just read.

// Suppress EOF error

// We're done, return the data.

// No data read.

// If we've filled this buffer, grow it by another page size.

/*
NewBufferRandom constructs an immutable buffer filled with cryptographically-secure random bytes.
*/
func NewBufferRandom(size int) *LockedBuffer {
	_ = "STUB: not implemented"
	// Construct a buffer of the specified size.
	return nil
}

// Fill the buffer with random bytes.

// Make the buffer immutable.

// Return the created Buffer object.

// Freeze makes a LockedBuffer's memory immutable. The call can be reversed with Melt.
func (b *LockedBuffer) Freeze() {
	_ = "STUB: not implemented"

	// Melt makes a LockedBuffer's memory mutable. The call can be reversed with Freeze.
	return
}

func (b *LockedBuffer) Melt() {
	_ = "STUB: not implemented"

	/*
	   Seal takes a LockedBuffer object and returns its contents encrypted inside a sealed Enclave object. The LockedBuffer is subsequently destroyed and its contents wiped.

	   If Seal is called on a destroyed buffer, a nil enclave is returned.
	*/return
}

func (b *LockedBuffer) Seal() *Enclave { _ = "STUB: not implemented"; return nil }

/*
Copy performs a time-constant copy into a LockedBuffer. Move is preferred if the source is not also a LockedBuffer or if the source is no longer needed.
*/
func (b *LockedBuffer) Copy(src []byte) {
	_ = "STUB: not implemented"

	/*
	   CopyAt performs a time-constant copy into a LockedBuffer at an offset. Move is preferred if the source is not also a LockedBuffer or if the source is no longer needed.
	*/return
}

func (b *LockedBuffer) CopyAt(offset int, src []byte) { _ = "STUB: not implemented"; return }

/*
Move performs a time-constant move into a LockedBuffer. The source is wiped after the bytes are copied.
*/
func (b *LockedBuffer) Move(src []byte) {
	_ = "STUB: not implemented"

	/*
	   MoveAt performs a time-constant move into a LockedBuffer at an offset. The source is wiped after the bytes are copied.
	*/return
}

func (b *LockedBuffer) MoveAt(offset int, src []byte) { _ = "STUB: not implemented"; return }

/*
Scramble attempts to overwrite the data with cryptographically-secure random bytes.
*/
func (b *LockedBuffer) Scramble() { _ = "STUB: not implemented"; return }

/*
Wipe attempts to overwrite the data with zeros.
*/
func (b *LockedBuffer) Wipe() { _ = "STUB: not implemented"; return }

/*
Size gives you the length of a given LockedBuffer's data segment. A destroyed LockedBuffer will have a size of zero.
*/
func (b *LockedBuffer) Size() int { _ = "STUB: not implemented"; return 0 }

/*
Destroy wipes and frees the underlying memory of a LockedBuffer. The LockedBuffer will not be accessible or usable after this calls is made.
*/
func (b *LockedBuffer) Destroy() {
	_ = "STUB: not implemented"

	/*
	   IsAlive returns a boolean value indicating if a LockedBuffer is alive, i.e. that it has not been destroyed.
	*/return
}

func (b *LockedBuffer) IsAlive() bool { _ = "STUB: not implemented"; return false }

/*
IsMutable returns a boolean value indicating if a LockedBuffer is mutable.
*/
func (b *LockedBuffer) IsMutable() bool { _ = "STUB: not implemented"; return false }

/*
EqualTo performs a time-constant comparison on the contents of a LockedBuffer with a given buffer. A destroyed LockedBuffer will always return false.
*/
func (b *LockedBuffer) EqualTo(buf []byte) bool { _ = "STUB: not implemented"; return false }

/*
	Functions for representing the memory region as various data types.
*/

/*
Bytes returns a byte slice referencing the protected region of memory.
*/
func (b *LockedBuffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

/*
Reader returns a Reader object referencing the protected region of memory.
*/
func (b *LockedBuffer) Reader() *bytes.Reader { _ = "STUB: not implemented"; return nil }

/*
String returns a string representation of the protected region of memory.
*/
func (b *LockedBuffer) String() string { _ = "STUB: not implemented"; return "" }

/*
Uint16 returns a slice pointing to the protected region of memory with the data represented as a sequence of unsigned 16 bit integers. Its length will be half that of the byte slice, excluding any remaining part that doesn't form a complete uint16 value.

If called on a destroyed LockedBuffer, a nil slice will be returned.
*/
func (b *LockedBuffer) Uint16() []uint16 {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Compute size of new slice representation.

// Construct the new slice representation.

// Cast the representation to the correct type and return it.

/*
Uint32 returns a slice pointing to the protected region of memory with the data represented as a sequence of unsigned 32 bit integers. Its length will be one quarter that of the byte slice, excluding any remaining part that doesn't form a complete uint32 value.

If called on a destroyed LockedBuffer, a nil slice will be returned.
*/
func (b *LockedBuffer) Uint32() []uint32 {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Compute size of new slice representation.

// Construct the new slice representation.

// Cast the representation to the correct type and return it.

/*
Uint64 returns a slice pointing to the protected region of memory with the data represented as a sequence of unsigned 64 bit integers. Its length will be one eighth that of the byte slice, excluding any remaining part that doesn't form a complete uint64 value.

If called on a destroyed LockedBuffer, a nil slice will be returned.
*/
func (b *LockedBuffer) Uint64() []uint64 {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Compute size of new slice representation.

// Construct the new slice representation.

// Cast the representation to the correct type and return it.

/*
Int8 returns a slice pointing to the protected region of memory with the data represented as a sequence of signed 8 bit integers. If called on a destroyed LockedBuffer, a nil slice will be returned.
*/
func (b *LockedBuffer) Int8() []int8 {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Construct the new slice representation.

// Cast the representation to the correct type and return it.

/*
Int16 returns a slice pointing to the protected region of memory with the data represented as a sequence of signed 16 bit integers. Its length will be half that of the byte slice, excluding any remaining part that doesn't form a complete int16 value.

If called on a destroyed LockedBuffer, a nil slice will be returned.
*/
func (b *LockedBuffer) Int16() []int16 {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Compute size of new slice representation.

// Construct the new slice representation.

// Cast the representation to the correct type and return it.

/*
Int32 returns a slice pointing to the protected region of memory with the data represented as a sequence of signed 32 bit integers. Its length will be one quarter that of the byte slice, excluding any remaining part that doesn't form a complete int32 value.

If called on a destroyed LockedBuffer, a nil slice will be returned.
*/
func (b *LockedBuffer) Int32() []int32 {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Compute size of new slice representation.

// Construct the new slice representation.

// Cast the representation to the correct type and return it.

/*
Int64 returns a slice pointing to the protected region of memory with the data represented as a sequence of signed 64 bit integers. Its length will be one eighth that of the byte slice, excluding any remaining part that doesn't form a complete int64 value.

If called on a destroyed LockedBuffer, a nil slice will be returned.
*/
func (b *LockedBuffer) Int64() []int64 {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Compute size of new slice representation.

// Construct the new slice representation.

// Cast the representation to the correct type and return it.

/*
ByteArray8 returns a pointer to some 8 byte array. Care must be taken not to dereference the pointer and instead pass it around as-is.

The length of the buffer must be at least 8 bytes in size and the LockedBuffer should not be destroyed. In either of these cases a nil value is returned.
*/
func (b *LockedBuffer) ByteArray8() *[8]byte {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Check if the length is large enough.

// Cast the representation to the correct type.

/*
ByteArray16 returns a pointer to some 16 byte array. Care must be taken not to dereference the pointer and instead pass it around as-is.

The length of the buffer must be at least 16 bytes in size and the LockedBuffer should not be destroyed. In either of these cases a nil value is returned.
*/
func (b *LockedBuffer) ByteArray16() *[16]byte {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Check if the length is large enough.

// Cast the representation to the correct type.

/*
ByteArray32 returns a pointer to some 32 byte array. Care must be taken not to dereference the pointer and instead pass it around as-is.

The length of the buffer must be at least 32 bytes in size and the LockedBuffer should not be destroyed. In either of these cases a nil value is returned.
*/
func (b *LockedBuffer) ByteArray32() *[32]byte {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Check if the length is large enough.

// Cast the representation to the correct type.

/*
ByteArray64 returns a pointer to some 64 byte array. Care must be taken not to dereference the pointer and instead pass it around as-is.

The length of the buffer must be at least 64 bytes in size and the LockedBuffer should not be destroyed. In either of these cases a nil value is returned.
*/
func (b *LockedBuffer) ByteArray64() *[64]byte {
	_ = "STUB: not implemented"

	// Check if still alive.
	return nil
}

// Check if the length is large enough.

// Cast the representation to the correct type.
