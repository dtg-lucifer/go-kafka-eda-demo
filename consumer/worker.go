package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type ConsumerDTO struct {
	Topic string
}

func StartWorker(dto ConsumerDTO) error {
	worker, err := ConnectCosumer([]string{"localhost:9092"})
	if err != nil {
		slog.Error("Cannot connect to the kafka broker as consumer", "err", err.Error())
		return err
	}

	consumer, err := worker.ConsumePartition(dto.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		slog.Error("Cannot consume partition", "err", err.Error())
		return err
	}

	slog.Info("Started consuming!!")

	sig_chan := make(chan os.Signal, 1)
	signal.Notify(sig_chan, syscall.SIGINT, syscall.SIGTERM)

	msg_count := 0

	done_chan := make(chan any)

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				slog.Error("Error consuming data!!", "err", err)
			case msg := <-consumer.Messages():
				msg_count++
				slog.Info("Received message", "topic", string(msg.Topic), "messages", msg_count)
				slog.Info("Message ---------->", "msg", msg)
				slog.Info("Value ---------->", "value", string(msg.Value))
			case <-sig_chan:
				slog.Error("Received interruption")
				done_chan <- "DONE"
			}
		}
	}()

	<-done_chan
	slog.Info("Processed messages", "count", msg_count)

	if err := worker.Close(); err != nil {
		slog.Error("Error closing the worker", "err", err)
		return err
	}

	return nil
}

func ConnectCosumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()

	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		slog.Error("Error connecting the consumer to the broker", "err", err)
		return nil, err
	}

	return conn, nil
}

func main() {
	StartWorker(ConsumerDTO{Topic: "comments"})
}
