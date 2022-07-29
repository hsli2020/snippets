package orderbook

import (
	"math"
	"sort"
	"untitled/data"
)

const (
	MaxFactor = 1e8
)

// Order represents an order in the books. Ask or Bid.
type Order struct {
	Price float64
	Size  float64
}

// Orderbook represents the book for a specific ticker, holding
// both the asks and bids.
type Orderbook struct {
	Asks Orders
	Bids Orders
}

// NewFromFile loads a new Orderbook form the given ask and bid .gob encoded file.
// it will sort the orders by best bid or ask.
func NewFromFile(asksSrc, bidsSrc string) *Orderbook {
	return &Orderbook{
		Asks: loadAndSortOrders(asksSrc, false),
		Bids: loadAndSortOrders(bidsSrc, true),
	}
}

// Aggregate aggregates all the orders for the given orderbook depth.
// After aggregation all orders, both the asks and bids will be sorted by their best price.
// NOTE: Using goroutines here will cause computational overhead and
// will not increase performance. However, this might be different if we need to aggregate much
// a larger amount of orders.
func (b *Orderbook) Aggregate(depth int) {
	b.Asks = aggregate(b.Asks, depth)
	sort.Sort(ByBestAsk{b.Asks})

	b.Bids = aggregate(b.Bids, depth)
	sort.Sort(ByBestBid{b.Bids})
}

func aggregate(orders []*Order, depth int) []*Order {
	aggregated := make(map[int]*Order)
	for i := 0; i < len(orders); i++ {
		pInc := roundToDepth(float64(depth), orders[i].Price)
		order, ok := aggregated[pInc]
		if !ok {
			aggregated[pInc] = orders[i]
		} else {
			order.Size = Round(order.Size + orders[i].Size)
		}
	}

	aggOrders := make([]*Order, len(aggregated))
	i := 0
	for price, order := range aggregated {
		order.Price = float64(price)
		aggOrders[i] = order
		i++
	}
	return aggOrders
}

// Orders is a slice of Order. For sorting purposes.
type Orders []*Order

// ByBestAsk is used to sort orders "by best ask"
type ByBestAsk struct{ Orders }

func (a ByBestAsk) Len() int           { return len(a.Orders) }
func (a ByBestAsk) Swap(i, j int)      { a.Orders[i], a.Orders[j] = a.Orders[j], a.Orders[i] }
func (a ByBestAsk) Less(i, j int) bool { return a.Orders[i].Price < a.Orders[j].Price }

// ByBestBid is used to sort orders "by best bid"
type ByBestBid struct{ Orders }

func (a ByBestBid) Len() int           { return len(a.Orders) }
func (a ByBestBid) Swap(i, j int)      { a.Orders[i], a.Orders[j] = a.Orders[j], a.Orders[i] }
func (a ByBestBid) Less(i, j int) bool { return a.Orders[i].Price > a.Orders[j].Price }

func loadAndSortOrders(src string, bids bool) Orders {
	var (
		orders     = data.LoadOrdersFromFile(src)
		orderSlice = make([]*Order, len(orders))
		i          = 0
	)
	for price, size := range orders {
		orderSlice[i] = &Order{
			Price: price,
			Size:  Round(size),
		}
		i++
	}

	if !bids {
		sort.Sort(ByBestAsk{orderSlice})
	} else {
		sort.Sort(ByBestBid{orderSlice})
	}

	return orderSlice
}

func roundToDepth(weight float64, val float64) int {
	return int(math.Ceil(val/weight) * weight)
}

// Round rounds a float64 to the provided precision factor.
func Round(n float64) float64 {
	return math.Round(n*MaxFactor) / MaxFactor
}
