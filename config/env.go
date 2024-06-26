package config

import (
	"log"

	"github.com/spf13/viper"
)

var Envs = initConfig()

type Config struct {
	PublicHost     string
	Port           string
	SwaggerAddress string
	SecretTokenJWT string
	JWTExpiredTime int64
	DBConfig
	DBGoose
}

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

type DBGoose struct {
	DBMigrations string
	DBSeeds      string
}

func initConfig() Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		PublicHost:     viper.GetString("PUBLIC_HOST"),
		Port:           viper.GetString("PORT"),
		SwaggerAddress: viper.GetString("SWAGGER_ADDRESS"),
		SecretTokenJWT: viper.GetString("JWT_TOKEN_KEY"),
		JWTExpiredTime: viper.GetInt64("JWT_EXPIRED_TIME"),
		DBConfig: DBConfig{
			DBUser:     viper.GetString("DB_USER"),
			DBPassword: viper.GetString("DB_PASSWORD"),
			DBAddress:  viper.GetString("DB_ADDRESS"),
			DBName:     viper.GetString("DB_NAME"),
		},
		DBGoose: DBGoose{
			DBMigrations: viper.GetString("DB_MIGRATIONS"),
			DBSeeds:      viper.GetString("DB_SEEDS"),
		},
	}
}
