# golang-eBPF

eBPF with golang, based on https://ebpf-go.dev/guides/getting-started/#the-go-application

## Setup (ubuntu 22.04)

```
sudo apt install -y llvm-11 llvm-11-dev llvm-11-tools clang-11 llvm 
sudo apt install linux-headers-`uname -r`
sudo ln -s /usr/include/x86_64-linux-gnu/asm /usr/include/asm
```

## Run

```
go generate && go build && sudo ./ebpf-test
```