# Rethinking Concurrency Patterns

## Principles

1. [Start goroutines when you have concurrent work.](First Point)
1. [Share by communicating.](Second Point)

## Links

- [Asynchronous APIs: Futures & Queues](Asynchronous APIs)
- [Condition Variables](Condition Variables)
- [Worker Pools](Worker Pools)
- [Recap](Recap)

## Asynchronous APIs

An asynchronous API returns to the caller **before** its result is ready.
Asynchronous does not mean concurrent as a callback can be given and a program
can sit around waiting for the results.

### How JS does it

```go
// Fetch immediately returns, then fetches the item and invokes f in a
// goroutine when the item is available.
// If the item does not exist, Fetch invokes f on the zero item.
func Fetch(name string, f func(Item)) {
  go func() {
    [...]
    f(item)
  }()
}
// Not how Go is written. Use channels and goroutines.
```

### How Java does it

Sometimes reffered to as `async/await`

```go
// Fetch immediately returns a channel, then fetches the requested item and
// sends it on the channel.
// If the item does not exist, Fetch closes the channel without sending.
func Fetch(name string) <-chan Item {
  c := make(chan Item, 1)
  go func() {
    [...]
    c <- item
  }()
  return c
}
// The Go analogue to a Future is a single-element buffered channel.

///////////////////////////////////////////////////////////////////////////////

// Yes:            |  // No: This blocks at a, then b
a := Fetch("a")    |  a := <-Fetch("a")
b := Fetch("b")    |  b := <-Fetch("b")
consume(<-a, <-b)  |  consume(a, b)
// Using Futures for concurrency, the caller must set up concurrent work
  **before** retrieving results.
```

### How RabbitMQ does it

```go
// Glob finds all items with names matching pattern and sends them on the
// returned channel.
// It closes the channel when all items have been sent.
func Glob(pattern string) <-chan Item {
  c := make(chan Item)
  go func() {
    defer close(c)
    for [...] {
      [...]
      c <- item
    }
  }()
  return c
}
// A channel fed by one goroutine and read by another acts as a queue.

///////////////////////////////////////////////////////////////////////////////

// Call site
for item := range Glob("[ab]*") {
  [...]
}
```

### Why use Asynchronous APIs?

#### Classical Benefits

1. Avoid blocking UI and network threads.
1. Reduce idle threads.
1. Reclaim stack frames.

#### Actual Benefit

- Initiate concurrent work.

### Why Go doesn't care

#### Not so good Classical Benefits

> **A goroutine [...] is lightweight**, costing little more than the allocation
> of stack space. And the stacks start small, so they are cheap, and grow by
> allocating (and freeing) heap storage as required.

1. Runtime manages threads for us so their is no single UI or network thread to
   block. No need to touch the kernel to switch goroutines.
1. Goroutines are ~2KB which is half the size of the **smallest** AMD64 page
1. Each variable in Go exists as long as there are references to it. **The
   storage location chosen by the implementation is irrelevant** to the semantics
   of the language.

#### Not so good Actual Benefit

- Asynchronous APIs introduce Caller-side ambiguity

```go
a := Fetch("a")
b := Fetch("b")
if err := [...] {
  return err
}
consume(<-a, <-b)
// What happens if we return early? Are resources still being used by the
// `Fetch` functions? Will we start fetches faster than they can be retired and
// run out of memory?

///////////////////////////////////////////////////////////////////////////////

a := Fetch(ctx, "a")
b := Fetch(ctx, "b")
[...]
consume(<-a, <-b)
// Will Fetch keep using the passed in context after it is returned?
// If so what happens if we cancel it and try to read from the channel --
// zero value, sentinel value, block?

///////////////////////////////////////////////////////////////////////////////

for result := range Glob("[ab]*") {
  if err := [...] {
    return err
  }
}
// If we return without draining the channel, is a goroutine leaked that is
// sending to Glob?

for result := range Glob(ctx, "[ab]*") {
  [...]
}
// Will Glob keep using the context as we iterate of the results?
// What happens if we cancel it?
// When, if ever, will the channel be closed?
```

In order to get the answers to all of these questions you have to hope and pray
🙏 that somewhere in the docs 📄 the answers can be found.

### History Lesson

> **A goroutine [...] is a function executing concurrently** with other
> goroutines in the same address space.

### First Point

> Start goroutines when **you** have concurrent work.

### Asynchronous == Synchronous

```go
func Async(x In) (<-chan Out) {
  c := make(chan Out, 1)
  go func() {
    c <- Synchronous(x)
  }()
  return c
}
// `c <- Synchronous(x)` blocks until a value can be passed
func Synchronous(x In) Out {
  c := Async(x)
  return <-c
}
// `return <-c` blocks until a value can be passed

///////////////////////////////////////////////////////////////////////////////

// Fetch returns the requested item.
func Fetch(context.Context, name string) (Item, error) {
  [...]
}
// There are no questions to answer as we are sure what we will produce.

///////////////////////////////////////////////////////////////////////////////

// We can introduce concurrency on the Caller-side with whichever pattern we so
// choose to implement

var a, b Item
g, ctx := errgroup.WithContext(ctx)
g.Go(func() (err error) {
    a, err = Fetch(ctx, "a")
    return err
})
g.Go(func() (err error) {
    b, err = Fetch(ctx, "a")
    return err
})
err := g.Wait()
[...]
consume(a, b)

// The caller can invoke synchronous functions concurrently, and often won't
// need to use channels at all.
```

