package db

import (
	"encoding/json"
	"log"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	ent "github.com/godpeny/goserv/ent"

	"github.com/streadway/amqp"
)

var (
	MQ_Client MessageQueue
	APIc      chan int
)

func InitMQClient() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	mq := MessageQueue{
		Conn:     conn,
		Channel:  ch,
		Queue:    q,
		Delivery: msgs,
	}

	MQ_Client = mq
	APIc = make(chan int)
}

func RunMQ_User(msgType string, req ent.User, c *gin.Context) {
	go func() {
		//conn := MQ_Client.Conn
		//defer conn.Close()

		ch := MQ_Client.Channel
		//defer ch.Close()

		q := MQ_Client.Queue

		corrId := randomString(32)

		// Marshalling
		body, err := json.Marshal(req)
		if err != nil {
			failOnError(err, "Failed to Marshal req")
		}

		err = ch.Publish(
			"",          // exchange
			"rpc_queue", // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: corrId,
				ReplyTo:       q.Name,
				Type:          msgType,
				Body:          body,
			})
		failOnError(err, "Failed to publish a message")

		for d := range MQ_Client.Delivery {
			if corrId == d.CorrelationId {
				res, err := strconv.Atoi(string(d.Body))
				log.Printf("[.] Got %d", res)
				failOnError(err, "Failed to convert body to integer")

				APIc <- res
				break
			}
		}
		return
	}()
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
