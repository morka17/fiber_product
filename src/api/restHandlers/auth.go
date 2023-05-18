package resthandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/morka17/fiber_product/src/api/restutils"
	"github.com/morka17/fiber_product/src/features/authentication/models"

	authservice "github.com/morka17/fiber_product/src/features/authentication/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandlers interface {
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authSvcClient authservice.AuthService
}

func NewAuthHandlers(authSvc authservice.AuthService) AuthHandlers {
	return &authHandler{authSvcClient: authSvc}
}

func (h *authHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutils.WriteError(w, http.StatusBadRequest, restutils.ErrEmptyBody)
		return
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		log.Printf("cannot read %v\n", r.Body)
		return
	}

	user := new(models.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	/// Better placed in `int64`
	user.Created = time.Now()
	user.Updated = user.Created
	/// #[warning] use default ID
	user.Id = primitive.NewObjectID()

	resp, err := h.authSvcClient.SignUp(user)
	if err != nil {
		restutils.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	restutils.WriteAsJson(w, http.StatusCreated, resp)
}

func (h *authHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutils.WriteError(w, http.StatusBadRequest, restutils.ErrEmptyBody)
		return
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		log.Printf("cannot read %v\n", r.Body)
		return
	}

	input := new(models.SigninRequest)
	err = json.Unmarshal(body, input)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.authSvcClient.SignIn(input)
	if err != nil {
		restutils.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	restutils.WriteAsJson(w, http.StatusOK, resp)
}

func (h *authHandler) PutUser(w http.ResponseWriter, r *http.Request) {

	tokenPayload, err := restutils.AuthRequestWithId(r)
	if err != nil {

	}

	if r.Body == nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := new(models.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		restutils.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Id, err = primitive.ObjectIDFromHex(tokenPayload.UserId)
	if err != nil {
		restutils.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	resp, err := h.authSvcClient.UpdateUser(user)

	restutils.WriteAsJson(w, http.StatusOK, resp)
}

func (h *authHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	tokenPayload, err := restutils.AuthRequestWithId(r)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.authSvcClient.GetUser(&models.GetUserRequest{Id: tokenPayload.UserId})
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	restutils.WriteAsJson(w, http.StatusOK, resp)
}

func (h *authHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.authSvcClient.ListUsers(&models.ListUserRequest{})
	if err != nil {
		restutils.WriteError(w, http.StatusOK, err)
		return
	}

	restutils.WriteAsJson(w, http.StatusOK, users)
}


func (h *authHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	tokenPayload, err := restutils.AuthRequestWithId(r)
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return 
	}
	resp, err := h.authSvcClient.DeleteUser(&models.DeleteUsersRequest{Id: tokenPayload.UserId})
	if err != nil {
		restutils.WriteError(w, http.StatusBadRequest, err)
		return 
	}

	w.Header().Set("Entity", resp.Id)
	restutils.WriteAsJson(w, http.StatusNoContent, nil)
}