package value

import (
	"log"
	"regexp"

	"github.com/ipe-dev/menu_project/errors"
	"golang.org/x/crypto/bcrypt"
)

type Password string

func NewPassword(s string) (Password, error) {
	if !regexp.MustCompile(`/^[a-zA-Z0-9]*$/`).Match([]byte(s)) {
		return Password(s), errors.NewCustomError("パスワードが不正です", s)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return Password(hash), nil
}
