package main
 ​
import (
   "context"
   "fmt"
   "runtime"
   "sync"
   "time"
)
 ​
func hardWork(job interface{}) error {
   time.Sleep(time.Second * 10)
   return nil
}
 ​
func requestWork(ctx context.Context, job interface{}) error {
   ctx, cancel := context.WithTimeout(ctx, time.Second*2)
   defer cancel()
 ​
   done := make(chan error, 1)
   panicChan := make(chan interface{}, 1)
   go func() {
     defer func() {
       if p := recover(); p != nil {
         panicChan <- p
       }
     }()
 ​
     done <- hardWork(job)
   }()
 ​
   select {
   case err := <-done:
     return err
   case p := <-panicChan:
     panic(p)
   case <-ctx.Done():
     return ctx.Err()
   }
}
 ​
func main() {
   const total = 10
   var wg sync.WaitGroup
   wg.Add(total)
   now := time.Now()
 ​
   for i := 0; i <total; i++ {
     go func() {
       defer func() {
         if p := recover(); p != nil {
           fmt.Println("oops, panic")
         }
       }()
 ​
       defer wg.Done()
       requestWork(context.Background(), "any")
     }()
   }
 ​
   wg.Wait()
   fmt.Println("elapsed:", time.Since(now))
   time.Sleep(time.Second * 20)
   fmt.Println("number of goroutines:", runtime.NumGoroutine())
}