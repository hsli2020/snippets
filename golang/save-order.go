package amazon

import "fmt"

// ConvertToStdOrder/MakeStdOrder
type Order struct {
	OrderID string
	Items   []OrderItem
	ShipTo  Address
}

type OrderItem struct {
	OrderID string
	OrderUK string
	SKU     string
}

type Address struct {
	OrderID string
	Name    string
}

func main() {
	orders := []Order{
		Order{
			OrderID: "111",
			Items: []OrderItem{
				{SKU: "SKU-AAA"},
			},
			ShipTo: Address{Name: "Buyer-1"},
		},
		Order{
			OrderID: "222",
			Items: []OrderItem{
				{SKU: "SKU-AAA"},
				{SKU: "SKU-BBB"},
			},
			ShipTo: Address{Name: "Buyer-2"},
		},
		Order{
			OrderID: "333",
			Items: []OrderItem{
				{SKU: "SKU-AAA"},
				{SKU: "SKU-BBB"},
				{SKU: "SKU-CCC"},
			},
			ShipTo: Address{Name: "Buyer-3"},
		},
	}

	for _, order := range orders {
		SaveMasterOrder(order)

		orderID := order.OrderID

		if len(order.Items) > 1 {
			for i, item := range order.Items {
				uk := fmt.Sprintf("%s_%d", orderID, i+1)

				item.OrderID = orderID
				item.OrderUK = uk
				SaveOrderItem(item)
			}
		} else {
			item := order.Items[0]
			item.OrderID = orderID
			item.OrderUK = orderID
			SaveOrderItem(item)
		}

		shipTo := order.ShipTo
		shipTo.OrderID = orderID
		SaveShipTo(shipTo)

		fmt.Println()
	}
}

func SaveMasterOrder(order Order) {
	fmt.Println("Master Order:", order.OrderID)
}

func SaveOrderItem(item OrderItem) {
	fmt.Println("Order Item  :", item.OrderID, item.OrderUK, item.SKU)
}

func SaveShipTo(shipTo Address) {
	fmt.Println("Ship To     :", shipTo.OrderID, shipTo.Name)
}
