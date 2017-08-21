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
	orm.CreateTables()
	user := orm.NewUser("Almendra", "8773", "almindra12@gmail.com")
	user.Save()
	// fmt.Println(user)
	users := orm.GetUsers()
	fmt.Println(users)
	user = orm.GetUser(1)
	user.Username = "Almendra Cárdenas Mesa"
	user.Password = "878"
	user.Update()
	fmt.Println(user)
	user.Delete()
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
