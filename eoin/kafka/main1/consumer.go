package main

import (
	"github.com/Shopify/sarama"
	"fmt"
)

func main() {
	//配置
	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_11_0_0
	//新建一个消费者
	consumer, e := sarama.NewConsumer([]string{"192.168.103.88:9092", "192.168.103.88:9092", "192.168.103.88:9092"}, config)
	if e != nil {
		panic("error get consumer")
	}
	defer consumer.Close()

	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err := consumer.ConsumePartition("logstash_test", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("error get partition consumer", err)
	}
	defer partitionConsumer.Close()
	//循环等待接受消息.
	for {
		select {
		//接收消息通道和错误通道的内容.
		case msg := <-partitionConsumer.Messages():
			fmt.Println("msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Println(err.Err)
		}
	}
}
/*
配置
新建一个消费者
通过消费者,指定主题的分区,获取一个特定的分区消费者.
通过分区消费者接收消息.
 */