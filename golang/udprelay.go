// main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var flagTimeout = flag.String("timeout", "10m", "duration to keep connections alive after their last packet")
var flagProtocol = flag.Bool("protocol", false, "enables the udprelay command protocol as defined in udprelay(7)")

var flagVersion = flag.Bool("version", false, "print the version to stdout and exit immediately")

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "udprelay %s\nusage: %s [OPTION...] port\n\noptions:\n", Version, os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "\nsee `man udprelay.1` for more information")
	}
	flag.Parse()

	if *flagVersion {
		fmt.Printf("udprelay %s\n", Version)
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	listenPort, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}

	log.Printf("udprelay %s\n\n", Version)

	timeoutDuration, err := time.ParseDuration(*flagTimeout)
	if err != nil {
		log.Println("error: parsing -timeout: %s\n", err)
		os.Exit(2)
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: listenPort})
	if err != nil {
		panic(err)
	}
	log.Printf("listen: %d\n", listenPort)

	relay := &Relay{
		Log:             log.New(os.Stderr, "", 0),
		Timeout:         timeoutDuration,
		CommandProtocol: *flagProtocol,
	}

	buf := make([]byte, 65536)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("error: %s\n", err)
			continue
		}
		packet := buf[:n]

		relay.HandlePacket(conn, addr, packet)
	}
}

// asciiutils.go
package main

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

func TrimLeftSpace(buf []byte) []byte {
	for i := 0; len(buf) > 0; i++ {
		if asciiSpace[buf[0]] != 1 {
			break
		}
		buf = buf[1:]
	}
	return buf
}

func TrimRightSpace(buf []byte) []byte {
	for i := 0; len(buf) > 0; i++ {
		if asciiSpace[buf[len(buf)-1]] != 1 {
			break
		}
		buf = buf[:len(buf)-1]
	}
	return buf
}

func TrimSpace(buf []byte) []byte {
	return TrimLeftSpace(TrimRightSpace(buf))
}

func Split2Space(buf []byte) ([]byte, []byte) {
	for i := 0; i < len(buf); i++ {
		if asciiSpace[buf[i]] == 1 {
			return buf[:i], buf[i+1:]
		}
	}
	return buf, []byte{}
}

// relay.go
package main

import (
	"bytes"
	"log"
	"net"
	"time"
)

var cmdPrefix = []byte("udprelay!")

type Relay struct {
	Log             *log.Logger
	CommandProtocol bool

	Timeout time.Duration

	peers    map[string]*Peer
	channels map[string]*Channel
}

type Channel struct {
	Peers map[string]*Peer
}

type Peer struct {
	Addr    *net.UDPAddr
	Timeout time.Time
	Channel string
}

func (relay *Relay) HandlePacket(conn *net.UDPConn, addr *net.UDPAddr, packet []byte) {
	peer := relay.peers[addr.String()]
	if peer == nil {
		peer = &Peer{
			Addr:    addr,
			Channel: "",
		}
		if relay.peers == nil {
			relay.peers = make(map[string]*Peer)
			relay.channels = make(map[string]*Channel)
			relay.channels[""] = &Channel{
				Peers: make(map[string]*Peer),
			}
		}
		channel := relay.channels[""]
		channel.Peers[peer.Addr.String()] = peer
		relay.peers[addr.String()] = peer
	}
	peer.Timeout = time.Now().Add(relay.Timeout)

	if relay.CommandProtocol {
		if bytes.HasPrefix(packet, cmdPrefix) {
			cmd := packet[len(cmdPrefix):]
			cmd, args := Split2Space(cmd)
			args = TrimSpace(args)
			relay.handleCommand(conn, packet, peer, string(cmd), args)
			return
		}
	}

	relay.sendPacket(conn, peer, peer.Channel, packet)
}

func (relay *Relay) sendPacket(conn *net.UDPConn, sender *Peer, channel string, packet []byte) {
	for _, peer := range relay.channels[channel].Peers {
		if sender != nil && peer.Addr == sender.Addr {
			continue
		}

		if time.Now().After(peer.Timeout) {
			relay.dropPeer(peer)
			continue
		}

		_, err := conn.WriteToUDP(packet, peer.Addr)
		if err != nil {
			relay.Log.Printf("error: writing to %s: %s\n", peer.Addr.String(), err)
			continue
		}
	}
}

func (relay *Relay) dropPeer(peer *Peer) {
	relay.switchChannel(peer, "")
	delete(relay.peers, peer.Addr.String())
}

func (relay *Relay) switchChannel(peer *Peer, channelName string) {
	delete(relay.channels[peer.Channel].Peers, peer.Addr.String())
	if !(peer.Channel == "") && len(relay.channels[peer.Channel].Peers) == 0 {
		delete(relay.channels, peer.Channel)
	}
	peer.Channel = channelName
	channel := relay.channels[channelName]
	if channel == nil {
		channel = &Channel{
			Peers: make(map[string]*Peer),
		}
		relay.channels[channelName] = channel
	}
	channel.Peers[peer.Addr.String()] = peer
}

func (relay *Relay) handleCommand(conn *net.UDPConn, packet []byte, peer *Peer, cmd string, args []byte) {
	switch cmd {
	case "echo":
		_, err := conn.WriteToUDP(packet, peer.Addr)
		if err != nil {
			relay.Log.Printf("error: replying to ping: %s\n", err)
		}
	case "channel":
		channelID := args
		var payload []byte
		hasPayload := false
		for i, ch := range args {
			if asciiSpace[ch] == 1 {
				channelID = args[:i]
				payload = args[i+1:]
				hasPayload = true
				break
			}
		}

		relay.switchChannel(peer, string(channelID))

		if hasPayload {
			relay.sendPacket(conn, peer, string(channelID), payload)
		}

		_, err := conn.WriteToUDP([]byte("udprelay!channel "+string(channelID)), peer.Addr)
		if err != nil {
			relay.Log.Printf("error: replying to channel switch message: %s\n", err)
		}
	}
}

// version.go
package main

var Version string = "v1.0.1"
