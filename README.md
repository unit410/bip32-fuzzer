BIP-32 Comparison
=================

Compare BIP-32 key derivations and report any inconsistencies.

Note: This is experimental software with experimental quality.

### Usage

```shell
make build compare

35000 Cases, 137 Errors for 0.39% Error rate
- inconsistent:  true
  hd_path:       m/44'/60'/0'/0/0
  mnemonic:      foam skirt great wife version tower trouble lion guard window floor replace
  passphrase:
  bins:          [../bin/golang ../bin/rust]
  addresses:     [0xBb56980521a9418796E67e8115376Ebf92108040 0x9C91428b700FAf641eCDd45992626FCF82508420]
```

### Adding Derivers

BIP-32 implementations can be added to this repo under `./derive/<name>`. Each of these sub-projects should (1) implement the following call interface, (2) return a derived ethereum address to stdout and (3) be added to a `make build` target:

```shell
# Input
./derive <derivation path> <mnemonic phrase> (<passphrarse>)
./derive "m/44'/60'/0'/0/0" "word word word" "password"
```
