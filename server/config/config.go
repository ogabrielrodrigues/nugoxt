package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
	Cors string
}

type DBConfig struct {
	ConnString string
}

func init() {
	viper.SetDefault("api.port", 9000)
	viper.SetDefault("api.cors", "*")
	viper.SetDefault("database.conn_string", "postgres://goshop:goshopdev@localhost:5432/goshop_dev")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = &config{
		API: APIConfig{
			Port: viper.GetString("api.port"),
			Cors: viper.GetString("api.cors"),
		},
		DB: DBConfig{
			ConnString: viper.GetString("database.conn_string"),
		},
	}

	return nil
}

func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetAPIConfig() APIConfig {
	return cfg.API
}
