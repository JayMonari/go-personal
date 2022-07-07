# Gophercon

## 2019

### [Gabbi Fisher - Where do Sockets live in Go](https://www.youtube.com/watch?v=pGR3r0UhoS8)

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

### [Natalie Pistunovich - Going Secure With Go](https://www.youtube.com/watch?v=9e2gRtzemGo)

- Use gosec
- Use golangci-lint (depgaurd)
- Dependabot

### [Junade Ali - Higher Reliability Software in Go](https://www.youtube.com/watch?v=gB2dxBDjHP4)

Recursion is illegal Barr Testimony (Bookout V. Toyota) according to MISRA-C

### [Jacob Walker - Understanding Allocations - The Stack And The Heap](https://www.youtube.com/watch?v=ZMZpH4yT7M0)

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

### [Ignat Korchagin - Go as a scripting language in Linux](https://www.youtube.com/watch?v=fcyHqDwGchI&list=PLMW8Xq7bXrG5B_gvikeSf3Du3NGBs4yVi&index=3)

`echo ':golang:E::go::/usr/local/bin/gorun:OC' | sudo tee /proc/sys/fs//binfmt_misc/register`

### [Bryan Boreham - Go Tune Your Memory](https://www.youtube.com/watch?v=uyifh6F_7WM&list=PLMW8Xq7bXrG5B_gvikeSf3Du3NGBs4yVi&index=6)

| Situation                     | Action     | Pro           | Con         |
|-------------------------------|------------|---------------|-------------|
| Large static data set         | GOGC ⬇️     | Smaller heap  | More CPU    |
| Tiny heap, rapid GC           | GOGC ⬆️     | Lower letency | More RAM    |
| One-shot execution (go build) | `GOGC=off` | Runs faster   | May Explode |

### [Joan López de la Franca Beltran - Dockerization of Go](https://www.youtube.com/watch?v=GnXmON9rLQw&list=PLMW8Xq7bXrG5B_gvikeSf3Du3NGBs4yVi&index=13)

Making an extremely small docker image with `distroless`

```Dockerfile
FROM golang:stretch AS builder
WORKDIR $GOPATH/src/github.com/freindsofgo/bacon-ipsum
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o
/go/bin/bacon-ipsum cmd/bacon-ipsum/main.go

FROM gcr.io/distroless/base
COPY --from-builder /go/bin/bacon-ipsum /go/bin/bacon-ipsum
ENTRYPOINT ["/go/bin/bacon-ipsum"]
```

Why distroless?

- No package manager
- Immutability
- Secure-friendly (certs, etc.)
- No shell access (Hacker getting in does nothing)

### [Dave Cheney - Constant Time](https://www.youtube.com/watch?v=pN_lm6QqHcw&list=PLMW8Xq7bXrG5B_gvikeSf3Du3NGBs4yVi&index=14)

`const uintSize = 32 << (^uint(0) >> 32 & 1)`

On 64 bit platform
```
^uint(0) == 1111111111111111111111111111111111111111111111111111111111111111
>> 32 ==    0000000000000000000000000000000011111111111111111111111111111111
&1 ==       0000000000000000000000000000000000000000000000000000000000000001
32 << 1 ==  0000000000000000000000000000000000000000000000000000000001000000
```

Constants are **fungible** aka **identical** aka **equal** aka **`Singleton`**

```go
type Error string
func (e Error) Error() string {
  return string(e)
}
const err = Error("EOF")
// therefore
const str1 = "EOF"
const str2 = "EOF"
fmt.Println(str1 == str2) // true
const err1 = Error("EOF")
const err2 = Error("EOF")
fmt.Println(err1 == err2) // true
```

### [Daniel Marti - Optimizing `Go` code without a blindfold](https://www.youtube.com/watch?v=jiXnzkAzy30&list=PLMW8Xq7bXrG5B_gvikeSf3Du3NGBs4yVi&index=15)

- Better `benchcmp` 👉 `golang.org/x/perf/cmd/banchstat`
- Use `go test -bench=XXX -count=8 > old.txt; benchstat old.txt`
- Need `perflock` for avoiding noise in benchmarks
- `perflock -daemon & perflock -governor=70& go test -...`
- `benchstat old.txt new.txt`
- Ask compiler for optimizing decisions `go build -gcflags='-m -m' io`
- `... | grep 'function too complex'` or `'escapes to heap'`
- bounds check eliminiation `... -gcflags=-d=ssa/check_bce/debug=1 io`
- prove stuff `... -gcflags=-d=ssa/prove/debug=1 io` or `=2`
- `GOSSAFUNC={pattern} go build`

## 2017

### [Pascal Costanza - Go, C++ or Java for DNA sequencing?](https://www.youtube.com/watch?v=8zfC4xLb6YQ&list=PLMW8Xq7bXrG7acNjsU5YMGl5MMK5gl2vn&index=9)

`Go` wins over C++???? 🤨 (Another reason to not to use C++)
