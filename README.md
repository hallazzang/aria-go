<h1 align="center">aria-go</h1>

<p align="center">
	<a href="https://travis-ci.org/hallazzang/aria-go">
		<img alt="Build Status" src="https://travis-ci.org/hallazzang/aria-go.svg?branch=master">
	</a>
	<a href="https://codecov.io/gh/hallazzang/aria-go">
		<img alt="Codecov Status" src="https://codecov.io/gh/hallazzang/aria-go/branch/master/graph/badge.svg">
	</a>
	<a href="https://godoc.org/github.com/hallazzang/aria-go">
		<img alt="Godoc Reference" src="https://godoc.org/github.com/hallazzang/aria-go?status.svg">
	</a>
	<a href="https://goreportcard.com/report/github.com/hallazzang/aria-go">
		<img alt="Goreportcard Badge" src="https://goreportcard.com/badge/github.com/hallazzang/aria-go">
	</a>
</p>

Go implementation of the [ARIA] encryption algorithm.

[aria]: https://tools.ietf.org/html/rfc5794

## Installation

```
go get -u github.com/hallazzang/aria-go
```

## Usage

This package is compatible with the standard `crypto` package.

```go
package main

import (
	"fmt"

	"github.com/hallazzang/aria-go"
)

func main() {
	key := []byte("0123456789abcdef")
	plaintext := []byte("fedcba9876543210")

	block, err := aria.NewCipher(key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("plaintext: %s\n", plaintext)

	ciphertext := make([]byte, 16)
	block.Encrypt(ciphertext, plaintext)
	fmt.Printf("ciphertext: % x\n", ciphertext)

	decrypted := make([]byte, 16)
	block.Decrypt(decrypted, ciphertext)
	fmt.Printf("decrypted: %s\n", decrypted)
}
```

This will print out:

```
plaintext: fedcba9876543210
ciphertext: 93 52 1e f2 67 65 0b d1 fc 75 20 5a d3 44 4d 9d
decrypted: fedcba9876543210
```

## Benchmarks

Here's the benchmark result compared with `aes` package:

```
goos: windows
goarch: amd64
BenchmarkAES128  	100000000	        17.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAES192  	100000000	        21.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAES256  	100000000	        23.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkARIA128 	 1000000	      1225 ns/op	       0 B/op	       0 allocs/op
BenchmarkARIA192 	 1000000	      1468 ns/op	       0 B/op	       0 allocs/op
BenchmarkARIA256 	 1000000	      1610 ns/op	       0 B/op	       0 allocs/op
```

## License

MIT
