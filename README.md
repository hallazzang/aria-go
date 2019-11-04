# aria-go

[![godoc-badge]](https://godoc.org/github.com/hallazzang/aria-go)
[![goreportcard-badge]](https://goreportcard.com/report/github.com/hallazzang/aria-go)

[godoc-badge]: https://godoc.org/github.com/hallazzang/aria-go?status.svg
[goreportcard-badge]: https://goreportcard.com/badge/github.com/hallazzang/aria-go

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
	input := []byte("fedcba9876543210")

	block, err := aria.NewCipher(key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("plain: %s\n", input)

	cipher := make([]byte, 16)
	block.Encrypt(cipher, input)
	fmt.Printf("cipher: % x\n", cipher)

	plain := make([]byte, 16)
	block.Decrypt(plain, cipher)
	fmt.Printf("decrypted: %s\n", plain)
}
```

This will print out:

```
plain: fedcba9876543210
cipher: 93 52 1e f2 67 65 0b d1 fc 75 20 5a d3 44 4d 9d
decrypted: fedcba9876543210
```
