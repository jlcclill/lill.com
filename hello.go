// hello.go
package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

func main() {
	var hs hash.Hash
	hs = sha256.New()
	fmt.Println("Hello World!")
	fmt.Println("%x", hs)
}
