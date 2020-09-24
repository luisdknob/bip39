package main

import (
	"crypto/sha512"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func main() {

	length := len(os.Args) - 1
	password := ""
	if length%3 != 0 {
		length--
		password = os.Args[length+1]
	}

	if length >= 12 && length <= 24 && length%3 == 0 {

		seedString := ""
		for _, element := range os.Args[1 : length+1] {
			seedString += element + " "
		}
		seedString = strings.TrimSuffix(seedString, " ")

		seed := pbkdf2.Key([]byte(seedString), []byte("mnemonic"+password), 2048, 64, sha512.New)
		fmt.Printf("%x", seed)
		return

	}

	fmt.Printf("bip39Seed SEEDPHRASE [PASSWORD]\n")
	return
}
