// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
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
		},
	},
	TxOut: []*wire.TxOut{
		{
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
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x23, 0x90, 0x63, 0x3b, 0x70, 0xf0, 0x62, 0xcb,
	0x3a, 0x3d, 0x68, 0x14, 0xb6, 0x7e, 0x29, 0xa8,
	0x0d, 0x9d, 0x75, 0x81, 0xdb, 0x0b, 0xcc, 0x49,
	0x4d, 0x59, 0x7c, 0x92, 0xc5, 0x0a, 0x00, 0x00,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xbb, 0x28, 0x66, 0xaa, 0xca, 0x46, 0xc4, 0x42,
	0x8a, 0xd0, 0x8b, 0x57, 0xbc, 0x9d, 0x14, 0x93,
	0xab, 0xaf, 0x64, 0x72, 0x4b, 0x6c, 0x30, 0x52,
	0xa7, 0xc8, 0xf9, 0x58, 0xdf, 0x68, 0xe9, 0x3c,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    112,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // 3ce968df58f9c8a752306c4b7264afab93149dbc578bd08a42c446caaa6628bb
		Timestamp:  time.Unix(1395342829, 0), // Thu, 20 Mar 2014 19:13:49 UTC
		Bits:       0x1e0fffff,
		Nonce:      220035,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x36, 0xcd, 0xf2, 0xdc, 0xb7, 0x55, 0x62, 0x87,
	0x28, 0x2a, 0x05, 0xc0, 0x64, 0x01, 0x23, 0x23,
	0xba, 0xe6, 0x63, 0xc1, 0x6e, 0xd3, 0xcd, 0x98,
	0x98, 0xfc, 0x50, 0xbb, 0xff, 0x00, 0x00, 0x00,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    3,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 3ce968df58f9c8a752306c4b7264afab93149dbc578bd08a42c446caaa6628bb
		Timestamp:  time.Unix(1440000002, 0), // Wed 19 Aug 16:00:02 UTC 2015
		Bits:       0x1e00ffff,
		Nonce:      6556309,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// test network (version 3).
var testNet3GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x36, 0xcd, 0xf2, 0xdc, 0xb7, 0x55, 0x62, 0x87,
	0x28, 0x2a, 0x05, 0xc0, 0x64, 0x01, 0x23, 0x23,
	0xba, 0xe6, 0x63, 0xc1, 0x6e, 0xd3, 0xcd, 0x98,
	0x98, 0xfc, 0x50, 0xbb, 0xff, 0x00, 0x00, 0x00,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNet3GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    3,
		PrevBlock:  chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet3GenesisMerkleRoot, // 3ce968df58f9c8a752306c4b7264afab93149dbc578bd08a42c446caaa6628bb
		Timestamp:  time.Unix(1440000002, 0),  // Wed 19 Aug 16:00:02 UTC 2015
		Bits:       0x1e00ffff,
		Nonce:      6556309,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x16, 0x44, 0xca, 0x92, 0xa1, 0xab, 0x60, 0xa7,
	0xb4, 0x53, 0xd9, 0x7c, 0x66, 0x5f, 0x48, 0x9a,
	0xd0, 0x61, 0xdf, 0x79, 0xc5, 0x20, 0x73, 0xd8,
	0xdd, 0x69, 0xc0, 0x78, 0x89, 0xa5, 0x0b, 0xf2,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 3ce968df58f9c8a752306c4b7264afab93149dbc578bd08a42c446caaa6628bb
		Timestamp:  time.Unix(1551661551, 0), // Mon  4 Mar 01:05:51 UTC 2019
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// sigNetGenesisHash is the hash of the first block in the block chain for the
// signet test network.
var sigNetGenesisHash = chainhash.Hash{
	0x31, 0xab, 0x14, 0xbb, 0x92, 0x35, 0xf2, 0xa2, 0xeb, 0x6c, 0x87, 0x7b,
	0x51, 0xaf, 0x57, 0x43, 0x25, 0x8c, 0x81, 0xe7, 0xe9, 0xcd, 0xc6, 0x93,
	0x79, 0xa2, 0xa2, 0xca, 0x7f, 0x00, 0x00, 0x00,
}

// sigNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the signet test network. It is the same as the merkle root for
// the main network.
var sigNetGenesisMerkleRoot = genesisMerkleRoot

// sigNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the signet test network.
var sigNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    3,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: sigNetGenesisMerkleRoot,  // 3ce968df58f9c8a752306c4b7264afab93149dbc578bd08a42c446caaa6628bb
		Timestamp:  time.Unix(1606082400, 0), // Sun 22 Nov 22:00:00 UTC 2020
		Bits:       0x1e00ffff,
		Nonce:      14675970,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
