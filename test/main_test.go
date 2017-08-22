package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/TheoRev/apirest_go/models"
)

func TestMain(m *testing.M) {
	result := m.Run()
	os.Exit(result)
}

func beforeTest() {
	fmt.Println("<< Antes de las pruebas")
	models.CreateConnection()
	models.CreateTables()
	createDefaultUser()
}

func createDefaultUser() {
	sql := fmt.Sprintf("insert users set id='%d', username='%s',password='%s',email='%s'", "José", "722", "pepe02@gmail.com")
	_, err := models.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func afterTest() {
	fmt.Println(">> Después de las pruebas")
	models.CloseConnection()
}
