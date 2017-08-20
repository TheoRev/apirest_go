package models

import (
	"database/sql"
	"fmt"
	"log"

	// Driver de mysql
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username string = "theo"
const password string = "ambu"
const host string = "localhost"
const port int = 3308
const database string = "project_go_web"

// CreateConnection establece la coneccion con mysql
func CreateConnection() {
	connection, err := sql.Open("mysql", generateURL())
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
}

func truncateTable(tableTable) {

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

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}
