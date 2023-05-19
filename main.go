package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
	resthandlers "github.com/morka17/fiber_product/src/api/restHandlers"
	"github.com/morka17/fiber_product/src/api/routes"
	"github.com/morka17/fiber_product/src/db"
	"github.com/morka17/fiber_product/src/features/authentication/repository"
	authservice "github.com/morka17/fiber_product/src/features/authentication/service"
)

var (
	port     int
	authAddr string
)

func init() {
	flag.IntVar(&port, "port", 9000, "api service port")
	flag.StringVar(&authAddr, "auth addr", "localhost:9001", "authentication service address")
	flag.Parse()

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Panicln(err)
	// }
}

func main() {

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	authRepository := repository.NewUsersRepository(conn)
	authservice := authservice.NewAuthService(authRepository)
	authHandlers := resthandlers.NewAuthHandlers(authservice)

	authRoutes := routes.NewAuthRoutes(authHandlers)

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, authRoutes)

	log.Printf("API service running on [::%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routes.WithCORS(router)))

}
