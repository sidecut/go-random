package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
)

func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%v\n", math_rand.Int63())
	}
}
