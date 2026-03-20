package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

)

const(
	length = 28
	letters ="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()"
)

func main() {


// Generate a password of length 28 characters
// 1. use a cryptographic randomness to generate a random password with the char set. This randomness is based on real OS entropy, look into csprng
// Each character is chosen from a set of 72 possible values (~6.17 bits of entropy per char) log base 2 (72).
// A raw byte has 256 possibilities (8 bits) 2 power 8, so using a charset slightly reduces entropy
// but makes the password human-readable while still remaining highly secure.
// 2. print the password generated

lol ,_:= generatePassword()

fmt.Println("The password generated:",lol)

}

func generatePassword() (string,error){

	b := make([]byte, length)
	for i := range b{
		index ,err := rand.Int(rand.Reader,big.NewInt(int64(len(letters))))
		if err!=nil{
			return "",err
		}
		b[i] = letters[index.Int64()]
	}
    return string(b), nil



}
