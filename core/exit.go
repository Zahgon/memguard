package core

/*
Purge wipes all sensitive data and keys before reinitialising the session with a fresh encryption key and secure values. Subsequent library operations will use these fresh values and the old data is assumed to be practically unrecoverable.

The creation of new Enclave objects should wait for this function to return since subsequent Enclave objects will use the newly created key.

This function should be called before the program terminates, or else the provided Exit or Panic functions should be used to terminate.
*/
func Purge() { _ = "STUB: not implemented"; return }

// Halt the re-key cycle and prevent new enclaves or keys being created.

// Get a snapshot of existing Buffers.

// Destroy them, performing the usual sanity checks.

// buffer destroy failed; wipe instead

// couldn't change it to mutable; we can't wipe it! (could this happen?)
// not sure what we can do at this point, just warn and move on

// wipe in subprocess?

// If we encountered an error, panic.

/*
Exit terminates the process with a specified exit code but securely wipes and cleans up sensitive data before doing so.
*/
func Exit(c int) {
	_ = "STUB: not implemented"
	// Wipe the encryption key used to encrypt data inside Enclaves.
	return
}

// Get a snapshot of existing Buffers.
// copy ensures the buffers stay in the list until they are destroyed.

// Destroy them, performing the usual sanity checks.

// Exit with the specified exit code.

/*
Panic is identical to the builtin panic except it purges the session before calling panic.
*/
func Panic(v any) {
	_ = "STUB: not implemented"
	// creates a new key so it is safe to recover from this panic
	return
}
