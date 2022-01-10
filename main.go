package main

import (
	"crypto/rand"
	"ecc/pkcs12"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {

	priv, err := pkcs12.ReadPrivateKeyFromSMFile("test_gm.sm2", "cfca1234")
	if err != nil {
		log.Fatal(err)
	}
	msg := "Hello,World!"
	sign, err := priv.Sign(rand.Reader, []byte(msg), nil)
	fmt.Println(hex.EncodeToString(sign))

	verify := priv.Verify([]byte(msg), sign)
	fmt.Println(verify)
}
