package security

import (
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewToken(t *testing.T){

	id := primitive.NewObjectID()

	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

}



func TestNewTokenPayload(t *testing.T) {

	id := primitive.NewObjectID()

	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := NewTokenPayload(token)
	assert.NoError(t, err )
	assert.NotNil(t, payload)

	tokenExpired := getTokenExpired(id.Hex())
	payload, err = NewTokenPayload(tokenExpired)
	assert.Error(t, err)
	assert.Nil(t, payload) 
}


func getTokenExpired(id string) string{
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5* -1).Unix(),
		Issuer:    id,
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JWTSecretKey)
	return tokenString
} 