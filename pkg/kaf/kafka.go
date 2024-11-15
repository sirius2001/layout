package kaf

import (
	"encoding/json"
	"github.com/sirius2001/loon/config"
	"github.com/sirius2001/loon/pkg/grpc/pb"
	"github.com/sirius2001/loon/pkg/log"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

func NewProducer() error {
	var err error
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	producer, err = sarama.NewSyncProducer(config.Conf().Kafka.Addr, kafkaConfig)
	if err != nil {
		return err
	}
	log.Info("connet to kafka successfully", "addr", config.Conf().Kafka.Addr, "topic", config.Conf().Kafka.Topic)
	return nil
}

type message struct {
	*pb.AuditRecord
	Result string `json:"result"`
}

func Message(record *pb.AuditRecord) {
	m := &message{
		AuditRecord: record,
		Result:      string(record.Result),
	}

	data, err := json.Marshal(m)
	if err != nil {
		log.Error("Message Marshal", "err", err)
		return
	}

	// 创建要发送的消息
	msg := &sarama.ProducerMessage{
		Topic: config.Conf().Kafka.Topic, // 目标 Topic
		Value: sarama.ByteEncoder(data),  // 消息内容
	}

	// 发送消息
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Error("kafka producer sendMessage", "err", err)
		return
	}

	log.Info("new record send to kafka suceesfully")
}
