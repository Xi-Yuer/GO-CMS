package bcrypt

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
}

func (b *Bcrypt) HashPassword(password string) (string, error) {
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(generateFromPassword), nil
}

func (b *Bcrypt) VerifyPassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
