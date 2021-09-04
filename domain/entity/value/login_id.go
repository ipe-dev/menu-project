package value

import (
	"regexp"

	"github.com/ipe-dev/menu_project/errors"
)

type LoginID string

func NewLoginID(loginID string) (LoginID, error) {
	if !regexp.MustCompile(`/^[a-zA-Z0-9]*$/`).Match([]byte(loginID)) {
		return LoginID(loginID), errors.NewCustomError("ログインIDが不正です", loginID)
	}
	return LoginID(loginID), nil
}
