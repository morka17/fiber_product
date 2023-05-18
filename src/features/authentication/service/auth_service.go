package authservice

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/morka17/fiber_product/src/features/authentication/models"
	"github.com/morka17/fiber_product/src/features/authentication/repository"
	"github.com/morka17/fiber_product/src/features/authentication/validators"
	"github.com/morka17/fiber_product/src/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService interface {
	DeleteUser(req *models.DeleteUsersRequest) (*models.DeleteUsersResponse, error)
	GetUser(req *models.GetUserRequest) (*models.User, error)
	ListUsers(req *models.ListUserRequest) ([]*models.User, error)
	UpdateUser(req *models.User) (*models.User, error)
	SignUp(req *models.User) (*models.User, error) 
	SignIn(req *models.SigninRequest) (*models.SigninResponse, error)
}

type authService struct {
	usersRepository repository.UsersRepository
}

func NewAuthService(usersRepository repository.UsersRepository) AuthService {
	return &authService{usersRepository: usersRepository}
}



func (s *authService) SignUp(req *models.User) (*models.User, error ){
	err := validators.ValidateSignUp(req)
	if err != nil {
		return nil, err 
	}

	req.Password, err = security.EncryptPassword(req.Password)
	if err != nil {
		return nil, err 
	}

	req.Name = strings.TrimSpace(req.Name) 
	req.Email	= validators.NormalizeEmail(req.Email)

	found, err := s.usersRepository.GetByEmail(req.Email)
	if err == mongo.ErrNoDocuments {
		user := new(models.User)
		user = req
		err := s.usersRepository.Save(user)
		if err != nil {
			return nil, fmt.Errorf("error in %v", err)
		}
		return user, nil 
	}	

	if found == nil {
		return nil, fmt.Errorf("Invalid arugement %v", err)
	}

	return nil, validators.ErrEmailAlreadyExists

}


func (s *authService) SignIn(req *models.SigninRequest) (*models.SigninResponse, error)  {
	req.Email = validators.NormalizeEmail(req.Email)

	user, err := s.usersRepository.GetByEmail(req.Email)
	if err != nil {
		log.Printf("Signin failed: %v\n", err.Error())
		return nil, validators.ErrSigninFailed
	}

	err = security.VerifyPassword(user.Password, req.Password)
	if err != nil {
		log.Printf("signin failed: %v", err.Error())
		// Invalid username and password 
		return nil, validators.ErrSigninFailed
	}

	// Create new signin token 
	token, err := security.NewToken(user.Id.Hex())
	return &models.SigninResponse{User: *user, Token: token }, nil 
}


/// UPDATE USER
func (s *authService) UpdateUser(req *models.User) (*models.User, error) {
	if !primitive.IsValidObjectID(req.Id.Hex()){
		return nil, validators.ErrInvalidUserId
	}

	user, err := s.usersRepository.GetById(req.Id.Hex())
	if err != nil {
		return nil, err 
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == ""{
		return nil, validators.ErrEmptyName
	}

	user.Name = req.Name 
	user.Updated =  time.Now()

	err =  s.usersRepository.Update(user)
	return user, err 
}

/// GET A SPECIFIC USER 
func (s *authService) GetUser(req *models.GetUserRequest) (*models.User, error) {
	if !primitive.IsValidObjectID(req.Id) {
		return nil, validators.ErrInvalidUserId
	}

	found, err := s.usersRepository.GetById(req.Id)
	if err != nil {
		return nil, err 
	}

	return found, nil 
}



/// List ALL USERS 
func (s *authService) ListUsers(req *models.ListUserRequest) ([]*models.User, error) {
	users, err := s.usersRepository.GetAll()
	if err != nil {
		return nil, err 
	}

	return users, nil 
}


/// DELETE USER 
func (s *authService) DeleteUser(req *models.DeleteUsersRequest) (*models.DeleteUsersResponse, error){
	if !primitive.IsValidObjectID(req.Id){
		return nil, validators.ErrInvalidUserId
	}

	err := s.usersRepository.Delete(req.Id)
	if err != nil {
		return nil, err
	}

	return &models.DeleteUsersResponse{Id: req.Id}, nil 

}
