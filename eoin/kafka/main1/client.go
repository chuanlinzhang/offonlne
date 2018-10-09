package main

import (
	"github.com/Shopify/sarama"
	"fmt"
)

func main() {
	config:=sarama.NewConfig()
	config.Version=sarama.V0_10_0_0
	client,err:=sarama.NewClient([]string{"192.168.103.88:9092","192.168.103.88:9092","192.168.103.88:9092"},config)
	if err!=nil{
		panic("client create error")
	}
	defer client.Close()
	//获取主题的名称集合
	topics,err:=client.Topics()
	if err!=nil{
		panic("get topics err")
	}
	for _,e:=range topics {
		fmt.Println(e)
	}
	//获取broker集合
	brokers:=client.Brokers()
	//输出每个机器的地址
	for _,broker:=range brokers {
		fmt.Println(broker.Addr())
	}
}
