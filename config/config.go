package config

import (
	"GoAPI/types"
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() (config types.Config, err error) {
	fmt.Printf("Cargando config...")
	database_connection := types.DatabaseConnection{}
	viper.AddConfigPath("config")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&database_connection)
	if err != nil {
		return
	}
	// fmt.Println(viper.Get("CORS"))

	loadConfigDB(&database_connection, &config)

	return
}

func loadConfigDB(database_connection *types.DatabaseConnection, config *types.Config) {
	config.DBDriver = database_connection.Driver
	config.DBSource = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Mexico_City",
		database_connection.Host,
		database_connection.User,
		database_connection.Pass,
		database_connection.Db_name,
		database_connection.Port,
	)
	config.ServerAddress = database_connection.Host
}
