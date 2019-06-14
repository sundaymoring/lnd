package input

import (
	"encoding/binary"
	"io"

	"github.com/btcsuite/btcd/wire"
)

// writeTxOut serializes a wire.TxOut struct into the passed io.Writer stream.
func writeTxOut(w io.Writer, txo *wire.TxOut) error {
	var scratch [8]byte

	binary.BigEndian.PutUint64(scratch[:], uint64(txo.Value))
	if _, err := w.Write(scratch[:]); err != nil {
		return err
	}

	if err := wire.WriteVarBytes(w, 0, txo.PkScript); err != nil {
		return err
	}

	if err := wire.WriteVarBytes(w, 0, txo.TokenId[:]); err != nil {
		return err
	}

	var tokkenAmount [8]byte
	binary.BigEndian.PutUint64(tokkenAmount[:], uint64(txo.TokenValue))
	if _, err := w.Write(tokkenAmount[:]); err != nil {
		return err
	}

	return nil
}

// readTxOut deserializes a wire.TxOut struct from the passed io.Reader stream.
func readTxOut(r io.Reader, txo *wire.TxOut) error {
	var scratch [8]byte

	if _, err := io.ReadFull(r, scratch[:]); err != nil {
		return err
	}
	value := int64(binary.BigEndian.Uint64(scratch[:]))

	pkScript, err := wire.ReadVarBytes(r, 0, 80, "pkScript")
	if err != nil {
		return err
	}

	tokenId, err := wire.ReadVarBytes(r, 0, 80, "tokenId")
	if err != nil {
		return err
	}

	var tokenAmountBytes [8]byte

	if _, err := io.ReadFull(r, tokenAmountBytes[:]); err != nil {
		return err
	}
	tokenAmount := int64(binary.BigEndian.Uint64(tokenAmountBytes[:]))

	*txo = wire.TxOut{
		Value:    value,
		PkScript: pkScript,
		TokenValue: tokenAmount,
	}
	txo.TokenId.SetBytes(tokenId)

	return nil
}
