# Marshallers

To solve a consensus problem, control needs to be offloaded to someone outside.
These challenges are:

- Control Hand-off
- Serialization
- Data Sharing
- Data Consistency

## `sync.WaitGroup`

### Used For

- Control Hand-off

### Not For

- Serialization
- Data Sharing
- Data Consistency

### Gotcha

```go
func passByValue(wg sync.WaitGroup) {
  defer wg.Done()
  log.Println("...")
}

func main() {
  var wg sync.WaitGroup
  for i := 0; i < 100; i++ {
    wg.Add(1)
    // XXX(jay): This will fail because we pass a copy here, not a reference
    go passByValue(wg)
  }
  wg.Wait()
}
```

## `sync.Mutex`

### Used For

- Serialization
- Data Consistency

### Not For

- Control Hand-off
- Data Sharing

### Gotcha

```go
type Counter struct {
  sync.Mutex
  n int
}

func (c *Counter) Add(d int) {
  c.Lock()
  defer c.Unlock()
  c.n += d
}

// XXX(jay): This is a data race. There is a copy of the lock here not the
// reference. Needs to be a `c *Counter` pointer method
func (c Counter) Value() int {
  c.Lock()
  defer c.Unlock()
  return c.n
}
```

## `sync/atomic`

### Used For

- Data Consistency

### Not For

- Control Hand-off
- Serialization
- Data Sharing

### Gotcha

```go
func CombineMarshallers() {
  type Map map[string]string
  var m atomic.Value
  var mu sync.Mutex
  m.Store(make(Map))

  read := function(key string) (val string){
    m1 := m.Load().(Map)
    return m1[key]
  }
  insert := func(key, val string) {
    // XXX(jay): Without this mutex if multiple writers tried inserting a data
    // race would happen.
    mu.Lock()
    defer mu.Unlock()
    m1 := m.Load().(Map)
    m2 := make(Map)
    for k, v := range m1 {
      m2[k] = v
    }
    m2[key] = val
    m.Store(m2)
  }
}
```

```go
func works(x *int32) chan int32 {
  c := make(chan int32)
  go func() {
    atomic.AddInt32(x, 2)
    log.Println("1")
    time.Sleep(1 * time.Second)
    c <- *x
  }
  return c
}
func main() {
  var x, y int32
  // XXX(jay): This is not what we expected because we are using the `<-`
  // operator, which **blocks** until it receives a value.
  log.Println(<-works(&x), <-works(&x)) // executes first one then the other

  // actually concurrent
  ac, bc := works(&y), works(&y)
  log.Println(<-ac, <-bc) // executes both
}
```

## Channels

### Used For

- Control Hand-off
- Serialization
- Data Sharing

### Not For

- Data Consistency

### Gotcha

```go
func request() int {
  c := make(chan int)
  for i := 0; i < 5; i++ {
    i := i
    go func() {
      // XXX(jay): 4 out of 5 goroutines are blocked forever; memory leak.
      c <- i
    }()
  }
  return <-c
}
func request() int {
  c := make(chan int)
  for i := 0; i < 5; i++ {
    i := i
    go func() {
      select {
      case c <- i:
      default:
      }
    }()
  }
  // XXX(jay): If channel used outside of function and channel is drained this
  // return is blocked indefinitely
  return <-c
}

func wheresOne() {
  n := 10
  c := make(chan int)
  wg.Add(n)
  for i := 0; i < n; i++ {
    go func(val int) {
      defer wg.Done()
      c <- val
    }(i)
  }

  go func() {
    for val := range c {
      a = append(a, val)
    }
  }()

  wg.Wait()
  // XXX(jay): We don't wait until we drain all values from the `range` but
  // when all values have entered the channel.
  close(c)
}

func heresOne() {
  ...

  go func() {
    wg.Wait()
    close(c)
  }()

  for val := range c 
    a = append(a, val)
  }

  // OR

  done := make(chan struct{}, 1)
  go func() {
    for val := range c {
      a = append(a, val)
    }
    done <- struct{}{}
  }()

  wg.Wait()
  close(c)
  <-done
}

func longRunning(msgs <-chan string) {
  t := time.After(time.Minute)
  for {
    select {
    // XXX(jay): Memory leak; `<-` blocks until it receives. A new value is
    // created every time and the Timer is not GC'd until after Timer fires.
    // case <-time.After(time.Minute):
    case <-t.C:
      return
    case msg := <-msgs:
      fmt.Println(msg)
    }
    t.Reset(time.Minute)
  }
}
```
