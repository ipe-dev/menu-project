package value

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func NewPassword(s string) Password {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), 12)
	if err != nil {
		log.Println(err)
	}
	return Password(hash)
}
