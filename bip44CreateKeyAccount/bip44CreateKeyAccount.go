package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	bip32 "github.com/tyler-smith/go-bip32"
)

const hardenedKeyZeroIndex = 0x80000000
const hardenedKey44Index = 0x8000002C

func main() {

	if len(os.Args) == 4 {

		if seed, err := hex.DecodeString(os.Args[1]); err == nil {
			MasterKey, _ := bip32.Deserialize([]byte(seed))

			change, _ := strconv.ParseUint(os.Args[2], 10, 64)
			index, _ := strconv.ParseUint(os.Args[3], 10, 64)

			bip44Purpose, _ := MasterKey.NewChildKey(uint32(hardenedKey44Index))
			bip44Coin, _ := bip44Purpose.NewChildKey(uint32(hardenedKeyZeroIndex))
			bip44Account, _ := bip44Coin.NewChildKey(uint32(hardenedKeyZeroIndex))
			bip44Internal, _ := bip44Account.NewChildKey(uint32(change))
			bip44CC, _ := bip44Internal.NewChildKey(uint32(index))

			mainNetAddr, _ := btcutil.NewAddressPubKey(bip44CC.PublicKey().Key, &chaincfg.MainNetParams)

			fmt.Println(mainNetAddr.EncodeAddress())

			return
		}
	}

	fmt.Printf("bip44CreateKeyAccount B58 CHANGE INDEX")
	return
}
