package auth

import (
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("mieayam"), nil)
}

func GetTokenAuth() *jwtauth.JWTAuth {
	return tokenAuth
}

func CreateJWT(userId uint, username string) (string, error) {
	_, token, err := tokenAuth.Encode(map[string]interface{}{
		"user_id":  userId,
		"username": username,
	})

	return token, err
}

func HashedPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
