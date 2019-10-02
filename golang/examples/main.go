package main

import (
	"fmt"
	"os"

	btcspv "github.com/summa-tx/bitcoin-spv/golang/btcspv"
)

func route(command string, arguments [][]byte) string {
	var result string

	switch command {
	case "parseVin":
		result = ParseVin(arguments[0])
	case "parseVout":
		result = ParseVout(arguments[0])
	default:
		result = fmt.Sprintf("Unknown command: %s", command)
	}

	return result
}

// Map function to slice of strings
func Map(vs []string, f func(string) []byte) [][]byte {
	vsm := make([][]byte, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var result string

	if len(os.Args) < 2 {
		fmt.Print("Not enough arguments")
		return
	}

	command := os.Args[1]
	arguments := Map(os.Args[2:], btcspv.DecodeIfHex)

	result = route(command, arguments)
	fmt.Print(result)
}