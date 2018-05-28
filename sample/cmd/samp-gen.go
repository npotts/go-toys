//   $ ./pubsub client $url client0 & client0=$!
//   $ ./pubsub client $url client1 & client1=$!
//   $ ./pubsub client $url client2 & client2=$!
//   $ sleep 5
//   $ kill $server $client0 $client1 $client2
//
package main

import (
	sample ".."
	"fmt"
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/push"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"github.com/golang/protobuf/proto"
	"os"
	"time"
)

type SampleClient struct {
	sock mangos.Socket
	name string
}

func NewSocket(name, url string) (*SampleClient, error) {
	sock, err := push.NewSocket()
	if err != nil {
		return nil, err
	}
	s := &SampleClient{
		sock: sock,
		name: name,
	}
	s.sock.AddTransport(ipc.NewTransport())
	s.sock.AddTransport(tcp.NewTransport())
	return s, s.sock.Dial(url)
}

func (ss *SampleClient) Generate() {
	for {
		d := sample.New(nil)
		d.AddTags(ss.name, "test", "testing", "mangos", "pb")
		d.SampleTime.Seconds -= 40
		/*	d.Values = map[string]float64{
				"secs": float64(d.SampleTime.Seconds),
				"ns":   float64(d.SampleTime.Nanos),
			}
		*/
		payload, err := proto.Marshal(d)
		if err != nil {
			panic(err)
		}
		err = ss.sock.Send(payload)
		if err != nil {
			fmt.Println("Error: unable to send: ", err)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: %s <name> <dial>")
		os.Exit(1)
	}
	s, e := NewSocket(os.Args[1], os.Args[2])
	fmt.Println("Dial succeeded:", e)
	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}
	s.Generate()
}
