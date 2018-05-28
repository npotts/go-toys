package sample

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestSample_Contents(t *testing.T) {
	s := NewSample([]byte("abcdefghijlk"))
	s.AddTag("def")
	s.Values["sda"] = 43.232
	fmt.Println(s)
	fmt.Println(s.Contents())

	out, err := proto.Marshal(s)
	fmt.Printf("%v: [%d]%X\n", err, len(out), out)
}
