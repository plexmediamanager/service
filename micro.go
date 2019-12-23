package service

import (
    "github.com/plexmediamanager/service/helpers"
    "github.com/hashicorp/consul/api"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/broker"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/client/selector"
    "github.com/micro/go-micro/client/selector/registry"
    nats "github.com/micro/go-plugins/broker/stan"
    "github.com/micro/go-plugins/registry/consul"
    "github.com/micro/go-plugins/wrapper/select/roundrobin"
    "github.com/nats-io/stan.go"
    "time"
)

type Micro struct {
    Service         micro.Service
}

func (handler *Micro) Initialize(application *Application) error {
    if application.config.Broker == nil {
        return BrokerNotConfigured.ToError(nil)
    }

    var microBroker broker.Broker
    switch application.config.Broker.Name() {
    case "nats":
        b := application.config.Broker.(*NatsBroker)
        microBroker = nats.NewBroker(
            nats.Options(
                stan.Options{
                    NatsURL:            b.Endpoint,
                    ConnectTimeout:     b.Timeout.Connect.Duration,
                    AckTimeout:         b.Timeout.Ack.Duration,
                    DiscoverPrefix:     b.DiscoverPrefix,
                    MaxPubAcksInflight: b.MaxPubAcks,
                    PingInterval:       b.Ping.Interval,
                    PingMaxOut:         b.Ping.MaxOut,
                },
            ),
            nats.DurableName(application.config.DurableName),
            nats.ClusterID(b.ClusterID),
        )
    default:
        return BrokerNotSupported.ToErrorWithArguments(nil, application.config.Broker.Name())
    }

    apiConfiguration := &api.Config {
        Address:    helpers.GetEnvironmentVariableAsString("CONSUL_ENDPOINT", ""),
        Token:      helpers.GetEnvironmentVariableAsString("CONSUL_TOKEN", ""),
    }

    consulRegistry := consul.NewRegistry(
        consul.Config(apiConfiguration),
    )
    registrySelector := registry.NewSelector(
        registry.TTL(time.Second),
        selector.Registry(consulRegistry),
    )
    wrapper := roundrobin.NewClientWrapper()

    service := micro.NewService(
        micro.Context(application.ctx),
        micro.Name(application.ServiceName()),
        micro.Version(application.Version()),
        micro.Metadata(application.config.MetaInformation),
        micro.Registry(consulRegistry),
        micro.Selector(registrySelector),
        func(options *micro.Options) {
            options.Client.Init(client.Selector(registrySelector))
        },
        micro.WrapClient(wrapper),
        micro.RegisterInterval(application.config.Register.Interval),
        micro.RegisterTTL(application.config.Register.TimeToLive),
        micro.Broker(microBroker),
        micro.AfterStart(func() error {
            return nil
        }),
    )

    if err := microBroker.Init(); err != nil {
        return BrokerInitializationError.ToError(err)
    }
    if err := microBroker.Connect(); err != nil {
        return BrokerConnectionError.ToError(err)
    }

    handler.Service = service
    return nil
}
