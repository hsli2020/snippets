package main // https://callistaenterprise.se/blogg/teknik/2019/10/05/go-worker-cancellation/

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const workerPoolSize = 4

func main() {
	// create the consumer
	consumer := Consumer{
		ingestChan: make(chan int, 1),
		jobsChan:   make(chan int, workerPoolSize),
	}

	// Simulate external lib sending us 10 events per second
	producer := Producer{callbackFunc: consumer.callbackFunc}
	go producer.start()

	// Set up cancellation context and waitgroup
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Start consumer with cancellation context passed
	go consumer.startConsumer(ctx)

	// Start workers and Add [workerPoolSize] to WaitGroup
	wg.Add(workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		go consumer.workerFunc(wg, i)
	}

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	// Handle shutdown
	fmt.Println("***********************\nShutdown signal received\n***********************")
	cancelFunc() // Signal cancellation to context.Context
	wg.Wait()    // Block here until are workers are done

	fmt.Println("All workers done, shutting down!")
}

// The consumer struct:

// -- Consumer below here!
type Consumer struct {
	ingestChan chan int
	jobsChan   chan int
}

// callbackFunc is invoked each time the external lib passes an event to us.
func (c Consumer) callbackFunc(event int) {
	c.ingestChan <- event
}

// workerFunc starts a single worker function that will range on the jobsChan until that
// channel closes.
func (c Consumer) workerFunc(wg *sync.WaitGroup, index int) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", index)
	for eventIndex := range c.jobsChan {
		// simulate work  taking between 1-3 seconds
		fmt.Printf("Worker %d started job %d\n", index, eventIndex)
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d finished processing job %d\n", index, eventIndex)
	}
	fmt.Printf("Worker %d interrupted\n", index)
}

// startConsumer acts as the proxy between the ingestChan and jobsChan, with a select to
// support graceful shutdown.
func (c Consumer) startConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.ingestChan:
			c.jobsChan <- job
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.jobsChan)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}

// the Producer struct that simulates our external library:

// -- Producer simulates an external library that invokes the
// registered callback when it has new data for us once per 100ms.
type Producer struct {
	callbackFunc func(event int)
}

func (p Producer) start() {
	eventIndex := 0
	for {
		p.callbackFunc(eventIndex)
		eventIndex++
		time.Sleep(time.Millisecond * 100)
	}
}
