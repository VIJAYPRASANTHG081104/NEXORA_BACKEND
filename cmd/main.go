package main

import (
	"database/sql"
	"fmt"
	"log"
	"nexora_backend/api/server"
	"nexora_backend/config"
	"nexora_backend/pkg/db"
)

func main(){
	db, err := db.NewPostgresSQL(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	config.ENVS.DBHost,
	config.ENVS.DBPORT,
	config.ENVS.DBuser,
	config.ENVS.DBpassword,
	config.ENVS.DBName))

	if err != nil {
		log.Fatal(err)
	}

	if err := initDB(db);err != nil{
		log.Fatal("connition err",err)
	}
	apiServer := server.NewAPIServer(":8080",db)
	
	if err := apiServer.Run(); err != nil{
		log.Fatal("error while running the API Server")
	}
}

func initDB(db *sql.DB) error{
	err := db.Ping()
	if err != nil{
		return err
	}
	fmt.Println("DB connection success")
	return nil;
}