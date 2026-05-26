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

package stream

// SlowRandByte writes 16KiB of random data to a stream and then operates on it in chunks, returning a random number between 0 and 255.
func SlowRandByte() byte {
	_ = "STUB: not implemented"
	// Get 16KiB bytes of random data.
	// In the real world we might be reading from a socket instead.
	// Also we are free to write data in arbitrarily sized chunks.
	return 0
}

// Allow mutation so stream writer can wipe source buffer.

// Create a stream object.
// Implements io.Reader and io.Writer interfaces.

// Write the data to it.
// Should never error or write less data.
// No longer need the source buffer. (Has been wiped.)

// Create a buffer to work on this data in chunks.

// Read the data back in chunks.

// Reads directly into guarded allocation.

// end of data

// other error

// Do some example computation on this data.

// Return the result.
