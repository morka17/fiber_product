package security

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidToken = errors.New("Invalid jwt")
	JWTSecretKey    = []byte(os.Getenv("JWT_SECRET_KEY"))
)

func NewToken(userId string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    userId,
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecretKey)
}

func parseJWTCallback(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexcepted signing Method: %v\n", token.Header)
	}
	return JWTSecretKey, nil
}

func ExtractToken(r *http.Request) (string, error) {
	// Authorization => Bearer Token ....
	header := strings.TrimSpace(r.Header.Get("Authorization"))
	splitted := strings.Split(header, " ")
	if len(splitted) != 2 {
		log.Println("error on extract token from header:", header)
		return "", ErrInvalidToken
	}
	return splitted[1], nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, parseJWTCallback)
}

type TokenPayload struct {
	UserId    string
	CreatedAt time.Time
	ExpiresAt time.Time
}


func NewTokenPayload(tokenString string) (*TokenPayload, error ){
	token, err := ParseToken(tokenString)
	if err != nil {
		return nil, err 
	}
	claims , ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, ErrInvalidToken
	}

	id, _ := claims["iss"].(string)
	createdAt, _ := claims["iat"].(int64)
	expiresAt, _ := claims["exp"].(int64)

	return &TokenPayload{
		UserId:    id,
		CreatedAt: time.Unix(createdAt, 0),
		ExpiresAt: time.Unix(expiresAt, 0),
	}, nil

}