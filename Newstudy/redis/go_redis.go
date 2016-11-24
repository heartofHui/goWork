package main

import (
	"fmt"
	//"log"
	"os"
	"github.com/alphazero/Go-Redis"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func main() {
	spec := redis.DefaultSpec().Host("127.0.0.1").Port(6379).Db(0).Password("")
	client, err := redis.NewSynchClientWithSpec(spec)
	checkError(err)

	dbkey := "Game:TEST:info"

	value := []byte("hello world")
	client.Set(dbkey, value)

	value, err = client.Get(dbkey)
	checkError(err)

	fmt.Printf("Get:%s\n", value)
}