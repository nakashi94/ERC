package main

import (
	"log"
	"web-service-api/infrastructures/database"
	"web-service-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの環境変数を読み込む
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// databaseと接続する
	conn, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	// databaseの接続を切断する
	defer func() {
		if err = conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	router := routes.NewRouter()

	router.SetMiddleware()
	router.SetProxy()
	router.SetHealthChecker()
	router.NewTaskRouter(conn)

	router.Serve()
}
