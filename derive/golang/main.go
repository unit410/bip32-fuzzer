package main

import (
	"fmt"

	"flag"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var hdPath, mnemonic, passphrase string
	flag.Parse()
	args := flag.Args()

	hdPath = args[0]
	mnemonic = args[1]
	if len(args) > 2 {
		passphrase = args[2]
	}
	fmt.Println(deriveAddress(hdPath, mnemonic, passphrase))
}

func deriveAddress(hdPath string, mnemonic string, passphrase string) string {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, passphrase)
	check(err)

	wallet, err := hdwallet.NewFromSeed(seed)
	check(err)

	path := hdwallet.MustParseDerivationPath(hdPath)
	account, err := wallet.Derive(path, false)
	check(err)

	return account.Address.Hex()
}
