package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	ent "github.com/godpeny/goserv/ent"
	ent_client "github.com/godpeny/goserv/internal/clients/ent"

	"github.com/streadway/amqp"
)

const (
	defaultExchange = "default"
	rpc             = "rpc"
)

type MessageQueue struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel

	Queue    amqp.Queue
	Delivery <-chan amqp.Delivery
}

var (
	MQ MessageQueue
)

func InitMQServer() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		"default", // name
		"direct",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,    // queue name
		"rpc",     // routing key
		"default", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
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

	MQ = mq
}

func RunMQServer() {
	go func() {
		conn := MQ.Conn
		defer conn.Close()

		ch := MQ.Channel
		defer ch.Close()

		forever := make(chan bool)

		go func() {
			for d := range MQ.Delivery {
				processExchange(conn, ch, d)
			}
		}()

		log.Printf(" [*] Awaiting RPC requests")
		<-forever
	}()
}

func processExchange(conn *amqp.Connection, ch *amqp.Channel, delivery amqp.Delivery) {
	excName := delivery.Exchange
	switch excName {
	case defaultExchange:
		processRoutingKey(conn, ch, delivery)
	}

}

func processRoutingKey(conn *amqp.Connection, ch *amqp.Channel, delivery amqp.Delivery) {
	routingKey := delivery.RoutingKey

	switch routingKey {
	case rpc:
		processRoutingKey_RPC(conn, ch, delivery)
	}
}

func processRoutingKey_RPC(conn *amqp.Connection, ch *amqp.Channel, delivery amqp.Delivery) {
	ctx := context.Background()

	if delivery.Type == "READ" {
		request := ent.User{}
		if err := json.Unmarshal(delivery.Body, &request); err != nil {
			log.Println(err, "Failed to Unmarshal req")
		}

		client, err := ent_client.GetClient()
		if err != nil {
			log.Println(err, "Failed to Get Ent Client")
		}

		lst, err := client.User.Query().All(ctx)
		if err != nil {
			log.Println(err, "Failed to Create User db")
		}

		for idx, val := range lst {
			fmt.Println(idx, " : ", val.Name)
		}

		err = ch.Publish(
			"",               // exchange
			delivery.ReplyTo, // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: delivery.CorrelationId,
				Body:          []byte(strconv.Itoa(2)),
			})

		if err != nil {
			log.Println(err, "Failed to publish a message")
		}
	} else if delivery.Type == "CREATE" {

		request := ent.User{}
		if err := json.Unmarshal(delivery.Body, &request); err != nil {
			log.Println(err, "Failed to Unmarshal req")
		}

		client, err := ent_client.GetClient()
		if err != nil {
			log.Println(err, "Failed to Get Ent Client")
		}

		_, err = client.User.Create().SetAge(int(request.Age)).SetName(request.Name).Save(ctx)
		if err != nil {
			log.Println(err, "Failed to Create User db")
		}

		err = ch.Publish(
			"",               // exchange
			delivery.ReplyTo, // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: delivery.CorrelationId,
				Body:          []byte(strconv.Itoa(1)),
			})

		if err != nil {
			log.Println(err, "Failed to publish a message")
		}
	}
	delivery.Ack(false)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
