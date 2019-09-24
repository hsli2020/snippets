package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
)

const BTEROOT = "c:/xampp/htdocs/btenew"

func dpr(v interface{}) {
	fmt.Printf("%#v\n", v)
}

func php(args ...string) error {
	cmd := exec.Command("c:/xampp/php/php", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = BTEROOT + "/job"
	return cmd.Run()
}

func php64(args ...string) error {
	cmd := exec.Command("c:/xampp/php64/php", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = BTEROOT + "/job"
	return cmd.Run()
}

func main() {
	// Setup Logger
	logfile := BTEROOT + "/app/logs/jobworker.log"
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file")
	}

	multi := io.MultiWriter(file, os.Stdout)
	//logger := log.New(multi, "", log.LstdFlags)
	log.SetOutput(multi)
	defer file.Close()

	jobQueue := NewJobQueue("job:queue")

	//msg := []string{"Test.php", "arg1", "arg2"}
	//jobQueue.Push(msg)

	job, err := jobQueue.Pop(1)
	if err == nil {
		log.Println("Run Job: " + strings.Join(job, " "))
		err = php64(job...)
		if err != nil {
			log.Println(err)
		}
	}
}

type JobQueue struct {
	name  string
	redis *redis.Client
}

func NewJobQueue(name string) *JobQueue {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	return &JobQueue{
		redis: client,
		name:  name,
	}
}

func (q JobQueue) Push(message []string) error {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return err
	}

	if err = q.redis.RPush(q.name, string(data)).Err(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (q JobQueue) Pop(d time.Duration) ([]string, error) {
	result, err := q.redis.BLPop(d*time.Second, q.name).Result()
	if err != nil {
		//log.Println(err)
		return nil, err
	}
	//dpr(result)

	var message []string
	err = json.Unmarshal([]byte(result[1]), &message)
	if err != nil {
		//log.Println(err)
		return nil, err
	}
	//dpr(message)

	return message, nil
}
