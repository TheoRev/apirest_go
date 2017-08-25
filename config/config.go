package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

// DatabaseConfig estructura con atributos de configuracion
type DatabaseConfig struct {
	username string
	password string
	host     string
	port     int
	databse  string
	debug    bool
}

// ServerConfig estrucctura con parametros del servidor
type ServerConfig struct {
	host  string
	port  int
	debug bool
}

var database *DatabaseConfig
var server *ServerConfig

func init() {
	database = &DatabaseConfig{}
	database.username = gonv.GetStringEnv("USER", "theo")
	database.password = gonv.GetStringEnv("PASSWORD", "ambu")
	database.host = gonv.GetStringEnv("HOST", "localhost")
	database.port = gonv.GetIntEnv("PORT", 3308)
	database.databse = gonv.GetStringEnv("DATABASE", "project_go_web")
	database.debug = gonv.GetBoolEnv("DEBUG", true)

	server = &ServerConfig{}
}

func DirTemplate() string {
	return "templates/**/*.html"
}

func DirTemplateError() string {
	return "templates/error.html"
}

// GetUrlDatabase obtiene la cadena de coneccion con la db
func GetUrlDatabase() string {
	return database.url()
}

func (this *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", this.username, this.password, this.host, this.port, this.databse)
}

func Debug() bool {
	return server.debug
}

func (this *ServerConfig) url() string {
	return fmt.Sprintf("%s:%d", this.host, this.port)
}
