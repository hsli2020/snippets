package main

import "fmt"

// suppliers.Client
type Client interface {
	Purchase(sku string) string
	//GetName() string
}

// dh.Client
type ClientA struct {
}

func (c ClientA) Purchase(sku string) string {
	return "Purcase " + sku + " from Client A"
}

// synnex.Client
type ClientB struct {
}

func (c ClientB) Purchase(sku string) string {
	return "Purcase " + sku + " from Client B"
}

func GetClient(name string) Client { // * not working
	var client Client // * not working

	switch (name) {
	case "A":
		client = ClientA{}
	case "B":
		client = ClientB{}
	}

	return client
}

func GetClientP(name string) Client {
	var client Client

	switch (name) {
	case "A":
		client = &ClientA{} // & works
	case "B":
		client = &ClientB{} // & works
	}

	return client
}

func main() {
	client := GetClient("A")
	result := client.Purchase("Keyboard")
	fmt.Println(result)

	client = GetClientP("B")
	result = client.Purchase("Monitor")
	fmt.Println(result)
}
