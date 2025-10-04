package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct{
	DBHost string
	DBPORT string
	DBuser string
	DBpassword string
	DBName string
 }

 var ENVS = initConfig()

 func initConfig() config{
	godotenv.Load()

	return config{
		DBHost: getENV("DB_HOST","127.0.0.1"),
		DBPORT: getENV("DB_PORT","5432"),
		DBpassword: getENV("DB_PASSWORD","2004"),
		DBuser: getENV("DB_USER","jet"),
		DBName: getENV("DB_NAME","nexora_DB"),
	}
 }

 func getENV(key,fallback string) string{
	if value, ok := os.LookupEnv(key); ok{  //LookupEnv returns value, bool
		return  value
	}
	return  fallback
 }