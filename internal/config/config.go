package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Login struct {
	User string
	Pass string
	Ip   string
	Port string
	Name string
}

func CreateEnvFile() error {
	err := os.WriteFile(".env", nil, 0o666)
	if err != nil {
		return err
	}
	return nil
}

func LoadCredentials() Login {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Login{
		User: os.Getenv("DATABASE_USER"),
		Pass: os.Getenv("DATABASE_PASSWORD"),
		Ip:   os.Getenv("DATABASE_IP"),
		Port: os.Getenv("DATABASE_PORT"),
		Name: os.Getenv("DATABASE_NAME"),
	}
}
