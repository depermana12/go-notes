package auth

import (
	"net/http"

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

func GetIdFromAuthCtx(r *http.Request) (uint, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return 0, err
	}

	return uint(userId), nil
}

func HashedPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
