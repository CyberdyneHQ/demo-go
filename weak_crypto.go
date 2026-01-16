package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/deepsourcelabs/demo-go/multi"
	"os"
)

func makeMD5Hash() {
	for _, arg := range os.Args {
		fmt.Printf("%x - %s\n", md5.Sum([]byte(arg)), arg)
	}
}

func generateRSAKey() {
	//Generate Private Key
	pvk, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pvk)
	// Create a context with a value
	ctx := context.WithValue(context.Background(), "somekey", "somevalue")

	// Call the function with the context
	result := GetSomethingWithContext("example", ctx)

	// Print the result
	log.Println(result)
}
