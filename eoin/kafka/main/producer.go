package main

import (
	"log"
	"os"
	"github.com/Shopify/sarama"
	"strings"
)

var logger1=log.New(os.Stderr,"[srama]",log.LstdFlags)

func main() {
	sarama.Logger=logger1
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	msg:=&sarama.ProducerMessage{}
	msg.Topic="hello"
	msg.Partition=int32(-1)
	msg.Key=sarama.StringEncoder("key")
	msg.Value=sarama.ByteEncoder("世界你好")
	producer,err:=sarama.NewSyncProducer(strings.Split("192.168.100.74:9092",","),nil)
	if err!=nil{
		logger1.Println("Failed to produce message:%s",err)
		os.Exit(500)
	}
	defer producer.Close()
	partiton,offset,err:=producer.SendMessage(msg)
	if err!=nil{
		logger1.Println("Failed to produce massage:",err)
	}
	logger1.Printf("parttiton=%d,offset=%d\n",partiton,offset)

}