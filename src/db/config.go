package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo/options"
)




type Config interface {
	Dsn()	string 
	DbName() string 
	DbOpts()	*options.ClientOptions
}


type config struct {
	dbUser  string
	dbPass	string 
	dbHost 	string
	dbPort 	int 
	dbName 	string 
	dsn 	string 
	opts    *options.ClientOptions
}


func NewConfig() Config {
	var cfg config 
	var err error 

	cfg.dbUser = os.Getenv("DATABASE_USER")
	cfg.dbPass = os.Getenv("DATABASE_PASS")
	cfg.dbHost = os.Getenv("DATABASE_HOST")
	cfg.dbName = os.Getenv("DATABASE_NAME")
	cfg.dbPort, err =  strconv.Atoi(os.Getenv("DATABASE_PORT"))
	
	if err != nil {
		log.Fatalf("Error on load env var : %v", err.Error())
	}

	credential := options.Credential{
		Username: cfg.dbUser,
		Password: cfg.dbPass,
	} 

	cfg.dsn = fmt.Sprintf("mongodb://%s:%d/%s", cfg.dbHost, cfg.dbPort, cfg.dbName)


	cfg.opts = options.Client().ApplyURI(cfg.dsn).SetAuth(credential)

	return &cfg

}



/// RETURN THE DSN 
func (c *config) Dsn() string {
	return c.dsn
}

/// RETURNS THE NAME 
func (c *config) DbName() string {
	return c.dbName
}

func (c *config) DbOpts() *options.ClientOptions{
	return c.opts
}