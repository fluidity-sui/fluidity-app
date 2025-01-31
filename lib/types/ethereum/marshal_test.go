// Copyright 2022 Fluidity Money. All rights reserved. Use of this
// source code is governed by a GPL-style license that can be found in the
// LICENSE.md file.

package ethereum

import (
	"encoding/json"
	"testing"

	"github.com/fluidity-money/fluidity-app/lib/types/misc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type StructTest struct {
	HashKey    Hash
	AddressKey Address
	HashMap    map[Hash]Hash
	AddressMap map[Address]Address
	HashArr    []Hash
	AddressArr []Address
}

type MarshalTestSuite struct {
	suite.Suite
	testHeader  BlockHeader
	testTx      Transaction
	testReceipt Receipt
	testStruct  StructTest
}

func (suite *MarshalTestSuite) SetupTest() {
	testHeader := BlockHeader{
		ParentHash:      HashFromString("0xparent"),
		UncleHash:       HashFromString("0xuncle"),
		Coinbase:        AddressFromString("coin"),
		Root:            HashFromString("root"),
		TransactionHash: HashFromString("0xtxn"),
		Bloom:           misc.Blob{1, 2, 3},
		Difficulty:      misc.BigIntFromInt64(100),
		Number:          misc.BigIntFromInt64(100),
		GasLimit:        misc.BigIntFromInt64(100),
		GasUsed:         misc.BigIntFromInt64(100),
		Time:            100,
		Extra:           misc.Blob{4, 5, 6},
		MixDigest:       HashFromString("mix"),
		Nonce:           BlockNonce{1, 2, 3},
		ReceiptHash:     HashFromString("receipt"),
		BaseFee:         misc.BigIntFromInt64(20),
	}

	var (
		receiptBlockNumber      = misc.BigIntFromInt64(0)
		receiptIndex            = misc.BigIntFromInt64(0)
		receiptTransactionIndex = misc.BigIntFromInt64(2)
	)

	testReceipt := Receipt{
		Type:              1,
		PostState:         misc.Blob{1},
		Status:            0,
		CumulativeGasUsed: 0,
		Bloom:             misc.Blob{1, 2},
		Logs: []Log{
			{
				Address:     AddressFromString("logaddr"),
				Topics:      []Hash{HashFromString("topic1"), HashFromString("topic2")},
				Data:        misc.Blob{100, 255},
				BlockNumber: receiptBlockNumber,
				TxHash:      HashFromString("12"),
				TxIndex:     receiptTransactionIndex,
				BlockHash:   HashFromString("blockhash"),
				Index:       receiptIndex,
				Removed:     false,
			},
		},
		TransactionHash:  HashFromString("0xtxn"),
		ContractAddress:  AddressFromString("address"),
		GasUsed:          misc.BigIntFromInt64(123),
		BlockHash:        HashFromString("0xblock"),
		BlockNumber:      misc.BigIntFromInt64(1238),
		TransactionIndex: 0,
	}

	testTx := Transaction{
		BlockHash: HashFromString("0xblock"),
		Data:      misc.Blob{1, 2},
		GasFeeCap: misc.BigIntFromInt64(100),
		GasTipCap: misc.BigIntFromInt64(100),
		GasPrice:  misc.BigIntFromInt64(100),
		Hash:      HashFromString("hash"),
		To:        AddressFromString("toaddr"),
		From:      AddressFromString("fromaddr"),
		Type:      1,
	}

	testStruct := StructTest{
		HashKey:    HashFromString("0xhash"),
		AddressKey: AddressFromString("0xaddress"),
		HashMap: map[Hash]Hash{
			HashFromString("0xhash1"): HashFromString("0xvalue1"),
			HashFromString("0xhash2"): HashFromString("0xvalue2"),
		},
		AddressMap: map[Address]Address{
			AddressFromString("0xaddress1"): AddressFromString("0xvalue1"),
			AddressFromString("0xaddress2"): AddressFromString("0xvalue2"),
		},
		HashArr: []Hash{
			HashFromString("0xhash0"),
			HashFromString("0xhash1"),
		},
		AddressArr: []Address{
			AddressFromString("0xaddress0"),
			AddressFromString("0xaddress1"),
		},
	}

	suite.testHeader = testHeader
	suite.testTx = testTx
	suite.testReceipt = testReceipt
	suite.testStruct = testStruct
}

func TestBlobTestSuite(t *testing.T) {
	suite.Run(t, new(MarshalTestSuite))
}

func (suite *MarshalTestSuite) TestMarshal() {

	suite.T().Run("TestMarshalHeader", func(t *testing.T) {
		header := suite.testHeader

		bin, err := header.MarshalBinary()
		require.NoError(t, err)

		var newHeader BlockHeader
		json.Unmarshal(bin, &newHeader)

		assert.Equal(t, header, newHeader)
	})

	suite.T().Run("TestMarshalTransaction", func(t *testing.T) {
		txn := suite.testTx

		bin, err := txn.MarshalBinary()
		require.NoError(t, err)

		var newTxn Transaction
		json.Unmarshal(bin, &newTxn)

		assert.Equal(t, txn, newTxn)
	})

	suite.T().Run("TestMarshalLog", func(t *testing.T) {
		log := suite.testReceipt.Logs[0]

		bin, err := log.MarshalBinary()
		require.NoError(t, err)

		var newLog Log
		json.Unmarshal(bin, &newLog)

		assert.Equal(t, log, newLog)
	})

	suite.T().Run("TestMarshalReceipt", func(t *testing.T) {
		receipt := suite.testReceipt

		bin, err := receipt.MarshalBinary()
		require.NoError(t, err)

		var newReceipt Receipt
		json.Unmarshal(bin, &newReceipt)

		assert.Equal(t, receipt, newReceipt)
	})

	suite.T().Run("TestMarshalStruct", func(t *testing.T) {
		testStruct := suite.testStruct

		bin, err := json.Marshal(testStruct)
		require.NoError(t, err)

		var newStruct StructTest
		json.Unmarshal(bin, &newStruct)

		assert.Equal(t, testStruct, newStruct)
	})
}
