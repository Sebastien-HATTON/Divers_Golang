package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashed)
}

// HashPassword prend un mot de passe en chaîne de caractères et retourne un hash sous forme de []byte.
func hashPasswordByte(password string) ([]byte) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil
	}
	return hash
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func main() {
	password1 := "password1"

	hashedPassword1 := hashPassword(password1)
	fmt.Println("Hash pour password1:", hashedPassword1)


	hashPasswordByte1 := hashPasswordByte(password1)
	fmt.Println("Hash []byte pour password1:", hashPasswordByte1)

}
