# GoDS (Go Data Structures)

<img align="right" width="320" height="280"
 src="https://raw.githubusercontent.com/pzaino/gods/main/images/logo.png" alt="GoDS Logo">

![Go build: ](https://github.com/pzaino/gods/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/pzaino/gods)](https://goreportcard.com/report/github.com/pzaino/gods)
[![Go-VulnCheck](https://github.com/pzaino/gods/actions/workflows/go-vulncheck.yml/badge.svg)](https://github.com/pzaino/gods/actions/workflows/go-vulncheck.yml)
![Scorecard supply-chain security](https://github.com/pzaino/gods/actions/workflows/scorecard.yml/badge.svg)
![CodeQL](https://github.com/pzaino/gods/actions/workflows/codeql.yml/badge.svg)
[![FOSSA license](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpzaino%2Fgods.svg?type=shield&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpzaino%2Fgods?ref=badge_shield&issueType=license)
[![FOSSA Security](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpzaino%2Fgods.svg?type=shield&issueType=security)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpzaino%2Fgods?ref=badge_shield&issueType=security)
[![OSV Security](https://github.com/pzaino/gods/actions/workflows/osv-scanner.yml/badge.svg)](https://github.com/pzaino/gods/actions/workflows/osv-scanner.yml/badge.svg)

**WIP**: This project is still a work in progress. I will be adding more data
 structures as I implement them. Please do not use it yet, as it is not ready
  for production.

This repository contains my implementations of various data structures in Go.
 The data structures are implemented in two ways:

- Lockless: You can find these in the `pkg` directory with their "standard"
 names, e.g., stack, linkedList.
- Concurrency-safe wrapper: You can find these in the `pkg` directory with the
 prefix "cs", e.g., csstack, cslinkedList.

Use the non-concurrent-safe data structures if you need the best performance
 in a non-concurrent application or if you want to handle concurrency yourself.
  Use the concurrent-safe data structures for the best approach in concurrent
   applications or if you prefer not to handle concurrency yourself.

Both implementations generally come with three special methods:

- `ConfinedForEach`: This method iterates over the data structure and executes
 a function for each element. The function is executed in a separate goroutine
  for each element.
- `ConfinedForFrom`: This method iterates over the data structure and executes
 a function for each element starting from a given element. The function is
  executed in a separate goroutine for each element.
- `ConfinedForRange`: This method iterates over the data structure and
 executes a function for each element within a given range. The function is
 executed in a separate goroutine for each element.

These three methods offer better performance when there are multiple CPU cores
 available. They are also useful when you want to execute a function for each
  element in parallel. The number of goroutines created will be equal to the
   number of CPU cores available.

Every other method is not parallel. The Buffer has a special set of methods
 called:

- `Blit`: This method combines/overwrites the elements of the buffer with the
 elements of another buffer using a custom function. If the two buffers have
  different lengths, the function will use the size of the smaller buffer.
- `BlitFrom`: This method combines/overwrites the elements of the buffer with
 the elements of another buffer starting from a given index using a custom
  function. If the two buffers have different lengths, the function will use
   the size of the smaller buffer.
- `BlitRange`: This method combines/overwrites the elements of the buffer
 with the elements of another buffer within a given range using a custom
  function.

These functions are special because, for large amounts of data, they use
 parallelism to speed up the process. The number of goroutines created will
  be equal to the number of CPU cores available.

## Packaging and General Design

All data structures come with a set of tests to ensure that they work as
 expected.

Each data structure is implemented as a separate package in the `pkg`
 directory. The `cmd` directory contains a set of example programs that
  demonstrate how to use the data structures.

This design should make it easy to use the data structures in your own
 projects without significantly increasing your code size. You can simply
  import the package you need and start using it.

All data structures are designed to use generics, so some method calls may
 require you to provide a comparison function, hash function, etc.

## Installation / Usage

To use a library, you need to import it into your code. For example, to use
 the stack data structure, you would do:

```go
import "github.com/pzaino/gods/pkg/stack"
```

Then you can start using the stack data structure in your code.

You can use more than one data structure in your code. Just import the ones
 you need.

You may need to "install" the library first. To do that, use the following
 command:

```bash
go get github.com/pzaino/gods
```

This will download the library and install it in your `$GOPATH`.

## Data Structures

- [x] [Stack](./pkg/stack)
- [x] [Concurrent Stack](./pkg/csstack)
- [x] [Buffer](./pkg/buffer)
- [x] [Concurrent Buffer](./pkg/csbuffer)
- [ ] [Ring Buffer](./pkg/ringBuffer)
- [ ] [Concurrent Ring Buffer](./pkg/csringBuffer)
- [ ] [A/B Buffer](./pkg/abBuffer)
- [ ] [Concurrent A/B Buffer](./pkg/csabBuffer)
- [x] [Queue](./pkg/queue)
- [ ] [Concurrent Queue](./pkg/csqueue)
- [x] [Priority Queue](./pkg/pqueue)
- [ ] [Concurrent Priority Queue](./pkg/cspqueue)
- [x] [Linked List](./pkg/linkList)
- [x] [Concurrent Linked List](./pkg/cslinkList)
- [x] [Doubly Linked List](./pkg/dlinkList)
- [x] [Concurrent Doubly Linked List](./pkg/csdlinkList)
- [x] [Circular Linked List](./pkg/circularLinkList)
- [ ] [Concurrent Circular Linked List](./pkg/cscircularLinkList)
- [ ] [Binary Search Tree](./pkg/binarySearchTree)
- [ ] [AVL Tree](./pkg/avlTree)
- [ ] [Trie](./pkg/trie)
- [ ] [Graph](./pkg/graph)
- [ ] [Disjoint Set](./pkg/disjointSet)
- [ ] [Segment Tree](./pkg/segmentTree)
- [ ] [Fenwick Tree](./pkg/fenwickTree)

Legend:

- [x] Implemented
- [ ] Not implemented yet (WIP)

Given that this project is still a WIP, I may change the APIs as I move forward
 with the work. I will try to keep the changes to a minimum and make them as
  easy to adapt to as possible. However, I also want to achieve a good design.

## License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpzaino%2Fgods.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpzaino%2Fgods?ref=badge_large)
