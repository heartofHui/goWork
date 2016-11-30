package main

import (
	"batu/demo"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"time"
)

const (
	NetWorkAddr = "127.0.0.1:9000" # listen addr&port
)

type batuThrift struct{
}

func (this * batuThrift) CallBack(callTime int64,name string,paramMap map[string]string)(r []string,err error){
	fmt.Println("--->from client Call:",time.Unix(callTime,0).Format("2006-01-02 15:02:03"),name,paramMap)
	r = append(r, "key:"+paramMap["a"]+"    value:"+paramMap["b"])
}

func (this *batuThrift) Put(s *demo.Article)(err error){
	fmt.Println("Article--->id: %d\tTitle:%s\tContent:%t\tAuther:%d\n",s.Id,s.Title,s.Content,s.Auther)
	return nil	
}

func main(){
	
}