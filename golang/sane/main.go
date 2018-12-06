package main

import (
	"fmt"
	"github.com/bloom42/sane-go"
	"io/ioutil"
)

type Creator struct {
	Name    string `sane:"name"`
	Website string `sane:"website"`
}

type Database struct {
	Server         string `sane:"server"`
	Ports          []int  `sane:"ports"`
	ConnectionMax  int    `sane:"connection_max"`
	Enabled        bool   `sane:"enabled"`
}

type Alpha struct {
	IP string `sane:"ip"`
	DC string `sane:"dc"`
}

type Beta struct {
	IP string `sane:"ip"`
	DC string `sane:"dc"`
}

type Servers struct {
	Alpha Alpha `sane:"alpha"`
	Beta  Beta  `sane:"beta"`
}

type EmptyMap struct {
}

type Config struct {
	Title     string   `sane:"title"`
	Creator   Creator  `sane:"creator"`
	Database  Database `sane:"database"`
	Servers   Servers  `sane:"servers"`
	EmptyMap  EmptyMap `sane:"empty_map"`
	Hosts     []string `sane:"hosts"`
}

func main() {
    var err error
	var cfg Config

	err = sane.Load("config.sane", &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n\n", cfg)

	data, err := ioutil.ReadFile("config.sane")
	err = sane.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n\n", cfg)

	b, err := sane.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
