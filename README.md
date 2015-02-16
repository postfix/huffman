[![GoDoc](http://godoc.org/github.com/mozu0/huffman?status.png)](http://godoc.org/github.com/mozu0/huffman)
[![Build Status](https://travis-ci.org/mozu0/huffman.svg?branch=master)](https://travis-ci.org/mozu0/huffman)
# huffman
Computes Huffman encoding of elements, given frequencies.

```go
huffman.FromInts([]int{3, 1, 2}) //=> ["1", "00", "01"]
```

http://en.wikipedia.org/wiki/Huffman_coding
