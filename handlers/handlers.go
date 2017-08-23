package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/TheoRev/go_web/rest/models"
	"github.com/gorilla/mux"
)

// GetUsers Se obtiene todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

// GetUser Se obtiene un usuario
func GetUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}

// CreateUser Se crea un usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	} 

	if err:= user.Valid() != nil {
		models.SendUnprocessableEntity(w)
		return
	}
}

// UpdateUser Se actualiza un usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	userResponse := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	user = models.UpdateUser(user, userResponse.Username, userResponse.Password)
	models.SendData(w, user)
}

// DeleteUser Se elimina un usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	} else {
		models.DeleteUser(user.Id)
		models.SendNoContent(w)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.GetUserById(userId)

	if user.Id == 0 {
		return user, errors.New("El usuario no existe en la base de datos.")
	}

	return user, nil
}
