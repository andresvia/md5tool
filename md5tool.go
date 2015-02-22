package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	h := md5.New()
	b := make([]byte, md5.BlockSize * md5.BlockSize)
	for {
		n, _ := os.Stdin.Read(b)
		if n != 0 {
			h.Write(b[:n])
		} else {
			break
		}
	}
	fmt.Printf("%x", h.Sum(nil))
}
