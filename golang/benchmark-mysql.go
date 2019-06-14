package main    // https://blog.huoding.com/2019/08/21/768

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var number, concurrency int

var cmd = &cobra.Command{
	Use:   "benchmark sql",
	Short: "a sql benchmark tool",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		b := benchmark{
			sql:         args[0],
			number:      number,
			concurrency: concurrency,
		}

		b.run()
	},
}

func init() {
	cobra.OnInitialize(config)

	cmd.Flags().IntVarP(&number, "number", "n", 100, "number")
	cmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "concurrency")
	cmd.Flags().SortFlags = false
}

func config() {
	viper.AddConfigPath(".")
	viper.SetConfigName("db")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	driver := viper.GetString("driver")
	dsn := viper.GetString("dsn")

	db, err = sql.Open(driver, dsn)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

type benchmark struct {
	sql         string
	number      int
	concurrency int
	duration    chan time.Duration
	start       time.Time
	end         time.Time
}

func (b *benchmark) run() {
	b.duration = make(chan time.Duration, b.number)
	b.start = time.Now()
	b.runWorkers()
	b.end = time.Now()

	b.report()
}

func (b *benchmark) runWorkers() {
	var wg sync.WaitGroup

	wg.Add(b.concurrency)

	for i := 0; i < b.concurrency; i++ {
		go func() {
			defer wg.Done()
			b.runWorker(b.number / b.concurrency)
		}()
	}

	wg.Wait()
	close(b.duration)
}

func (b *benchmark) runWorker(num int) {
	for i := 0; i < num; i++ {
		start := time.Now()
		b.request()
		end := time.Now()

		b.duration <- end.Sub(start)
	}
}

func (b *benchmark) request() {
	if _, err := db.Exec(b.sql); err != nil {
		log.Fatal(err)
	}
}

func (b *benchmark) report() {
	sum := 0.0
	num := float64(len(b.duration))

	for duration := range b.duration {
		sum += duration.Seconds()
	}

	qps := int(num / b.end.Sub(b.start).Seconds())
	tpq := sum / num * 1000

	fmt.Printf("qps: %d [#/sec]\n", qps)
	fmt.Printf("tpq: %.3f [ms]\n", tpq)
}