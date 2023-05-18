package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	pwd, err := EncryptPassword("123456")
	assert.NoError(t, err)
	assert.NotEmpty(t, pwd)
	assert.Len(t, pwd, 60)
}


func TestVerifyPassword(t *testing.T){
	code := "1234567"
	pwd, err := EncryptPassword(code)
	assert.NoError(t, err)
	assert.NotEmpty(t, pwd)
	assert.Len(t, pwd, 60)

	assert.NoError(t, VerifyPassword(pwd, code))

	assert.Error(t, VerifyPassword(pwd, "123456"))
	assert.Error(t, VerifyPassword(pwd, pwd))
	assert.Error(t, VerifyPassword(code, pwd))

}