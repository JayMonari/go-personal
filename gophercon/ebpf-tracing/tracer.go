package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/iovisor/gobpf/bcc"
)

// const eBPFProgram = `
// #include <uapi/linux/ptrace.h>
//
// BPF_PERF_OUTPUT(events);
//
// int handler_was_called(struct pt_regs *ctx) {
//   char message[] = "handler function was called";
//   events.perf_submit(ctx, &message, sizeof(message));
//   return 0;
// }
// `

const eBPFProgram = `
#include <uapi/linux/ptrace.h>

BPF_PERF_OUTPUT(events);

int addTwoNumbersWasCalled(struct pt_regs *ctx) {
  void* stackAddress = (void*)ctx->sp;

  long argument1;
  bpf_probe_read(&argument1, sizeof(argument1), stackAddress+8);
  events.perf_submit(ctx, &argument1, sizeof(argument1));

  char argument2;
  bpf_probe_read(&argument2, sizeof(argument2), stackAddress+16);
  events.perf_submit(ctx, &argument2, sizeof(argument2));

  return 0;
}
`

func main() {
	bpfModule := bcc.NewModule(eBPFProgram, []string{})

	// uprobeFd, err := bpfModule.LoadUprobe("handler_was_called")
	uprobeFd, err := bpfModule.LoadUprobe("addTwoNumbersWasCalled")
	if err != nil {
		log.Fatal(err)
	}

	// err = bpfModule.AttachUprobe(os.Args[1], "main.handlerFunction", uprobeFd, -1)
	err = bpfModule.AttachUprobe(os.Args[1], "main.addTwoNumbers", uprobeFd, -1)
	if err != nil {
		log.Fatal(err)
	}

	table := bcc.NewTable(bpfModule.TableId("events"), bpfModule)
	perfCh := make(chan []byte)

	perfMap, err := bcc.InitPerfMap(table, perfCh, nil)
	if err != nil {
		log.Fatal(err)
	}

	perfMap.Start()
	defer perfMap.Stop()

	go func() {
		for {
			// val := <-perfCh
			// fmt.Println(string(val))
			firstBytes := <-perfCh
			first := binary.LittleEndian.Uint64(firstBytes)

			secondBytes := <-perfCh
			second := binary.LittleEndian.Uint16(secondBytes)

			fmt.Printf("Arguments: %d %d\n", first, second)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}
