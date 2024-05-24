# golang-eBPF

eBPF with golang, based on https://ebpf-go.dev/guides/getting-started/#the-go-application

## Setup (ubuntu 22.04)

```
sudo apt install -y llvm-11 llvm-11-dev llvm-11-tools clang-11 llvm linux-tools-6.5.0-1023-oem linux-tools-common 
sudo apt install linux-headers-`uname -r`
sudo ln -s /usr/include/x86_64-linux-gnu/asm /usr/include/asm
```

## Run

### packetcounter

```
cd packetcounter
```

Packet counter counts network packets on an interface.

```
go generate && go build && sudo ./ebpf-test
```

### kprobe

```
cd kprobe
```

Kprobe attaches to the `sys_execve` symbol and reports the executed programs.

Run once:

```
bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
```


```
go generate && go build && sudo ./ebpf-test
```