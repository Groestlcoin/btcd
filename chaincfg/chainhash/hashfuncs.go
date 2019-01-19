// Copyright (c) 2015 The Decred developers
// Copyright (c) 2016-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chainhash

import (
	"crypto/sha256"
	"github.com/Groestlcoin/go-groestl-hash/groestl"
)

// HashB calculates hash(b) and returns the resulting bytes.
func HashB(b []byte) []byte {
	hash := sha256.Sum256(b)
	return hash[:]
}

// HashH calculates hash(b) and returns the resulting bytes as a Hash.
func HashH(b []byte) Hash {
	return Hash(sha256.Sum256(b))
}

// DoubleHashB calculates hash(hash(b)) and returns the resulting bytes.
func DoubleHashB(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

// DoubleHashH calculates hash(hash(b)) and returns the resulting bytes as a
// Hash.
func DoubleHashH(b []byte) Hash {
	first := sha256.Sum256(b)
	return Hash(sha256.Sum256(first[:]))
}

func DoubleGroestlB(b []byte) []byte {
	var h1, h2 [64]byte
	g := groestl.New()
	g.Write(b)
	g.Close(h1[:], 0, 0)
	g.Write(h1[:])
	g.Close(h2[:], 0, 0)
	return h2[:32]
}
func DoubleGroestlH(b []byte) (ret Hash) {
	var h1, h2 [64]byte
	g := groestl.New()
	g.Write(b)
	g.Close(h1[:], 0, 0)
	g.Write(h1[:])
	g.Close(h2[:], 0, 0)
	copy(ret[:], h2[:32])
	return
}
