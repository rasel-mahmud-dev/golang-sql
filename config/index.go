package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

var (
	PORT = 0
	DATABASE_NAME = ""
	DB_DRIVER = ""
	DATABASE_PORT = 0
	DATABASE_HOST = ""
	DATABASE_USERNAME = ""
	DATABASE_PASSWORD = ""
	SECRET_KEY = ""
	JWT_EXPIRE = time.Hour * 24
)

func Envload()  {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("can't read env file")
	}
	
	
	DB_DRIVER = os.Getenv("DB_DRIVER")
	PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	DATABASE_PORT, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	DATABASE_NAME = os.Getenv("DATABASE_NAME")
	DATABASE_HOST = os.Getenv("DATABASE_HOST")
	DATABASE_USERNAME = os.Getenv("DATABASE_USERNAME")
	DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	SECRET_KEY = os.Getenv("API_SECRET_KEY")
	
}

// @/?charset-utf8&parseTime=True&loc=Local