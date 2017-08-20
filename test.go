package main

import (
	"github.com/TheoRev/apirest_go/models"
	"fmt"
	"github.com/TheoRev/apirest_go/orm"
)

func main() {
	// testCrud()
	testOrm()
}

func testOrm()  {
	orm.CreateConnection()
	orm.CloseConnection()
}

func testCrud()  {
	models.CreateConnection()
	// models.CreateTables()
	// user := models.CreateUser("Almendra", "8773", "almindra12@gmail.com")
	// user.Save()

	// user := models.GetUser(6)
	// fmt.Println(user)

	users:=models.GetUsers()
	fmt.Println(users)
	models.CloseConnection()
}
