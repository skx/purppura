package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage %s password1 password2 .. passwordN\n", os.Args[0])
		os.Exit(1)
	}
	for _, arg := range os.Args[1:] {
		pBytes, _ := bcrypt.GenerateFromPassword([]byte(arg), 14)
		pCrypt := string(pBytes)

		fmt.Printf("%s %s\n", arg, pCrypt)
	}
}
