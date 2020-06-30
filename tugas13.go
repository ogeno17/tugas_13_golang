package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	var nama = "Hugo Ghally Imanaka"

	var encodeStr = base64.StdEncoding.EncodeToString([]byte(nama))
	fmt.Println("EncodeString : ", encodeStr)

	var sha = sha1.New()
	sha.Write([]byte(encodeStr))

	var enkripsi = sha.Sum(nil)
	var enkripsiStr = fmt.Sprintf("%x", enkripsi)

	fmt.Println("EnkripsiString : ", enkripsiStr)
}
