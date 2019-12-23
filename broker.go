package service

import (
    "encoding/json"
    "errors"
    format "fmt"
    "github.com/nats-io/stan.go"
)

type BrokerInterface interface {
    Name()              string
    Decode([]byte)      error
}

type brokerConfiguration struct {
    Type                string
    Config              json.RawMessage
}

type NatsBroker struct {
    Endpoint            string                  `json:"endpoint"`
    ClusterID           string                  `json:"cluster_id"`
    Timeout             struct {
        Connect         Duration       `json:"connect"`
        Ack             Duration       `json:"ack"`
    }
    DiscoverPrefix      string                  `json:"discover_prefix"`
    MaxPubAcks          int                     `json:"max_pub_acks"`
    Ping                struct {
        Interval        int                     `json:"interval"`
        MaxOut          int                     `json:"max_out"`
    }
}

func (broker *NatsBroker) Name() string {
    return "nats"
}

func (broker *NatsBroker) Decode(decode []byte) error {
    err := json.Unmarshal(decode, broker)
    if err != nil {
        return errors.New(format.Sprintf("[NatsBroker::Decode] %v", err))
    }

    if broker.Timeout.Connect.Duration == 0 {
        broker.Timeout.Connect.Duration = stan.DefaultConnectWait
    }
    if broker.Timeout.Ack.Duration == 0 {
        broker.Timeout.Ack.Duration = stan.DefaultAckWait
    }
    if broker.DiscoverPrefix == "" {
        broker.DiscoverPrefix = stan.DefaultDiscoverPrefix
    }
    if broker.MaxPubAcks == 0 {
        broker.MaxPubAcks = stan.DefaultMaxPubAcksInflight
    }
    if broker.Ping.Interval == 0 {
        broker.Ping.Interval = stan.DefaultPingInterval
    }
    if broker.Ping.MaxOut == 0 {
        broker.Ping.MaxOut = stan.DefaultPingMaxOut
    }
    return nil
}
