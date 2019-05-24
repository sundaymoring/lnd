// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package helpers provides convenience functions to simplify wallet code.  This
// package is intended for internal wallet use only.
package helpers

import (
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

// SumOutputValues sums up the list of TxOuts and returns an Amount.
func SumOutputValues(outputs []*wire.TxOut) (btcutil.Amount, *wire.TokenId, btcutil.Amount) {

	var totalOutput, totalTokenOutput btcutil.Amount
	tokenId := &wire.TokenId{}
	for _, txOut := range outputs {
		totalOutput += btcutil.Amount(txOut.Value)
		if txOut.TokenId.IsValid() {
			if !tokenId.IsValid() {
				tokenId.SetBytes(txOut.TokenId[:])
			}

			totalTokenOutput += btcutil.Amount(txOut.TokenValue)
		}
	}
	return totalOutput, tokenId, totalTokenOutput
}

// SumOutputSerializeSizes sums up the serialized size of the supplied outputs.
func SumOutputSerializeSizes(outputs []*wire.TxOut) (serializeSize int) {
	for _, txOut := range outputs {
		serializeSize += txOut.SerializeSize()
	}
	return serializeSize
}
