package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	flag.Parse()

	pw := flag.Arg(0)
	if len(pw) > 72 {
		fmt.Println("WARNING : The string passed as an argument is longer than 72 characters.")
		fmt.Println("          Strings after the 73rd character are ignored.")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(flag.Arg(0)), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", hash)
}
