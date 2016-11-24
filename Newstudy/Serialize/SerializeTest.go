package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	//fmt.Println("hello")
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(128),
		int8(-54),
		uint8(255),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("%x", buf.Bytes())
}