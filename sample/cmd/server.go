//   $ ./pubsub client $url client0 & client0=$!
//   $ ./pubsub client $url client1 & client1=$!
//   $ ./pubsub client $url client2 & client2=$!
//   $ sleep 5
//   $ kill $server $client0 $client1 $client2
//
package main

import (
	"fmt"
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pull"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"os"

	sample ".."
	"github.com/golang/protobuf/proto"
)

type SampleServer struct {
	sock mangos.Socket
}

func NewSocket(url string) (*SampleServer, error) {
	sock, err := pull.NewSocket()
	if err != nil {
		return nil, err
	}
	s := &SampleServer{
		sock: sock,
	}
	s.sock.AddTransport(ipc.NewTransport())
	s.sock.AddTransport(tcp.NewTransport())

	return s, s.sock.Listen(url)
}

func (ss *SampleServer) Listen() {
	fmt.Println("Listen")
	for {
		msg, err := ss.sock.Recv()
		if err != nil {
			fmt.Println("recv err", err)
			continue
		}
		//unmarhall data
		smpl := &sample.Sample{}
		if err := proto.Unmarshal(msg, smpl); err != nil {
			fmt.Println("Unable to marshall", err)
		}
		fmt.Printf("Rx'd %d bytes from %s\n", len(msg), smpl.Tags[0])
		//fmt.Println(smpl.Contents())
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: %s <dial>")
		os.Exit(1)
	}
	s, e := NewSocket(os.Args[1])
	fmt.Println("Connected", e)
	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}
	s.Listen()
}
