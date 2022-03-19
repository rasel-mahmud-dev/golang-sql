package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var (
	PORT = "4000"
	DATABASE_NAME = ""
	DB_DRIVER = ""
	DATABASE_PORT = ""
	DATABASE_HOST = ""
	MYSQL_USER = ""
	MYSQL_PASSWORD = ""
	SECRET_KEY = ""
	JWT_EXPIRE = time.Hour * 24
)

func Envload()  {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("can't read env file")
	}


	DB_DRIVER = os.Getenv("DB_DRIVER")
	PORT = os.Getenv("PORT")
	DATABASE_PORT = os.Getenv("DATABASE_PORT")
	DATABASE_NAME = os.Getenv("MYSQL_DATABASE")
	DATABASE_HOST = os.Getenv("DATABASE_HOST")
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	SECRET_KEY = os.Getenv("API_SECRET_KEY")

}

// @/?charset-utf8&parseTime=True&loc=Local