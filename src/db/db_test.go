package db

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}
}

func TestNewConfig(t *testing.T) {

	//if we crash happens
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := NewConfig()
	assert.Equal(t, cfg.DbName(), "microservices2")

}

func TestNewConnection(t *testing.T) {

	cfg := NewConfig()

	conn, err := NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()
}
