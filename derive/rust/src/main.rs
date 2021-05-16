extern crate bip39;
use bip39::{Language, Mnemonic, Seed};

extern crate eth_checksum;
extern crate hdwallet;
use hdwallet::{ExtendedPrivKey, ExtendedPubKey, KeyChain};

extern crate crypto;
use crypto::digest::Digest;
use crypto::sha3::Sha3;

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    let hd_path = args[1].to_string();
    let words = args[2].to_string();
    let mut passphrase = "";
    if args.len() > 3 {
        passphrase = &args[3];
    }
    println!("{}", derive_address(hd_path, words, passphrase));
}

fn derive_address(
    hd_path: std::string::String,
    words: std::string::String,
    passphrase: &str,
) -> std::string::String {
    // Derive the seed
    let mnemonic = Mnemonic::from_phrase(&words, Language::English).expect("Load mnemonic");
    let seed = Seed::new(&mnemonic, passphrase);

    // Derive the master keychain
    let private_key = ExtendedPrivKey::with_seed(seed.as_bytes()).unwrap();
    let key_chain = hdwallet::DefaultKeyChain::new(private_key);

    // Derive the child key
    let (child_key, _d) = key_chain.derive_private_key(hd_path.into()).unwrap();
    let pubchild_key = ExtendedPubKey::from_private_key(&child_key)
        .public_key
        .serialize_uncompressed();

    // Format the address
    let mut hasher = Sha3::keccak256();
    hasher.input(&pubchild_key[1..]);
    let address = format!("0x{}", &hasher.result_str()[24..]);
    eth_checksum::checksum(&address)
}
