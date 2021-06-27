package value

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func NewPassword(s string) Password {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return Password(hash)
}
