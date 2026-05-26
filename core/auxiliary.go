package core

import (
	"os"
)

var (
	// Ascertain and store the system memory page size.
	pageSize = os.Getpagesize()
)

// Round a length to a multiple of the system page size.
func roundToPageSize(length int) int { _ = "STUB: not implemented"; return 0 }
