package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
)

var ErrInvalidToken = errors.New("invalid token")

// ErrForbidden is returned when a auth issue is identified.
var ErrForbidden = errors.New("attempted action is not allowed")

type Claims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Iat   string `json:"iat"`
	jwt.StandardClaims
}

func Verify(cookie *http.Cookie) (Claims, error) {
	tokenStr := cookie.Value

	claims := Claims{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, keyFunc)
	if err != nil {
		return Claims{}, err
	}

	if !token.Valid {
		return Claims{}, ErrInvalidToken
	}

	return claims, nil
}

// ctxKey represents the type of value for the context key.
type ctxKey int

// key is used to store/retrieve a Claims value from a context.Context.
const key ctxKey = 1

// SetClaims stores the claims in the context.
func SetClaims(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, key, claims)
}

// GetClaims returns the claims from the context.
func GetClaims(ctx context.Context) (Claims, error) {
	v, ok := ctx.Value(key).(Claims)
	if !ok {
		return Claims{}, errors.New("claim value missing from context")
	}
	return v, nil
}
