package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Usage: compare <directory of derivation binaries>")
		os.Exit(1)
	}

	// Get derivation binaries
	binDir := args[0]
	log.Println("Using bin directory: ", binDir)
	files, err := ioutil.ReadDir(binDir)
	check(err)

	derivationBinaries := []string{}
	for _, f := range files {
		binPath := path.Join(binDir, f.Name())
		derivationBinaries = append(derivationBinaries, binPath)
	}

	// Compare derivations forever
	cases := 0
	errors := 0
	for {
		mnemonic := getNewMnemonic()
		passphrase := ""
		hdPath := "m/44'/60'"

		inconsistent, addresses := compareDerivations(derivationBinaries, hdPath, mnemonic, passphrase)

		if inconsistent {
			fmt.Println("- inconsistent: ", true)
			fmt.Println("  hd_path:      ", hdPath)
			fmt.Println("  mnemonic:     ", mnemonic)
			fmt.Println("  passphrase:   ", passphrase)
			fmt.Println("  bins:         ", derivationBinaries)
			fmt.Println("  addresses:    ", addresses)
			errors++
		}

		cases++
		if cases%100 == 0 {
			fmt.Printf("%v Cases, %v Errors for %.2f%% Error rate\n", cases, errors, 100.*float32(errors)/float32(cases))
		}
	}
}

func compareDerivations(derivationBinaries []string, hdPath string, mnemonic string, passphrase string) (bool, []string) {
	addresses := []string{}
	inconsistent := false

	for _, bin := range derivationBinaries {
		address := deriveAddress(bin, hdPath, mnemonic, passphrase)
		for _, r := range addresses {
			if r != address {
				inconsistent = true
			}
		}
		addresses = append(addresses, address)
	}
	return inconsistent, addresses
}

func deriveAddress(binPath string, hdPath string, mnemonic string, passphrase string) string {
	cmd := exec.Command(binPath, hdPath, mnemonic, passphrase)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return strings.TrimSpace(string(stdout))
}

func getNewMnemonic() string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	words := strings.Fields(mnemonic)
	formatted := strings.Join(words, " ")
	return formatted
}
