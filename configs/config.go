package configs

import (
	"log"
	"os"
)

type PortConfig struct {
	Value string
}

type MongoDbConfig struct {
	Url string
}

type JwtConfig struct {
	Key string
}

type Config struct {
	PORT    PortConfig
	Mongo   MongoDbConfig
	JWT_KEY JwtConfig
}

func LoadConfig() *Config {
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8000"
	}
	Port := PortConfig{Value: portString}
	MongoDbUrlString := os.Getenv("MONGODBURI")
	if MongoDbUrlString == "" {
		log.Fatal("Doesn't find any database url to connect with..!!")
	}
	MongoDb := MongoDbConfig{
		Url: MongoDbUrlString,
	}
	jwtString := os.Getenv("JWT_KEY")
	if (jwtString == "") {
		log.Fatal("Couldn't Find jwt key");
	}
	Jwt := JwtConfig {
		Key: jwtString,
	}

	return &Config{
		PORT:  Port,
		Mongo: MongoDb,
		JWT_KEY: Jwt,
	}
}
