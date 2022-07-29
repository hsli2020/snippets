// https://www.abilityrush.com/working-with-aws-simple-queue-service-sqs-in-golang/
package main

import (
	"flag"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var shouldDelete, push, pop bool

func main() {

	flag.BoolVar(&shouldDelete, "delete", false, "delete after reading")
	flag.BoolVar(&push, "push", false, "push the message")
	flag.BoolVar(&pop, "pop", false, "pop message")
	flag.Parse()

	if !push && !pop {
		log.Fatalln("At least push or pop flag is required.")
	}

	svc, qurlOut := createSQSObj()

	if push {
		pushToQueue(svc, qurlOut.QueueUrl)
	}

	if pop {
		readFromQueue(svc, qurlOut.QueueUrl)
	}

}

func createSQSObj() (*sqs.SQS, *sqs.GetQueueUrlOutput) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
	queueName := "SQS_TEST_QUEUE"

	qurlOut, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return svc, qurlOut
}

func pushToQueue(svc *sqs.SQS, qurl *string) {

	now := time.Now().Local().String()
	msgBody := "Hi, I am entered the queue at " + now

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: &msgBody,
		QueueUrl:    qurl,
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Pushed to queue, Data : %v", msgBody)
}

func readFromQueue(svc *sqs.SQS, qurl *string) {

	var maxNumMsg int64 = 1

	rmo, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            qurl,
		MaxNumberOfMessages: &maxNumMsg,
	})

	if err != nil {
		log.Fatalln(err)
	}

	if len(rmo.Messages) <= 0 {
		log.Println("No Messages to read.")
	}

	for _, msg := range rmo.Messages {
		log.Printf("Read from queue, Data : %v", *msg.Body)

		if shouldDelete {
			_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      qurl,
				ReceiptHandle: msg.ReceiptHandle,
			})

			if err != nil {
				log.Fatalln(err)
			}

			log.Printf("Delete from queue success")
		}

	}
}
