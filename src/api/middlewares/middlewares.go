package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/morka17/fiber_product/src/api/restutils"
	"github.com/morka17/fiber_product/src/security"
)

func LogRequests(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next(rw, r)
		log.Printf(`{"proto": "%s","method:" , "%s", "route":"%s%s", "request_time": "%v"}`,
			r.Proto, r.Method, r.Host, r.URL.Path, time.Since(t))
	}
}


func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tokenString, err := security.ExtractToken(r)
		if err != nil {
			log.Println("error on extract token from header:", err.Error())
			restutils.WriteAsJson(rw, http.StatusUnauthorized, restutils.ErrUnauthorized)
			return 
		}
		token, err := security.ParseToken(tokenString)
		if err != nil {
			log.Println("error on parse token", err.Error())
			restutils.WriteAsJson(rw, http.StatusUnauthorized, restutils.ErrUnauthorized)
			return 
		}

		if !token.Valid {
			log.Println("error on parse token:", tokenString)
			restutils.WriteAsJson(rw, http.StatusUnauthorized, restutils.ErrUnauthorized)
			return 
		}
		next(rw, r )

	}
}
