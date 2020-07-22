package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := "password123"
	// hashing password
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(bs)

	// comparing hashed pw to entered pw
	err = bcrypt.CompareHashAndPassword(bs, []byte(pw))
	if err != nil {
		fmt.Println("You can't login!")
		return
	}

	fmt.Println("You're logged in")
}