### Make concurrency an **internal detail**

```go
// Glob finds all items with names matching pattern.
func Glob(ctx context.Context, pattern string) ([]Item, error) {
  [...] // find matching names.
  c := make(chan Item)
  g, ctx := errgroup.WithContext(ctx)
  for _, n := range names {
    n := n
    g.Go(func() error {
      item, err := Fetch(ctx, n)
      if err == nil {
        c <- item
      }
      return err
    })
  }
  go func() {
    err = g.Wait()
    close(c)
  }()

  var items []Item
  for i := range c {
    items = append(items, i)
  }
  if err != nil {
    return nil, err
  }
  return items, nil
}
```

> In Go no need to pay the costs of Asynchronicity for Concurrency

## Condition Variables

Also called Monitors are defined with an example.

```go
// A condition variable is associated with a sync.Locker (e.g. a Mutex).
type Queue struct {
  mu sync.Mutex
  items []Item
  itemAdded sync.Cond
}

func New() *Queue {
  q := new(Queue)
  q.itemAdded.L = &q.mu
  return q
}

func (q *Queue) Get() Item {
  q.mu.Lock()
  defer q.mu.Unlock()
  for len(q.items) == 0 {
    // Wait atomically unlcoks the mutex and suspends the goroutine.
    q.itemAdded.Wait()
  }
  item := q.items[0]
  q.items = q.items[1:]
  return item
}

func (q *Queue) Put(item Item) {
  q.mu.Lock()
  defer q.mu.Unlock()
  q.items = append(q.items, item)
  // Signal locks the mutex and wakes up the goroutine.
  q.itemAdded.Signal()
}
```

```go
type Queue struct {
  [...]
  closed bool
}

func (q *Queue) Close() {
  q.mu.Lock()
  defer q.mu.Unlock()
  q.closed = true
  // Broadcast usually communicates events that affect all waiters.
  q.cond.Broadcast()
}

func (q *Queue) GetMany(n int) []Item {
  q.mu.Lock()
  defer q.mu.Unlock()
  for len(q.items) < n {
    q.itemAdded.Wait()
  }
  items := q.items[:n:n]
  q.items = q.items[n:]
  return items
}

func (q *Queue) Put(item Item) {
  q.mu.Lock()
  defer q.mu.Unlock()
  q.items = append(q.items, item)
  // Since we don't know which of the GetMany calls is ready after a Put, we
  // can wake them all up and let them decide.
  q.itemAdded.Broadcast()
}
```

### The Bad

- Spurious wakeups
- Forgotten signals
- Starvation
- Unresponsive cancellation

Fundamentally condition variables rely on communicating by shared memory; they
signal that a change has occurred, but leave it up to the caller to check other
shared variables to see what changed.

Go has a different approach: **Share by communicating.**

### Communicating by sharing

```go
type Pool struct {
  mu              sync.Mutex
  cond            sync.Cond
  numConns, limit int
  idle            []net.Conn
}

func NewPool(limit int) *Pool {
  p := &Pool{limit: limit}
  p.cond.L
  return p
}

func (p *Pool) Release(c net.Conn) {
  p.mu.Lock()
  defer p.mu.Unlock()
  p.idle = append(p.idle, c)
  p.cond.Signal()
}

func (p *Pool) Hijack(c net.Conn) {
  p.mu.Lock()
  defer p.mu.Unlock()
  p.numConns--
  p.cond.Signal()
}

// Acquire loops until a resouce is available, then extracts it from the shared
// state.
func (p *Pool) Acquire() (net.Conn, error){
  p.mu.Lock()
  defer p.mu.Unlock()
  for len(p.idle) == 0 && p.numConns >= p.limit {
    p.cond.Wait()
  }
  if len(p.idle) > 0 {
    c := p.idle[len(p.idle)-1]
    p.idle = p.idle[:len(p.idle)-1]
    return c, nil
  }
  c, err := dial()
  if err == nil {
    p.numConns++
  }
  return c, err
}
```

### Sharing by communicating

```go
// Channel operations combine synchronization, signaling, and data transfer.
type Pool struct {
  sem chan token
  idle chan net.Conn
}
type token struct{}

func NewPool(limit int) *Pool {
  sem := make(chan token, limit)
  idle := make(chan net.Conn, limit)
  return &Pool{sem, idle}
}

func (p *Pool) Release(c net.Conn) {
  p.idle <- c
}

func (p *Pool) Hijack(c net.Conn) {
  <-p.sem
}

// When we block on communicating others can **also** communicate with us, e.g.
// to cancel the call.
func (p *Pool) Acquire(ctx context.Context) (net.Conn, error) {
  select {
  case conn := <-p.idle:
    return conn, nil
  case p.sem <- token():
    conn, err := dial()
    if err != nil {
      <-p.sem
    }
    return conn, err
  case <-ctx.Done():
    return nil, ctx.Err()
  }
}
```