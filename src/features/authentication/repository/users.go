package repository

import (
	"context"
	"fmt"

	"github.com/morka17/fiber_product/src/db"
	"github.com/morka17/fiber_product/src/features/authentication/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollection = "micro_users"

type UsersRepository interface {
	Save(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Update(user *models.User) error
	Delete(id string) error
	DeleteAll() error
}

type usersRepository struct {
	c *mongo.Collection
}

/// New Instance of the `User` repository
func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{c: conn.DB().Collection(UsersCollection)}
}

func (r *usersRepository) Save(user *models.User) error {
	_, err := r.c.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("Insert operation failed, because: %v", err)
	}

	return nil
}

func (r *usersRepository) GetById(id string) (user *models.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("cannnot parse id")
	}

	filter := bson.M{"_id": oid}
	res := r.c.FindOne(context.Background(), filter)
	if err = res.Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	filter := bson.M{"email": email}
	res := r.c.FindOne(context.Background(), filter)
	if err = res.Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *usersRepository) GetAll() (users []*models.User, err error) {
	var user *models.User
	filter := bson.M{}

	curs, err := r.c.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("Unkown internal error:%v\n", err) // internal error
	}
	defer curs.Close(context.Background())

	for curs.Next(context.Background()) {
		err = curs.Decode(&user)
		if err != nil {
			return nil, fmt.Errorf(
				"Error while decoding data",
			)
		}

		users = append(users, user)
	}

	/// Check if there exist error in the database cursor
	if err = curs.Err(); err != nil {
		return nil, fmt.Errorf("Unkown internal error: %v", err)
	}

	return users, nil
}


func (r *usersRepository) Update(user *models.User) error {
	update := user 

	oid, err := primitive.ObjectIDFromHex(user.Id.Hex())
	if err  !=  nil {
		return fmt.Errorf("Invalid arugment, could not parse ID")
	} 

	filter := bson.M{"_id":oid}
	_, updateErr := r.c.ReplaceOne(context.Background(), filter, update)
	if updateErr != nil {
		return  fmt.Errorf("Failed to update object: %v", updateErr)
	}  

	return nil 
}


func (r *usersRepository) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("cannot parse ID: %v", err)
	}

	filter := bson.M{"_id": oid}

	res, err := r.c.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("Operation failed because : %v", err)
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("User not exist %v:", err)
	}
	
	return nil 
}



func (r *usersRepository) DeleteAll() error {
	filter := bson.M{}
	_, err := r.c.DeleteMany(context.Background(), filter)
	return err
}
