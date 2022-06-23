# Rethinking Concurrency Patterns

[Supporting YT Video](https://www.youtube.com/watch?v=5zXAHh5tJqQ)

## Principles

1. [Start goroutines when you have concurrent work.](#first-point)
1. [Share by communicating.](#second-point)

## Links

- [Asynchronous APIs: Futures & Queues](#asynchronous-apis)
- [Condition Variables](#condition-variables)
- [Worker Pools](#worker-pools)

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
// **before** retrieving results.
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
   storage location chosen by the implementation is irrelevant** to the
   semantics of the language.

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
    // Wait atomically unlocks the mutex and suspends the goroutine.
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

## Second Point

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

// Acquire loops until a resource is available, then extracts it from the
// shared state.
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
  case p.sem <- token{}:
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

### Indicate existence of new data.

```go
type Queue struct {
  items chan []Item // non-empty slices only
  // empty conveys metadata about the items channel
  // it indicates that no goroutine is sending to items.
  empty chan struct{} // holds true if the queue is empty
}

func NewQueue() *Queue {
  items := make(chan []Item, 1)
  empty := make(chan struct{}, 1)
  empty <- struct{}{}
  return &Queue{items, empty}
}

// Get grabs the item it needs and will either return the items back to the
// queue or communicate that items is empty.
func (q *Queue) Get(ctx context.Context) (Item, error) {
  var items []Item
  select {
  case <-ctx.Done():
    return 0, ctx.Err()
  case items = <-q.items:
  }

  item := items[0]
  if len(items) == 1 {
    q.empty <- struct{}{}
  } else {
    q.items <- items[1:]
  }
  return item, nil
}

// Put puts the item in the queue and updates all of the items back to whatever
// the last call to Get had gotten from it.
func (q *Queue) Put(item Item) {
  var items []Item
  select {
  case items = <-q.items:
  case <-q.empty:
  }
  items = append(items, item)
  q.items <- item
}

// Each waiter consumes the data they need and communicates any remaning data
// back to the channels.
```

```go
// to wait for specific data, the waiters communicate **their needs.**
type waiter struct {
  n int
  c chan []Item
}

type state struct {
  items []Item
  wait  []waiter
}

type Queue struct {
  s chan state
}

func NewQueue() *Queue {
  s := make(chan state, 1)
  s <- state{}
  return &Queue{s}
}

func (q *Queue) GetMany(n int) []Item {
  s := <-q.s
  if len(s.wait) == 0 && len(s.items) >= n {
    items := s.items[:n:n]
    s.items = s.items[n:]
    q.s <- s
    return items
  }
  c := make(chan []Item)
  s.wait = append(s.wait, waiter{n, c})
  q.s <- s
  return <-c
}

// Put sends to the next waiter if -- and only if -- it has enough items for
// that waiter
func (q *Queue) Put(item Item) {
  s := <-q.s
  s.items = append(s.items, item)
  for len(s.wait) > 0 {
    w := s.wait[0]
    if len(s.items) < w.n {
      break
    }
    w.c <- s.items[:w.n:w.n]
    s.items = s.items[w.n:]
    s.wait = s.wait[1:]
  }
  q.s <- s
}
```

### Share completion by completing communication

```go
type Idler struct {
  mu sync.Mutex
  idle sync.Cond
  busy bool
  idles int64
}

func NewIdler() *Idler {
  i := new(Idler)
  i.idle.L = &i.mu
  return i
}

func (i *Idler) AwaitIdle() {
  i.mu.Lock()
  defer i.mu.Unlock()
  idles := i.idles
  for i.busy && idles == i.idles {
    i.idle.Wait()
  }
}

func (i *Idler) SetBusy(b bool) {
  i.mu.Lock()
  defer i.mu.Unlock()
  wasBusy := i.busy
  i.busy = b
  if wasBusy && !i.busy {
    i.idles++
    i.idle.Broadcast()
  }
}
```

```go
type Idler struct {
  next chan chan struct{}
}

func NewIdler() *Idler {
  next := make(chan chan struct{}, 1)
  next <- nil
  return &Idler{next}
}

func (i *Idler) AwaitIdle(ctx context.Context) error {
  idle := <-i.next
  i.next <- idle
  if idle != nil {
    select {
    case <-ctx.Done():
      return ctx.Err()
    case <-idle:
      // noop
    }
  }
  return nil
}

// SetBusy allocates a new channel on the idle to busy transition and closes
// the previous channel, if any, on the busy to idle transition.
func (i *Idler) SetBusy(b bool) {
  idle := <-i.next
  if b && (idle == nil) {
    idle = make(chan struct{})
  } else if !b && (idle != nil) {
    close(idle) // Idle now.
  }
  i.next <- idle
}
```

## Worker Pools

Async patterns -> goroutines

Condition Variables -> sharing resources

Worker pool (Called Thread Pool in other languages): 
Treat a set of goroutines as resources

### Naive Version

This will have problems with leaking the workers forever without
`sync.WaitGroup` and even with it we can have idle workers all the way up until
end of work, if end of work can be reached instead of deadlocking.

1. Start workers

```go
work := make(chan Task)
var wg sync.WaitGroup
for n := limit; n > 0; n-- {
  // Often created by the same goroutine sends tasks to workers
  wg.Add(1)
  go func() {
    for task := range work {
      perform(task)
    }
    wg.Done()
  }()
}
```

2. Send work

```go
// Sender blocks until a worker is available to receive next task
for _, task := range hugeSlice {
  work <- task
}
close(work)
wg.Wait()
```

### Not so beneficial

Efficiency in other languages -- Distribute work across threads

Goroutines are multiplexed onto multiple OS threads by `GOMAXPROCS`, but what
they do allow is limiting work in flight.

### Better Version

> Start goroutines when you have concurrent work to do now.

Omit the worker pool and it's channel and only use the `sync.WaitGroup` if you
can have unlimited goroutines

```go
// WorkerPool is a fixed pool of goroutines that receives and performs tasks
// up to a given limit
func WorkerPool(limit int) {
  var wg sync.WaitGroup
  for _, task := range hugeSlice {
    wg.Add(1)
    go func(t Task) {
      perform(t)
      wg.Done()
    }(task)
  }
  wg.Wait()
}
```

Or a semaphore if you need to limit the amount of workers and to make sure
there will only ever be 1 idle worker.

```go
// WorkerPool is a fixed pool of goroutines that receives and performs tasks
// up to a given limit
func WorkerPool(limit int) {
  sem := make(chan token, limit)
  for _, task := range hugeSlice {
    sem <- token{}
    go func(t Task) {
      perform(t)
      <-sem
    }(task)
  }

  // Equivalent to wg.Wait()
  for n := limit; n > 0; n-- {
    sem <- token{}
  }
}
```
