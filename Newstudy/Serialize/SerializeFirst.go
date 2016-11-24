package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash/crc32"
	//"io"
	//"net"
	"os"
)

type Resister struct {
	ACTION int32
	SID    int32
}

type Packet struct {
	length uint32
	crc32  uint32
	info   string
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func (p Packet) Encode() []byte {
	bufTmp := new(bytes.Buffer)
	var length int = len([]byte(p.info))
	err := binary.Write(bufTmp, binary.LittleEndian, (int16)(length))
	checkError(err)

	err = binary.Write(bufTmp, binary.LittleEndian, []byte(p.info))
	checkError(err)

	buf := new(bytes.Buffer)
	p.length = uint32(bufTmp.Len() + 8)

	err = binary.Write(buf, binary.LittleEndian, p.length)
	checkError(err)

	p.crc32 = crc32.ChecksumIEEE(buf.Bytes())
	err = binary.Write(buf, binary.LittleEndian, p.crc32)
	checkError(err)

	err = binary.Write(buf, binary.LittleEndian, bufTmp.Bytes())
	checkError(err)
	return buf.Bytes()
}

func (p Packet) Decode(buff []byte) {
	buf := bytes.NewBuffer(buff)
	err := binary.Read(buf, binary.LittleEndian, &(p.length))
	checkError(err)

	fmt.Printf("length:%d\n", p.length)

	err = binary.Read(buf, binary.LittleEndian, &(p.crc32))
	checkError(err)
	fmt.Printf("crc32:%d\n", p.crc32)

	bufTmp := bytes.NewBuffer(buff[8:])
	crc := crc32.ChecksumIEEE(bufTmp.Bytes())

	if crc != p.crc32 {
		fmt.Errorf("crc not equal!")
	}

	var info_len int16
	binary.Read(buf, binary.LittleEndian, &(info_len))

	fmt.Printf("info_len:%d\n", info_len)

	bufStr := bytes.NewBuffer(buff[10:])

	p.info = (string)(bufStr.Bytes())
	fmt.Printf("info:%s\n", p.info)
}

func main() {
	m := Resister{20004, 1024}
	b, err := json.Marshal(m)
	checkError(err)
	var packet Packet
	packet.info = string(b)

	buf := packet.Encode()
	fmt.Println(len(buf))

	var msg Packet
	msg.Decode(buf)
	os.Exit(0)
}
