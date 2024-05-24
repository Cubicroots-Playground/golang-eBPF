package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/ringbuf"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {

		log.Fatal("Removing memlock:", err)
	}

	// Load the compiled eBPF ELF and load it into the kernel.
	var objs kprobeObjects

	if err := loadKprobeObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	// Attach count_packets to the network interface.
	link, err := link.Kprobe("sys_execve", objs.KprobeSysExecve, nil)
	if err != nil {
		log.Fatal("Attaching kprobe:", err)
	}
	defer link.Close()

	// Periodically fetch the packet counter from PktCount,
	// exit the program when interrupted.
	tick := time.Tick(time.Millisecond * 4)
	stop := make(chan os.Signal, 5)
	signal.Notify(stop, os.Interrupt)

	reader, err := ringbuf.NewReader(objs.Events)
	if err != nil {
		log.Fatal("new reader failed: ", err)
	}

	for {
		select {
		case <-tick:
			obj, err := reader.Read()
			if err != nil {
				log.Fatal("reader read:", err)
			}
			log.Printf("sys_execve: %s\n", string(obj.RawSample[4:]))
		case <-stop:
			log.Print("Received signal, exiting..")
			return
		}
	}
}
