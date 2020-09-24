package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/tyler-smith/go-bip32"
)

func main() {

	if len(os.Args) == 2 {
		if seed, err := hex.DecodeString(os.Args[1]); err == nil {

			MasterKey, _ := bip32.NewMasterKey(seed)
			b58seed, _ := MasterKey.Serialize()
			fmt.Printf("%x\n", b58seed)
			return
		}
	}

	fmt.Printf("bip32RootKey SEED")
	return
}
