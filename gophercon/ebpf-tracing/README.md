# eBPF Tracing

## Getting started

```go
go build -o demo ./main.go
go build -o tracer ./tracer.go
./demo & sudo ./tracer ./demo
```

## Why this is useful

- Logging and debugging of programs that are already compiled, deployed, or
running.
- Sometimes you don't control the source code.

## Tracing arguments

1. Access the top of the stack using the `sp` fielf of `pt_regs`
1. Calculate the arguments offset off the stack
1. Read those arguments off the stack using `bpf_probe_read`
1. Send them up to userpace via the perf buffer
1. Parse them in Go using the `encoding/binary` package

## Calculate stack offsets

- Arguments start 8 bytes off the stack (first 8 is for base pointer)
- Parameters are padded on the stack based on the largest data type amongst
them (e.g. int, int8 -- int is largest at 8 bytes)
- "Window" size based on largest data type.
- Parameters are put on the stack based on if they fit in the current "window"
- If not the stack is padded until the start of the next "window"
- Parameter ordering (kinda) affects memory efficiency!

### arguments.go example

- `int` is 8 bytes `int8` is 1 byte
- The first parameter starts at `sp`+8 and takes up 8 bytes
- The second parameter starts at `sp`+15 and takes up 1 byte followed by 7
bytes of padding
