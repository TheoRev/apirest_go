package test

import (
	"testing"

	"github.com/TheoRev/apirest_go/models"
)

var user *models.User

func TestNewUser(t *testing.T) {
	user := models.NewUser("username", "password", "email")
	if user.Username != "username" || user.Password != "password" || user.Email != "email" {
		t.Error("No es posible crear el objeto")
	}
}

func TestSave(t *testing.T) {
	user := models.NewUser("José", "722", "pepe02@gmail.com")
	if err := user.Save(); err != nil {
		t.Error("No es posible crear el usuario", err)
	}
}

func TestCreateUser(t *testing.T) {
	_, err := models.CreateUser("Anita", "454", "anita44@gmail.com")
	if err != nil {
		t.Error("No es posible insertar el objeto", err)
	}
}

func TestUniqueUsername(t *testing.T) {
	_, err := models.CreateUser("José", "722", "pepe02@gmail.com")
	if err != nil {
		t.Error("Es posible insertar registros con usernames duplicados")
	}
}

func TestGetUser(t *testing.T) {
	user := models.GetUser(1)
	if user.Username != "Theo" || user.Password != "123" || user.Email != "theorev@gmail.com" {
		t.Error("No es posible obtener el usuario", nil)
	}
}

func TestGetUsers(t *testing.T) {
	users := models.GetUsers()
	if len(users) == 0 {
		t.Error("No es posible obtener a los usuarios")
	}
}

func TestDeleteUser(t *testing.T) {

}
