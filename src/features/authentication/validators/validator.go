package validators

import (
	"errors"
	"strings"

	"github.com/morka17/fiber_product/src/features/authentication/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrEmptyName  = errors.New("Name field can't be empty")
	ErrEmptyEmail = errors.New("Email can't be empty")
	ErrEmptyPassword = errors.New("Password field can't be empty")
	ErrEmailAlreadyExists  = errors.New("Email already exits")
	ErrInvalidEmailAddress	=   errors.New("Invalid email address")
	ErrInvalidUserId 	=   errors.New("Invalid user id")
	ErrSigninFailed 	=  errors.New("signin failed")
)




func ValidateSignUp(user *models.User ) (error){
	if !primitive.IsValidObjectID(user.Id.Hex()){
		return ErrInvalidUserId
	}

	if user.Email == ""{
		return ErrEmptyEmail
	}

	if !strings.Contains(user.Email, "@"){
		return  ErrInvalidEmailAddress
	}

	if user.Name == ""{
		return ErrEmptyName 
	}

	if user.Password == ""{
		return ErrEmptyPassword
	}

	return nil 
}


func NormalizeEmail(email string ) string {
	return strings.TrimSpace(strings.ToLower(email))
}