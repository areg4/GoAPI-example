package types

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rate   float32 `json:"rate"`
	Resume string  `json:"resume"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MovieResponse struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rate   float32 `json:"rate"`
	Resume string  `json:"resume"`
}

type DatabaseConnection struct {
	User    string `json:"db_user" mapstructure:"DB_USER"`
	Pass    string `mapstructure:"DB_PASS" json:"db_pass"`
	Db_name string `mapstructure:"DB_NAME" json:"db_name"`
	Port    int    `mapstructure:"DB_PORT" json:"db_port"`
	Host    string `mapstructure:"DB_HOST" json:"db_host"`
	Driver  string `mapstructure:"DB_DRIVER" json:"db_driver"`
}

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}
