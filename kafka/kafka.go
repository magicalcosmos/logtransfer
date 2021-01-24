package kafka

import (
	"fmt"

	"logtransfer/es"

	"github.com/Shopify/sarama"
)

// Init init
func Init(address []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("failed to start consumer, err: %v\n", err)
		return
	}
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("failed to start consumer partition list, err: %v\n", err)
		return
	}
	for partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer partition %v, err: %v\n", partition, err)
			return err
		}
		// defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				ld := es.LogData{
					Topic: topic,
					Data:  string(msg.Value),
				}
				es.SendToESChan(&ld)
			}
		}(pc)
	}
	return
}
