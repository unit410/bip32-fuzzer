mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(dir $(mkfile_path))

build-rust:
	cd ${mkfile_dir}/derive/rust && cargo clean && cargo build
	ln -fs ${mkfile_dir}/derive/rust/target/debug/rust-key-derivation ${mkfile_dir}/bin/rust

build-golang:
	cd ${mkfile_dir}/derive/golang && go clean && go build main.go
	ln -fs ${mkfile_dir}/derive/golang/main ${mkfile_dir}/bin/golang

build: build-rust build-golang

compare:
	cd compare && go run main.go ../bin

.PHONY: compare
