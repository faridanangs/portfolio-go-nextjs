package utils

import (
	"portfolio/helpers"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(pass string) string {
	res, err := bcrypt.GenerateFromPassword([]byte(pass), 13)
	helpers.PanicIfError(err, "Error in generate password at utils")
	return string(res)
}
