package main

import (
	"encoding/hex"
	"fmt"

	btcspv "github.com/summa-tx/bitcoin-spv/golang/btcspv"
)

func prettifyInput(numInput int, outpoint []byte, index uint, inputType btcspv.InputType, sequence uint) string {
	outpointStr := hex.EncodeToString(outpoint)
	dataStr := fmt.Sprintf("\nInput #%d:\n  Outpoint: %s,\n  Index: %d,\n  Type: %d,\n  Sequence: %d\n", numInput, outpointStr, index, inputType, sequence)
	return dataStr
}

// ParseVin parses an input vector from hex
func ParseVin(vin []byte) string {
	// Validate the vin
	isVin := btcspv.ValidateVin(vin)
	if !isVin {
		return "Invalid Vin"
	}

	numInputs := int(vin[0])
	var inputs string
	for i := 0; i < numInputs; i++ {
		// Extract each vin at the specified index
		vin := btcspv.ExtractInputAtIndex(vin, uint8(i))

		// Use ParseInput to get more information about the vin
		sequence, inputID, inputIndex, inputType := btcspv.ParseInput(vin)

		// Format information about the vin
		numInput := i + 1
		vinData := prettifyInput(numInput, inputID, inputIndex, inputType, sequence)

		// Concat vin information onto `inputs`
		inputs = inputs + vinData
	}

	return inputs
}