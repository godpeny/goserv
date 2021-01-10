package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"

	model "github.com/godpeny/goserv/init/model"
	db "github.com/godpeny/goserv/internal/goserv-db"
)

func main() {

	// set db
	client, err := model.InitModel()
	if err != nil {
		log.Fatalf("failed init Model : %v", err)
	}
	defer client.Close()

	//mq
	log.Printf("Run RMQ Server)")
	db.InitMQServer()
	db.RunMQServer()

	log.Info("Waiting SIGTERM...")
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	<-sigterm
	log.Info("Receive SIGTERM and exit")
}
