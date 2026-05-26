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

package socketkey

// Save the data here so we can compare it later. Obviously this leaks the secret.
var data []byte

// NOTE: Some lines are commented out for the sake of tests.

/*
SocketKey is a streaming multi-threaded client->server transfer of secure data over a socket.
*/
func SocketKey(size int) {
	_ = "STUB: not implemented"
	// Create a server to listen on.
	return
}

// Catch signals and close the listener before terminating safely.

// Purge the session before returning.

// Create a client to connect to our server.

// Connect to our server

// Create a buffer filled with random bytes

// Save a copy of the key for comparison later.

// fmt.Printf("Sending key: %#v\n", buf.Bytes())

// Send the data to the server

// Accept connections from clients

// Read the data directly into a guarded memory region

// fmt.Printf("Received key: %#v\n", buf.Bytes())

// Compare the key to make sure it wasn't corrupted.

// Seal the key into an encrypted Enclave object.

// <-- buf is destroyed by this point

// fmt.Printf("Encrypted key: %#v\n", key)

// Decrypt the key into a new buffer.

// fmt.Printf("Decrypted key: %#v\n", buf.Bytes())

// Destroy the buffer.
