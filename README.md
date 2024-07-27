# GoDS (Go Data Structures)

![Go build: ](https://github.com/pzaino/gods/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/pzaino/gods)](https://goreportcard.com/report/github.com/pzaino/gods)
[![Go-VulnCheck](https://github.com/pzaino/gods/actions/workflows/go-vulncheck.yml/badge.svg)](https://github.com/pzaino/gods/actions/workflows/go-vulncheck.yml)
![Scorecard supply-chain security: ](https://github.com/pzaino/gods/actions/workflows/scorecard.yml/badge.svg)

**WIP**: This project is still a work in progress. I will be adding more data structures as I implement them.

This repository contains my implementations of various data structures in Go. The data structures are implemented in two ways:

- Traditional single thread optimized (you can find these in the `pkg`
directory with their "standard" name, for ex. stack, linkList).
- Concurrency-safe wrapper added (you can find these in the `pkg`
directory with their "cs" initials, for example cstsack, cslinkList).

Use the non concurrent-safe data structures if you need best performance
in a non-concurrent application. Use the concurrent-safe data structures
if you need to use them in a concurrent application.

All data structures comes with a set of tests to ensure that they work as
 expected.

Each data structure is implemented as a separate package in the `pkg`
directory. The `cmd` directory contains a set of example programs that
 demonstrate how to use the data structures.

This should make it easy to use the data structures in your own projects and
 without making your code get too big. You can just import the package you
  need and start using it.

All data structures were designed to use generics, so some method call may
 require you to provide a comparison function or a hash function.

## Data Structures

- [x] [Stack](./pkg/stack)
- [x] [Concurrent Stack](./pkg/csstack)
- [x] [Queue](./pkg/queue)
- [ ] [Concurrent Queue](./pkg/csqueue)
- [x] [Priority Queue](./pkg/pqueue)
- [ ] [Concurrent Priority Queue](./pkg/cspqueue)
- [x] [Linked List](./pkg/linkList)
- [x] [Concurrent Linked List](./pkg/cslinkList)
- [x] [Doubly Linked List](./pkg/dlinkList)
- [x] [Concurrent Doubly Linked List](./pkg/csdlinkList)
- [ ] [Circular Linked List](./pkg/circularLinkList)
- [ ] [Concurrent Circular Linked List](./pkg/cscircularLinkList)
- [ ] [Binary Search Tree](./pkg/binarySearchTree)
- [ ] [AVL Tree](./pkg/avlTree)
- [ ] [Trie](./pkg/trie)
- [ ] [Graph](./pkg/graph)
- [ ] [Disjoint Set](./pkg/disjointSet)
- [ ] [Segment Tree](./pkg/segmentTree)
- [ ] [Fenwick Tree](./pkg/fenwickTree)
