package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func addChecksum(data []byte, hash []byte) []byte {

	firstChecksumByte := hash[0]

	// len() is in bytes so we divide by 4
	checksumBitLength := uint(len(data) / 4)

	// For each bit of check sum we want we shift the data one the left
	// and then set the (new) right most bit equal to checksum bit at that index
	// staring from the left
	dataBigInt := new(big.Int).SetBytes(data)

	for i := uint(0); i < checksumBitLength; i++ {
		// Bitshift 1 left
		dataBigInt.Mul(dataBigInt, big.NewInt(2))

		// Set rightmost bit if leftmost checksum bit is set
		if uint8(firstChecksumByte&(1<<(7-i))) > 0 {
			dataBigInt.Or(dataBigInt, big.NewInt(1))
		}
	}

	return dataBigInt.Bytes()
}

func padByteSlice(slice []byte, length int) []byte {
	offset := length - len(slice)
	if offset <= 0 {
		return slice
	}
	newSlice := make([]byte, length)
	copy(newSlice[offset:], slice)
	return newSlice
}

func main() {

	if len(os.Args) == 2 {
		if entropy, err := hex.DecodeString(os.Args[1]); err == nil {
			if len(entropy) >= 16 && len(entropy) <= 32 {
				//fmt.Printf("%x\n", entropy)

				sentenceLength := int((len(entropy)*8 + len(entropy)/4) / 11)

				hasher := sha256.New()
				hasher.Write(entropy)
				hash := hasher.Sum(nil)

				entropy = addChecksum(entropy, hash)

				file, err := os.Open("english.txt")
				if err != nil {
					return
				}
				defer file.Close()

				var wordList []string
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					wordList = append(wordList, scanner.Text())
					if scanner.Err() != nil {
						return
					}
				}

				entropyInt := new(big.Int).SetBytes(entropy)

				// Slice to hold words in.
				words := make([]string, sentenceLength)

				// Throw away big.Int for AND masking.
				word := big.NewInt(0)

				for i := sentenceLength - 1; i >= 0; i-- {
					// Get 11 right most bits and bitshift 11 to the right for next time.
					word.And(entropyInt, big.NewInt(2047))
					entropyInt.Div(entropyInt, big.NewInt(2048))

					// Get the bytes representing the 11 bits as a 2 byte slice.
					wordBytes := padByteSlice(word.Bytes(), 2)

					// Convert bytes to an index and add that word to the list.
					words[i] = wordList[binary.BigEndian.Uint16(wordBytes)]
				}

				fmt.Println(strings.Join(words, " "))

				return
			}
		}
	}

	fmt.Printf("bip39Mnemonic ENT\n")
	return
}
