package memguard

/* Enhancement: check for low memory locking limit and print warning?*/

/*
ScrambleBytes overwrites an arbitrary buffer with cryptographically-secure random bytes.
*/
func ScrambleBytes(buf []byte) { _ = "STUB: not implemented"; return }

/*
WipeBytes overwrites an arbitrary buffer with zeroes.
*/
func WipeBytes(buf []byte) {
	_ = "STUB: not implemented"

	/*
	   Purge resets the session key to a fresh value and destroys all existing LockedBuffers. Existing Enclave objects will no longer be decryptable.
	*/return
}

func Purge() {
	_ = "STUB: not implemented"

	/*
	   SafePanic wipes all it can before calling panic(v).
	*/return
}

func SafePanic(v any) {
	_ = "STUB: not implemented"

	/*
	   SafeExit destroys everything sensitive before exiting with a specified status code.
	*/return
}

func SafeExit(c int) { _ = "STUB: not implemented"; return }
