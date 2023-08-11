package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBUSER          string
	DBPASSWORD      string
	DBNAME          string
	DBHOST          string
	DBPORT          string
	HTTPPORT        string
	MODE            string
	AWSREGION       string
	PHOTOBUCKETNAME string
)

func init() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	filename := ".env." + env
	e := godotenv.Load(filename)
	if e != nil {
		fmt.Println("couldn't load env file")
		panic(e)
	}
	loadDBVars()
	loadAppVars()
	loadAwsVars()
}

func loadDBVars() {
	DBUSER = os.Getenv("DB_USER")
	DBPASSWORD = os.Getenv("DB_PASSWORD")
	DBNAME = os.Getenv("DB_NAME")
	DBHOST = os.Getenv("DB_HOST")
	DBPORT = os.Getenv("DB_PORT")
}

func loadAppVars() {
	HTTPPORT = fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	MODE = os.Getenv("MODE")
}

func loadAwsVars() {
	AWSREGION = os.Getenv("REGION")
	PHOTOBUCKETNAME = os.Getenv("S3_PHOTO_BUCKET_NAME")
}
