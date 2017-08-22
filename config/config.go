package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

// DatabaseConfig estructura con atributos de configuracion
type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Databse  string
	Debug    bool
}

var database *DatabaseConfig

func init() {
	database = &DatabaseConfig{}
	database.Username = gonv.GetStringEnv("USER", "theo")
	database.Password = gonv.GetStringEnv("PASSWORD", "ambu")
	database.Host = gonv.GetStringEnv("HOST", "localhost")
	database.Port = gonv.GetIntEnv("PORT", 3308)
	database.Databse = gonv.GetStringEnv("DATABASE", "project_go_web")
	database.Debug = gonv.GetBoolEnv("DEBUG", true)
}

// GetUrlDatabase obtiene la cadena de coneccion con la db
func GetUrlDatabase() string {
	return database.url()
}

func (this *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", this.Username, this.Password, this.Host, this.Port, this.Databse)
}

func GetDebug() bool {
	return database.Debug
}
