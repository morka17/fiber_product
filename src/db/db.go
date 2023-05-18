package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)


type Connection interface {
	Close()
	DB() *mongo.Database

}


type conn struct {
	session *mongo.Client
	database  *mongo.Database
}


func NewConnection(cfg Config) (Connection, error) {
	fmt.Printf("Database url: %v\n", cfg.Dsn())

	// Connect to mongoDB 
	session, err := mongo.NewClient(cfg.DbOpts())

	log.Printf("session is: %v", session)
	if err != nil {
		return nil , fmt.Errorf("Session was not created because: %v", err)
	}

	// defer session.Disconnect(context.Background())
    err = session.Connect(context.Background());
	if err != nil  {log.Fatal(err)} 

	return &conn{session: session, database: session.Database(cfg.DbName())}, nil 
	
}


func (c *conn) Close() {
	c.session.Disconnect(context.Background())
}

func (c *conn) DB() *mongo.Database {
	return c.database
}
