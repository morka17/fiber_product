package repository

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/morka17/fiber_product/src/db"
	"github.com/morka17/fiber_product/src/features/authentication/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Panicln(err)
	}
}

func TestUsersRepositorySave(t *testing.T) {
	/// if we crash happens
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)

	defer conn.Close()

	id := primitive.NewObjectID()

	password := "password1"

	user := &models.User{
		Id:       id,
		Name:     "TEST1",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: password,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	/// Return the database collection
	r := NewUsersRepository(conn)
	/// ...
	/// Save user to db
	err = r.Save(user)
	assert.NoError(t, err)
}

func TestUsersRepositoryGetById(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)

	defer conn.Close()

	id := primitive.NewObjectID()

	password := "password1"

	user := &models.User{
		Id:       id,
		Name:     "TEST2",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: password,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	/// Return the database collection
	r := NewUsersRepository(conn)

	/// ... Save user to db
	err = r.Save(user)
	assert.NoError(t, err)

	///	...
	/// Retrieve a collection by `id`
	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)

	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)

}

func TestUsersRepositoryGetByEmail(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)

	defer conn.Close()

	id := primitive.NewObjectID()

	password := "password1"

	user := &models.User{
		Id:       id,
		Name:     "TEST2",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: password,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	/// Return the database collection
	r := NewUsersRepository(conn)

	/// ... Save user to db
	err = r.Save(user)
	assert.NoError(t, err)

	///	...
	/// Retrieve a collection by `id`
	found, err := r.GetByEmail(user.Email)
	assert.NoError(t, err)

	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)

}

func TestUsersRepositoryUpdate(t *testing.T) {
	//if we crash happens
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "1234566",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn) // return a database collec
	err = r.Save(user)            // saves user in the collection
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	chg := "UPDATE"

	found.Name = chg

	err = r.Update(found)
	assert.Nil(t, err)

	found, err = r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, chg, found.Name)

}



func TestUsersRespositoryDelete(t *testing.T) {
	//if we crash happens
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "1234566",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn) // return a database collec
	fmt.Print(r)
	err = r.Save(user) // saves user in the collection
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)

	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)

	err = r.Delete(user.Id.Hex())
	assert.Nil(t, err)

	found, err = r.GetById(user.Id.Hex())
	assert.Error(t, err)
	assert.EqualError(t, mongo.ErrNoDocuments, err.Error())
	assert.Nil(t, found)
}
