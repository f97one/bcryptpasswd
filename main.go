package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	flag.Parse()
	hash, err := bcrypt.GenerateFromPassword([]byte(flag.Arg(0)), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", hash)
}
