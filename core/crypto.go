package core

import (
	"errors"

	"golang.org/x/crypto/nacl/secretbox"
)

// Overhead is the size by which the ciphertext exceeds the plaintext.
const Overhead int = secretbox.Overhead + 24 // auth + nonce

// ErrInvalidKeyLength is returned when attempting to encrypt or decrypt with a key that is not exactly 32 bytes in size.
var ErrInvalidKeyLength = errors.New("<memguard::core::ErrInvalidKeyLength> key must be exactly 32 bytes")

// ErrBufferTooSmall is returned when the decryption function, Open, is given an output buffer that is too small to hold the plaintext. In practice the plaintext will be Overhead bytes smaller than the ciphertext returned by the encryption function, Seal.
var ErrBufferTooSmall = errors.New("<memguard::core::ErrBufferTooSmall> the given buffer is too small to hold the plaintext")

// ErrDecryptionFailed is returned when the attempted decryption fails. This can occur if the given key is incorrect or if the ciphertext is invalid.
var ErrDecryptionFailed = errors.New("<memguard::core::ErrDecryptionFailed> decryption failed: key is wrong or ciphertext is corrupt")

// Encrypt takes a plaintext message and a 32 byte key and returns an authenticated ciphertext.
func Encrypt(plaintext, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Check the length of the key is correct.
	return nil, nil
}

// Get a reference to the key's underlying array without making a copy.

// Allocate space for and generate a nonce value.

// Encrypt m and return the result.

/*
Decrypt decrypts a given ciphertext with a given 32 byte key and writes the result to the start of a given buffer.

The buffer must be large enough to contain the decrypted data. This is in practice Overhead bytes less than the length of the ciphertext returned by the Seal function above. This value is the size of the nonce plus the size of the Poly1305 authenticator.

The size of the decrypted data is returned.
*/
func Decrypt(ciphertext, key []byte, output []byte) (int, error) {
	_ = "STUB: not implemented"
	// Check the length of the key is correct.
	return 0, nil
}

// Check the capacity of the given output buffer.

// Get a reference to the key's underlying array without making a copy.

// Retrieve and store the nonce value.

// Decrypt and return the result.

// Decryption successful.
// Move plaintext to given output buffer.
// Return length of decrypted plaintext.

// Decryption unsuccessful. Either the key was wrong or the authentication failed.

// Hash implements a cryptographic hash function using Blake2b.
func Hash(b []byte) []byte { _ = "STUB: not implemented"; return nil }

// Scramble fills a given buffer with cryptographically-secure random bytes.
func Scramble(buf []byte) error { _ = "STUB: not implemented"; return nil }

// See Wipe

// Wipe takes a buffer and wipes it with zeroes.
func Wipe(buf []byte) { _ = "STUB: not implemented"; return }

// This should keep buf's backing array live and thus prevent dead store
// elimination, according to discussion at
// https://github.com/golang/go/issues/33325 .

// Copy is identical to Go's builtin copy function except the copying is done in constant time. This is to mitigate against side-channel attacks.
func Copy(dst, src []byte) { _ = "STUB: not implemented"; return }

// Move is identical to Copy except it wipes the source buffer after the copy operation is executed.
func Move(dst, src []byte) { _ = "STUB: not implemented"; return }

// Equal does a constant-time comparison of two byte slices. This is to mitigate against side-channel attacks.
func Equal(x, y []byte) bool { _ = "STUB: not implemented"; return false }
