package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"time"
)

func main() {
  go producer()
  go consumer()
 time.Sleep(time.Second*20)

}
/*
首先新建一个config,用于配置生产者相关的配置项
通过config和一个包含一个或多个kafka服务器的字符串数组,新建一个producer
定义一个 生产信息 ,包括发送的主题,哪个分区,重试次数等等信息和消息内容.
通过producer的输入通道,接受msg
如果配置中配置了,接收服务器反馈的响应,可以通过Successes和Errors通道来接受成功或失败的内容
 */
func producer()  {
	//设置配置
	config:=sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks=sarama.WaitForAll
	//随机的分区类型
	config.Producer.Partitioner=sarama.NewRoundRobinPartitioner
	//是否等待成功和失败的响应，只有上面的RequireAcks设置不是NoReponse才有用
	config.Producer.Return.Successes=true
	config.Producer.Return.Errors=true
	//设置使用kafka版本，如果低于v0_0_0版本，消息中的timestrap没有作用，需要消费和生产同时配置
	//config.Version=sarama.V0_11_0_0

	//使用配置，新建一个异步生产者
	producer,e:=sarama.NewAsyncProducer([]string{"192.168.103.88:9092"},config)
	if e!=nil{
		panic(e)
	}
	defer producer.AsyncClose()
	//发送的消息，主题，key
	msg:=&sarama.ProducerMessage{
		Topic:"logstash_test",
		Key:sarama.StringEncoder("test"),

	}
	var value string
	for {
		value="this is a message"
		//设置发送的真正内容
		fmt.Scanln(&value)
		//将字符串转化为字节数组
		msg.Value=sarama.ByteEncoder(value)
		fmt.Println(value,"___")
		//使用通道发送
		producer.Input()<-msg
		//循环判断哪个通道发送数据
		select {
		case suc:=<-producer.Successes():
			fmt.Println("offset:",suc.Offset,"timestmp:",suc.Timestamp.String(),"partiton:",suc.Partition)

		}
	}
}
func consumer()  {
	//配置
	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_11_0_0
	//新建一个消费者
	consumer, e := sarama.NewConsumer([]string{"192.168.103.88:9092"}, config)
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