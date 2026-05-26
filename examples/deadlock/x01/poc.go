package x01

import (
	"context"

	"github.com/awnumar/memguard"
)

func OpenEnclave(ctx context.Context) { _ = "STUB: not implemented"; return }

// buf := make([]byte, 1<<20)
// fmt.Println(string(buf[:runtime.Stack(buf, true)]))

func openVerify(lock *memguard.Enclave, exp []byte) error { _ = "STUB: not implemented"; return nil }

func immediateOpen(ctx context.Context, lock *memguard.Enclave, exp []byte) {
	_ = "STUB: not implemented"
	return
}
