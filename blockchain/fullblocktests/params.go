// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package fullblocktests

import (
	"encoding/hex"
	"math/big"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// newHashFromStr converts the passed big-endian hex string into a
// wire.Hash.  It only differs from the one available in chainhash in that
// it panics on an error since it will only (and must only) be called with
// hard-coded, and therefore known good, hashes.
func newHashFromStr(hexStr string) *chainhash.Hash {
	hash, err := chainhash.NewHashFromStr(hexStr)
	if err != nil {
		panic(err)
	}
	return hash
}

// fromHex converts the passed hex string into a byte slice and will panic if
// there is an error.  This is only provided for the hard-coded constants so
// errors in the source code can be detected. It will only (and must only) be
// called for initialization purposes.
func fromHex(s string) []byte {
	r, err := hex.DecodeString(s)
	if err != nil {
		panic("invalid hex in source file: " + s)
	}
	return r
}

var (
	// bigOne is 1 represented as a big.Int.  It is defined here to avoid
	// the overhead of creating it multiple times.
	bigOne = big.NewInt(1)

	// regressionPowLimit is the highest proof of work value a Bitcoin block
	// can have for the regression test network.  It is the value 2^255 - 1.
	regressionPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 255), bigOne)

	// regTestGenesisHash is the hash of the first block in the block chain for the
	// regression test network (genesis block).
	regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
		0x36, 0xcd, 0xf2, 0xdc, 0xb7, 0x55, 0x62, 0x87,
		0x28, 0x2a, 0x05, 0xc0, 0x64, 0x01, 0x23, 0x23,
		0xba, 0xe6, 0x63, 0xc1, 0x6e, 0xd3, 0xcd, 0x98,
		0x98, 0xfc, 0x50, 0xbb, 0xff, 0x00, 0x00, 0x00,
	})

	// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
	// block for the regression test network.  It is the same as the merkle root for
	// the main network.
	regTestGenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
		0xbb, 0x28, 0x66, 0xaa, 0xca, 0x46, 0xc4, 0x42,
		0x8a, 0xd0, 0x8b, 0x57, 0xbc, 0x9d, 0x14, 0x93,
		0xab, 0xaf, 0x64, 0x72, 0x4b, 0x6c, 0x30, 0x52,
		0xa7, 0xc8, 0xf9, 0x58, 0xdf, 0x68, 0xe9, 0x3c,
	})

	// regTestGenesisBlock defines the genesis block of the block chain which serves
	// as the public transaction ledger for the regression test network.
	regTestGenesisBlock = wire.MsgBlock{
		Header: wire.BlockHeader{
			Version:    3,
			PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
			MerkleRoot: regTestGenesisMerkleRoot, // 3ce968df58f9c8a752306c4b7264afab93149dbc578bd08a42c446caaa6628bb
			Timestamp:  time.Unix(1440000002, 0), // Wed 19 Aug 16:00:02 UTC 2015
			Bits:       0x1e00ffff,
			Nonce:      6556309,
		},
		Transactions: []*wire.MsgTx{{
			Version: 1,
			TxIn: []*wire.TxIn{{
				PreviousOutPoint: wire.OutPoint{
					Hash:  chainhash.Hash{},
					Index: 0xffffffff,
				},
				SignatureScript: []byte{
					0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x32,
					0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65,
					0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65,
					0x20, 0x70, 0x75, 0x74, 0x20, 0x6f, 0x6e, 0x20,
					0x56, 0x6c, 0x61, 0x64, 0x69, 0x6d, 0x69, 0x72,
					0x20, 0x50, 0x75, 0x74, 0x69, 0x6e, 0x20, 0x6f,
					0x76, 0x65, 0x72, 0x20, 0x43, 0x72, 0x69, 0x6d,
					0x65, 0x61,
				},
				Sequence: 0xffffffff,
			}},
			TxOut: []*wire.TxOut{{
				Value: 0,
				PkScript: []byte{
					0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
					0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
					0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
					0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
					0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
					0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
					0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
					0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
					0x1d, 0x5f, 0xac, /* |._.| */
				},
			}},
			LockTime: 0,
		}},
	}
)

// regressionNetParams defines the network parameters for the regression test
// network.
//
// NOTE: The test generator intentionally does not use the existing definitions
// in the chaincfg package since the intent is to be able to generate known
// good tests which exercise that code.  Using the chaincfg parameters would
// allow them to change out from under the tests potentially invalidating them.
//
// NOTE: Unfortunately information above is not true for the current btcd
// codebase. These definitions must be in sync with chaincfg package, otherwise
// tests will fail.
var regressionNetParams = &chaincfg.Params{
	Name:        "regtest",
	Net:         wire.TestNet,
	DefaultPort: "18888",

	// Chain parameters
	GenesisBlock:             &regTestGenesisBlock,
	GenesisHash:              &regTestGenesisHash,
	PowLimit:                 regressionPowLimit,
	PowLimitBits:             0x207fffff,
	CoinbaseMaturity:         100,
	BIP0034Height:            100000000,
	BIP0065Height:            1351,
	BIP0066Height:            1251,
	DGW3SwitchHeight:         5000,
	TargetTimespan:           time.Second * 24,
	TargetTimePerBlock:       time.Second,
	RetargetAdjustmentFactor: 3,
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Second * 2, // TargetTimePerBlock * 2
	GenerateSupported:        true,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Mempool parameters
	RelayNonStdTxs: true,

	// Address encoding magics
	PubKeyHashAddrID: 0x6f, // starts with m or n
	ScriptHashAddrID: 0xc4, // starts with 2
	PrivateKeyID:     0xef, // starts with 9 (uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}
