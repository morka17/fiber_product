package restutils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morka17/fiber_product/src/security"
)

var (
	ErrEmptyBody = errors.New("Body cannot be null")
	ErrUnauthorized = errors.New("Unauthorized")
)



type JError struct {
	Error string `json:"error"`
}



func WriteAsJson(w http.ResponseWriter, statusCode int, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}



func WriteError(w http.ResponseWriter, statusCode int, err error){
	e := "error"
	if err != nil {
		e = err.Error()
	}

	WriteAsJson(w, statusCode, JError{e})
}


///	.......
///	Authenticate request 
func AuthRequestWithId(r *http.Request) (*security.TokenPayload, error){
	token, err := security.ExtractToken(r)
	if err != nil {
		return nil, err 
	}

	payload, err := security.NewTokenPayload(token)
	if err != nil {
		return nil, err 
	}

	vars := mux.Vars(r)
	if payload.UserId != vars["id"] {
		return nil, ErrUnauthorized
	}

	return payload, nil 
}
