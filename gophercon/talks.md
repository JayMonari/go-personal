# Gophercon

## GopherCon 2019

### Where do Sockets live in Go

- What are Sockets?
  File Descriptors are used to access I/O resources on computers, encountered
  when writing/reading from files.
  Sockets are writing/reading to networks. They are a subset of FDs and there are
  two types:

1. Stream Sockets (TCP)
1. Datagram Sockets (UDP)

#### Stream Sockets

Server

- socket() - Creates socket that clients can send request to.
- bind() - Bind socket to an addr (IP:Port) pair for clients @ 1.2.3.4:80
- listen() - Accept incoming connections from client sockets
- accept() - blocks until client connects
- read() - output of data from the client that can be written
- write() - input data into stream for client to read
- close() - free FD (socket)

Client

- socket() - Creates socket that can be used for TCP streams
- connect() - connects to an address @ 1.2.3.4:80
- read() - output of data from the server that can be written
- write() - input data into stream for server to read
- close() - free FD (socket)

#### Datagram Sockets

Server

- socket()
- bind()
- sendto()
- recvfrom()
- close()

Client

- socket()
- sendto()
- recvfrom()
- close()

#### Practical Socket Applications

##### DNS Resolvers

DNS Resolvers have both UDP and TCP running on the same port, how so?

##### 5-Tuples for Forwarding to Sockets

|-------------------|  |-------------------|
|Protocol: **UDP**  |  |Protocol: **TCP**  |
|DstIP: **1.1.1.1** |  |DstIP: **1.1.1.1** |
|DstPort: **53**    |  |DstPort: **53**    |
|SrcIP: **2.3.4.5** |  |SrcIP: **2.3.4.5** |
|SrcPort: 4387      |  |SrcPort: 4387      |
|-------------------|  |-------------------|

#### Socket Options

- load balancing between worker processes reading from a shared queue
- running parallel ingress queues
- implementing packet filtering via BPF

## Singapore 2019

### Going Secure With Go

- Use gosec
- Use golangci-lint (depgaurd)
- Dependabot

### Optimizing Go code without a blindfold

- `go test --bench=. --count=8`
- `go tool pprof cpu.out`
- `benchstat old.txt`
- `perflock -daemon & perflock -governor=70% go test -...`
- Higher variance, higher `--count`
- `go build -gcflags='-m -m' io | grep -E '(function too complex)|(escapes to heap)'`
- `go build -gcflags=-d=ssa/check_bce/debug=1 io`
- `go build -gcflags=-d=ssa/prove/debug=1 io`
- count runes `len([]rune(str))`
- `GOSSAFUNC=HelloWorld go build && open ssa.html`
- `cmd/compile/{README,internal/ssa/README}`

### Higher Reliability Software in Go

- Recursion is illegal Barr Testimony (Bookout V. Toyota) according to MISRA-C

### Understanding Allocations - The Stack And The Heap

```go
// ESCAPES TO HEAP       ||  // STAY ON STACK
func main() {            ||  func main() {
  b := read()            ||    b := make([]byte, 32)
                         ||    read(b)
  // use b               ||    // use b
}                        ||  }
func read() []byte {     ||  func read(b []byte) {
  // return a new slice. ||    // write into slice.
  b := make([]byte, 32)  ||  }
  return b               ||
}                        ||
```

- Each Goroutine has its own stack
- Need to know only if:
  1. Program **NOT** fast enough
  1. and you have benchmarks to prove it
  1. and they show excessive heap allocations.
- Sharing down typically stays on the stack. (Passing values and pointers)
- Sharing up typically escapes to the Heap (returning pointers)
- Escape Analysis -- ask compiler
- `go tool compile -h` for gcflags options
- `go build -gcflags "-m=2"`
- Common allocated values
  - Pointers
  - Variables in Interface variables
  - Func literal variables
  - Variables capture by a closure
  - Backing data for maps, chans, slices, strings

## Gophercon Europe 2020

# dotGo

## 2019

### Ignat Korchagin - Go as a scripting language in Linux

`echo ':golang:E::go::/usr/local/bin/gorun:OC' | sudo tee /proc/sys/fs//binfmt_misc/register`

### Bryan Boreham - Go Tune Your Memory

| Situation                     | Action     | Pro           | Con         |
|-------------------------------|------------|---------------|-------------|
| Large static data set         | GOGC ⬇️     | Smaller heap  | More CPU    |
| Tiny heap, rapid GC           | GOGC ⬆️     | Lower letency | More RAM    |
| One-shot execution (go build) | `GOGC=off` | Runs faster   | May Explode |
