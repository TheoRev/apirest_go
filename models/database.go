package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TheoRev/apirest_go/config"
	// Driver de mysql
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// CreateConnection establece la coneccion con mysql
func CreateConnection() {
	url := config.GetUrlDatabase()
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	} else {
		db = connection
		fmt.Println("Conexión exitosa!")
	}
}

// CreateTables crea todas las tablas necesarias
func CreateTables() {
	createTable("users", userSchema)
}

// CreateTable crea la tabla en la db
func createTable(tableName, schema string) {
	if !existsTable(tableName) {
		Exec(schema)
	}
	// else {
	// 	truncateTable(tableName)
	// }
}

func truncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

// Exec overwrite function
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

// Query overwrite function
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	return rows, err
}

// CloseConnection cierra la conexion con la db
func CloseConnection() {
	db.Close()
}

// ExistsTable verifica si existe la tabla en la db
func existsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

// Ping prueba la conexion con la db
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
