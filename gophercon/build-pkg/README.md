# Building Package in Go

- `GODEBUG=gocachehash=1 go install -v ./` -- see info on how pkg hashed
- `go env GOCACHE` -- find where your cache is

There are 256 directories in the build cache, for concurrent reads and writes.

## Hash formation

![Hash](./calculate_hash.png)

## Cache

![Cache](./go_cache.png)

## Building dependencies

![Build Order](./build_order.png)

## Recalculating builds

![Dependencies](./dependencies.png)

## Caching in CI (speeding up builds)

1. Load the cache from `go env GOCACHE`
1. Build binaries
1. Run the tests
1. Save the contents of `go env GOCACHE`
