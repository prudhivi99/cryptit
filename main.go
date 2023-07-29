package main

import "github.com/prudhivi99/cryptit/encrypt"
import "github.com/prudhivi99/cryptit/decrypt"
import "fmt"

func main() {
	encryptedStr := encrypt.Nimbus("Hello")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
