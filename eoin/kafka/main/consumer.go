package main

import (
	"sync"
	"log"
	"os"
	"github.com/Shopify/sarama"
	"strings"
	"fmt"
)

var (
	wg     sync.WaitGroup
	logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
)

func main() {
	sarama.Logger = logger
	consummer, err := sarama.NewConsumer(strings.Split("192.168.100.74:9092", ","), nil)
	if err != nil {
		logger.Println("Failed to start consumer:%s", err)
	}
	partitionList, err := consummer.Partitions("hello")
	if err != nil {
		logger.Println("Failed to get the list of partitons: ", err)
	}
	for partition := range partitionList {
		pc, err := consummer.ConsumePartition("hello", int32(partition), sarama.OffsetNewest)
		if err != nil {
			logger.Printf("Failed to start consumer for partiton %d:%s\n", partition, err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Paetiton:%d,Offset:%d,Key:%s,Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
		}(pc)
	}
	wg.Wait()
	logger.Println("Done consuming topic hello")
	consummer.Close()
}
