package test

import (
	"testing"

	"github.com/TheoRev/apirest_go/models"
)

func TestConnection(t *testing.T) {
	connection := models.GetConnection()
	if connection == nil {
		t.Error("No es posible realizar la conexion", nil)
	}
}
