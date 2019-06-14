package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sohamkamani/go-job-processing-example/queue"
)

func main() {
	// initialize the queue connection
	queue.Init("amqp://localhost")
	// start a publisher or worker depending on the command line argument
	if os.Args[1] == "worker" {
		worker()
	} else {
		publisher()
	}
}

func publisher() {
	// the publisher publishes the message "1,1" every 500 milliseconds, perpetually
	for {
		if err := queue.Publish("add_q", []byte("1,1")); err != nil {
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func worker() {
	// obtain the channel which we subscribe to
	msgs, close, err := queue.Subscribe("add_q")
	if err != nil {
		panic(err)
	}
	defer close()
	forever := make(chan bool)

	go func() {
		// Receive messages from the channel forever
		for d := range msgs {
			// everytime a message is received, convert it to numbers, add the numbers
			i1, i2 := toNums(d.Body)
			// then print the result to STDOUT, along with the time
			fmt.Println(time.Now().Format("01-02-2006 15:04:05"), "::", i1+i2)
			// acknowledge the message so that it is cleared from the queue
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func toNums(b []byte) (int, int) {
	s := string(b)
	ss := strings.Split(s, ",")
	i1, _ := strconv.Atoi(ss[0])
	i2, _ := strconv.Atoi(ss[1])
	return i1, i2
}

// queue/queue.go
package queue

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection

func Init(c string) {
	var err error
	// Initialize the package level "conn" variable that represents the connection the the rabbitmq server
	conn, err = amqp.Dial(c)
	if err != nil {
		log.Fatalf("could not connect to rabbitmq: %v", err)
		panic(err)
	}
}

func Publish(q string, msg []byte) error {
	// create a channel through which we publish
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// create the payload with the message that we specify in the arguments
	payload := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         msg,
	}

	// publish the message to the queue specified in the arguments
	if err := ch.Publish("", q, false, false, payload); err != nil {
		return fmt.Errorf("[Publish] failed to publish to queue %v", err)
	}

	return nil
}

func Subscribe(qName string) (<-chan amqp.Delivery, func(), error) {
	// create a channel through which we publish
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	// assert that the queue exists (creates a queue if it doesn't)
	q, err := ch.QueueDeclare(qName, false, false, false, false, nil)

	// create a channel in go, through which incoming messages will be received
	c, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	// return the created channel
	return c, func() { ch.Close() }, err
}
