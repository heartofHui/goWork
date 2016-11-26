package main

import (
	"fmt"
	"log"
	"github.com/golang/protobuf/proto"
	"example"
)

func main() {
	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(1024),
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("Marshaling error:", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
	}
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(),
			newTest.GetLabel())
	}
	fmt.Printf("info = %s\n",newTest.String())
}