// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hash provides interfaces for hash functions.
package BitSHA

import "io"

// Hash is the common interface implemented by all hash functions.
type Hash interface {
	// Write adds more data to the running hash.
	// It never returns an error.
	io.Writer

	// Sum returns the current hash, without changing the
	// underlying hash state.
	Sum() []byte

	// Reset resets the hash to one with zero bytes written.
	Reset()
	
	//Gives a copy of the digest structure
	Copy() Hash

	//Calculates the midstate of a SHA hash
	Midstate() []byte

	// Size returns the number of bytes Sum will return.
	Size() int
}

// Hash32 is the common interface implemented by all 32-bit hash functions.
type Hash32 interface {
	Hash
	Sum32() uint32
}

// Hash64 is the common interface implemented by all 64-bit hash functions.
type Hash64 interface {
	Hash
	Sum64() uint64
}
