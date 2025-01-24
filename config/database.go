package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro loading .env file")
	}

	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// var err error

	db_params := fmt.Sprintf("host=localhost user=postgres password='%s' dbname='%s' sslmode=disable", password, dbname)

	DB, err = sql.Open("postgres", db_params)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Banco de dados não está acessível: ", err)
	}
	log.Println("Banco de dados conectado com sucesso!")
}
