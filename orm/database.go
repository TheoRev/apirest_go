package orm

import (
	"fmt"

	"github.com/TheoRev/apirest_go/config"
	"github.com/jinzhu/gorm"
	// Driver de mysql
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

// CreateConnection establece la coneccion con mysql
func CreateConnection() {
	url := config.GetUrlDatabase()
	connection, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	} else {
		db = connection
		fmt.Println("Conexión exitosa!")
	}
}

// CloseConnection cierra la conexion con la db
func CloseConnection() {
	db.Close()
}
