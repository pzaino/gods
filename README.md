# GODS (Go Data Structures)

This repository contains my implementations of various data structures in Go. The data structures are implemented in a way that they can be used in a concurrent environment.

All data structures comes with a set of tests to ensure that they work as expected.

Each data structure is implemented as a separate package in the `pkg` directory. The `cmd` directory contains a set of example programs that demonstrate how to use the data structures.

This should make it to use the data structures in your own projects and without making your code get too big. You can just import the package you need and start using it.

All data structures were designed to use generics, so some method call may require you to provide a comparison function or a hash function.

## Data Structures

- [x] [Stack](stack.go)
- [ ] [Queue](queue.go)
- [ ] [Priority Queue](priority_queue.go)
- [x] [Linked List](linked_list.go)
- [x] [Doubly Linked List](doubly_linked_list.go)
- [ ] [Circular Linked List](circular_linked_list.go)
- [ ] [Binary Search Tree](binary_search_tree.go)
- [ ] [AVL Tree](avl_tree.go)
- [ ] [Trie](trie.go)
- [ ] [Graph](graph.go)
- [ ] [Disjoint Set](disjoint_set.go)
- [ ] [Segment Tree](segment_tree.go)
- [ ] [Fenwick Tree](fenwick_tree.go)
