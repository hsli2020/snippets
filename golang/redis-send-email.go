import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/go-redis/redis/v8"
	"github.com/keithwachira/go-taskq"
)
//this is the redis key we want to use for our stream
var streamName = "send_order_emails"
func main() {
   //we start by creating a new go-redis client
   //that we will use to access our redis instance
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})
	//start processing any received job in redis
    //you can see the definition of this function below
	go StartProcessingEmails(rdb)
    //to mimic user behaviour we will create a single end point to send email requests.
	handler := http.NewServeMux()
	///we create a new router to expose our api
	//to our users
	s := Server{Redis: rdb}
	handler.HandleFunc("/api/order", s.NewOrderReceivedFromClient)
	//Every time a  request is sent to the endpoint ("/api/order")
	// will mock sending an email
	err := http.ListenAndServe("0.0.0.0:8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	Redis *redis.Client
}
// NewOrderReceivedFromClient this is mocking an endpoint that users use to place an order
//once we receive an order here we should register it to redis
//then the workers will pick it up and send an email to the user
func (S *Server) NewOrderReceivedFromClient(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"email": "redis@gmail.com", "message": "We have received you order and we are working on it."}
	//we have received  an order here send it to
	//redis has a function called xadd that we will use to add this to our stream
     //you can read more about it on the link shared above.
	err := S.Redis.XAdd(context.Background(), &redis.XAddArgs{
		///this is the name we want to give to our stream
		///in our case we called it send_order_emails
		//note you can have as many stream as possible
		//such as one for email...another for notifications
		Stream:       streamName,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		//values is the data you want to send to the stream
		//in our case we send a map with email and message keys
		Values: data,
	}).Err()
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, `We received you order`)
}

// RedisStreamsProcessing 
//you can also pass a database here too
//if you need it to process you work
type RedisStreamsProcessing struct {
	Redis *redis.Client
	//other dependencies e.g. logger database goes here
}

// Process this method implements JobCallBack
///it will read and process each email and send it to our users
//the logic to send the emails goes here
func (r *RedisStreamsProcessing) Process(job interface{}) {
     //the go redis client returns the redis stream data as type [redis.XMessage]
	if data, ok := job.(redis.XMessage); ok {
		email := data.Values["email"].(string)
		message := data.Values["message"].(string)
		fmt.Printf("I am sending an email to the email  %vwith message:%v   \n ", email, message)
		//here we can decide to delete each entry when it is processed
		//in that case you can use the redis xdel command i.e:
		///r.Redis.XDel(context.Background(),streamName,data.ID).Err()
	} else {
		log.Println("wrong type of data sent")
	}
}
func StartProcessingEmails(rdb *redis.Client) {
	//create a new consumer instance to process the job
    //and pass it to our task queue
	redisStreams := RedisStreamsProcessing{
		Redis: rdb,
	}

	//in this case we have started 5 goroutines so at any moment we will
	//be sending a maximum of 5 emails.
	//you can adjust these parameters to increase or reduce
	q := taskq.NewQueue(5, 10, redisStreams.Process)

	//call startWorkers  in a different goroutine otherwise it will block
	go q.StartWorkers()

	//with our workers running now we can start listening to new events from redis stream
	//we start from id 0 i.e. the first item in the stream
	id := "0"
	for {
		var ctx = context.Background()
		data, err := rdb.XRead(ctx, &redis.XReadArgs{
			Streams: []string{streamName, id},

			//count is number of entries we want to read from redis
			Count: 4,

			//we use the block command to make sure if no entry is found we wait 
			//until an entry is found
			Block: 0,
		}).Result()
		if err != nil {
			log.Println(err)
			log.Fatal(err)
		}
		///we have received the data we should loop it and queue the messages
        //so that our jobs can start processing
		for _, result := range data {
			for _, message := range result.Messages {
				///we use EnqueueJobBlocking to send out jobs to the workers
				q.EnqueueJobBlocking(message)

				//here we set a new start id because we don't want to process old emails
				//so we have set the id to the last id we saw
				id = message.ID
			}
		}
	}
}