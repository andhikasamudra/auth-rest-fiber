package util

import "golang.org/x/crypto/bcrypt"

func Encode(data []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
