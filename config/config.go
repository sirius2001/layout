package config

import (
	"encoding/json"
	"log/slog"
	"os"
)

type Config struct {
	HTTP struct {
		HTTPAddress string `json:"http_address"`
		HTTPEnable  bool   `json:"http_enable"`
		HTTPRouter  string `json:"http_router"`
	} `json:"http"`

	GRPC struct {
		RPCEnable   bool   `json:"rpc_enable"`
		GRPCAddress string `json:"rpc_address"`
	} `json:"grpc"`

	Kafka struct {
		KafkaAddr  []string `json:"kafka_addr"` // 支持kafka集群 多kafka
		KafkaTopic string   `json:"kafka_topic"`
	} `json:"kafka"`
}

var config = &Config{}

func LoadConfig(confPath string) (*Config, error) {
	file, err := os.Open(confPath)
	if err != nil {
		slog.Error("LoadConfig", "Err", err)
		return nil, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		slog.Error("LoadConfig", "Err", err)
		return nil, err
	}

	return config, nil
}

func GetConfig() *Config {
	return config
}
