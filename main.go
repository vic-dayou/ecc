package main

import (
	"crypto/rand"
	"ecc/pkcs12"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	data, err := ioutil.ReadFile("test_gm.sm2")
	if err != nil {
		return
	}
	smData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		fmt.Println(err)
	}
	priv, _, err := pkcs12.Decode(smData, "cfca1234")
	if err != nil {
		log.Fatal(err)
	}
	msg := "Hello,World!"
	sign, err := priv.Sign(rand.Reader, []byte(msg), nil)
	fmt.Println(hex.EncodeToString(sign))

	verify := priv.Verify([]byte(msg), sign)
	fmt.Println(verify)
}
