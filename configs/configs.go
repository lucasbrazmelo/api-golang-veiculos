package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Port string
	Host string
	User string
	Password string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.user", "")
	viper.SetDefault("adatabase.password", "")
	viper.SetDefault("adatabase.database", "veiculos")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Port: viper.GetString("database.port"),
		Host: viper.GetString("database.host"),
		User: viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}