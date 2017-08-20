package main

import "github.com/TheoRev/apirest_go/models"
import "fmt"

func main() {
	models.CreateConnection()
	models.CreateTables()
	user := models.CreateUser("Chris", "23234", "chris277@gmail.com")
	fmt.Println(user)
	models.CloseConnection()
}
