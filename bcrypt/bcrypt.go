package bcrypt

import (
	"geulSsi/app/dto/custom"
	"golang.org/x/crypto/bcrypt"
)

func Generate(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Compare(hash, pwd string) *custom.Error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return custom.BcryptMismatch
		}
		return custom.BcryptParsingFail
	}
	return nil
}

func CompareIfNotMatch(hash, pwd string) (*string, *custom.Error) {
	if compare := Compare(hash, pwd); compare == nil {
		return nil, custom.SamePassword
	}

	encryptPassword, err := Generate(pwd)
	if err != nil {
		return nil, custom.BcryptGenerateFail
	}

	return &encryptPassword, nil
}
