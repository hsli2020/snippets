
start := time.Now()
elapsed := time.Since(start)
TPS := int(float64(msgCount) / elapsed.Seconds())

// event := new(OrderEvent)
// var event OrderEvent

// Trap SIGINT to trigger a shutdown.
signals := make(chan os.Signal, 1)
signal.Notify(signals, os.Interrupt)

ConsumerLoop:
for {
    select {
		case msg := <-ch.Messages():
            // do something
		case <-signals:
			break ConsumerLoop
		}
	}
}

import log "github.com/sirupsen/logrus"

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
		FullTimestamp:   true,
	})

	log.SetLevel(log.InfoLevel)
    // log.SetLevel(log.DebugLevel)
}

func (engine *Engine) Match(order *Order) {
	switch order.Side {
	case "SELL":
		matchingOrder, matches := engine.matchAskOrder(order)
		if matchingOrder.Amount > 0 {
			addAskOrder(engine.askbook, order)
		}

	case "BUY":
		matchingOrder, matches := engine.matchBidOrder(order)
		if matchingOrder.Amount > 0 {
			addBidOrder(engine.bidbook, order)
		}
	}
}

