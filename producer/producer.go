package producer

import (
	"log/slog"

	"github.com/IBM/sarama"
)

type ProducerDTO struct {
	Topic string `json:"topic"`
	Data  []byte `json:"data"`
}

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	conn, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PushToQueue(dto ProducerDTO) error {
	brokers := []string{"localhost:9092"}

	// @INFO Open a connection to the Kafka cluster
	conn, err := ConnectProducer(brokers)
	if err != nil {
		return err
	}

	// @INFO Close the connection when the function exits
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()

	partition, offset, err := conn.SendMessage(&sarama.ProducerMessage{
		Topic: dto.Topic,
		Value: sarama.ByteEncoder(dto.Data),
	})
	if err != nil {
		return err
	}

	slog.Info("Sent message to", "topic", dto.Topic, "partition", partition, "offset", offset)

	return nil
}
