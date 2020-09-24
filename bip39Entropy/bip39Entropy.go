package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 2 {
		if bitSize, err := strconv.ParseInt(os.Args[1], 10, 64); err == nil {
			if bitSize >= 128 && bitSize <= 256 && bitSize%32 == 0 {
				entropy := make([]byte, bitSize/8)
				_, err = rand.Read(entropy)
				if err != nil {
					return
				}
				fmt.Printf("%x\n", entropy)
				return
			}
		}
	}

	fmt.Printf("bip39Entropy ENTSIZE\nENTSIZE = 128 160 192 224 256")
	return
}
