package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/godpeny/goserv/init/model"
	sw "github.com/godpeny/goserv/internal/goserv-api/go"
)

func main() {
	// set db
	client, err := model.InitModel()
	if err != nil {
		log.Fatalf("failed init Model : %v", err)
	}
	defer client.Close()

	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(router.Run(":8080"))
}
