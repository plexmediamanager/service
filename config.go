package service

import (
    "bytes"
    "encoding/json"
    format "fmt"
    "github.com/plexmediamanager/service/helpers"
    consulapi "github.com/hashicorp/consul/api"
    "github.com/spf13/viper"
    "io/ioutil"
    "strings"
    "time"
)

type subscription struct {
    Topic           string
    Queue           string
}

type subscriptions struct {
    list            map[string]*subscription
}

type config struct {
    Broker              BrokerInterface
    DurableName         string
    MetaInformation     map[string]string
    Topic               string
    Queue               string
    Register            struct {
        Interval        time.Duration
        TimeToLive      time.Duration
    }

    subscriptions       *subscriptions
}

func (config *config) init(application *Application) error {
    viper.SetEnvPrefix(application.EnvironmentPrefix())
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    viper.AutomaticEnv()
    config.setDefaultValuesForBroker(application)

    consulTokenFile := viper.GetString("consul.token.file")
    if consulTokenFile != "" {
        data, err := ioutil.ReadFile(format.Sprintf(consulTokenFile, application.EnvironmentPrefix))
        if err == nil {
            viper.SetDefault("consul.token", strings.TrimSpace(string(data)))
        }
    }

    apiOptions := consulapi.DefaultConfig()
    apiOptions.Address = viper.GetString("consul.endpoint")
    apiOptions.Token = viper.GetString("consul.token")

    if apiOptions.Address == "" {
        apiOptions.Address = helpers.GetEnvironmentVariableAsString("CONSUL_ENDPOINT", "")
    }

    if apiOptions.Token == "" {
        apiOptions.Token = helpers.GetEnvironmentVariableAsString("CONSUL_TOKEN", "")
    }

    consulClient, err := consulapi.NewClient(apiOptions)
    if err != nil {
        return UnableToConnectToConsul.ToError(err)
    }

    brokerKey := format.Sprintf("%s.broker-%s", application.Vendor(), application.Environment())
    keyValueStorage, _, err := consulClient.KV().Get(brokerKey, nil)
    if err != nil {
        return UnableToGetKeyFromConsulKV.ToErrorWithArguments(err, brokerKey)
    }

    if keyValueStorage == nil {
        return BrokerKeyNotFound.ToError(nil)
    } else {
        brokerConfiguration := brokerConfiguration{}
        bytesBuffer := bytes.NewBuffer(keyValueStorage.Value)
        err = json.NewDecoder(bytesBuffer).Decode(&brokerConfiguration)
        if err != nil {
            return InvalidBrokerConfiguration.ToError(err)
        }


        switch brokerConfiguration.Type {
            case "nats":
                config.Broker = &NatsBroker{}
            default:
                return BrokerNotSupported.ToErrorWithArguments(nil, brokerConfiguration.Type)
        }
        err = config.Broker.Decode(brokerConfiguration.Config)
        if err != nil {
            return InvalidBrokerConfiguration.ToError(err)
        }
    }

    key := format.Sprintf("%s-%s", application.ConsulConfigurationKey(), application.Environment())
    keyValueStorage, _, err = consulClient.KV().Get(key, nil)
    if err != nil || keyValueStorage == nil {
        return UnableToGetServiceConfigurationFromConsulKV.ToErrorWithArguments(err, key)
    } else {
        viper.SetConfigType("json")
        err = viper.MergeConfig(bytes.NewBuffer(keyValueStorage.Value))
        if err != nil {
            return ViperUnableToMergeConfigs.ToError(err)
        }
    }

    subscriptionsKey := format.Sprintf("%s.subscriptions", application.Vendor())
    keyValueStorage, _, err = consulClient.KV().Get(subscriptionsKey, nil)
    if err != nil {
        return UnableToGetKeyFromConsulKV.ToErrorWithArguments(err, key)
    }
    if keyValueStorage == nil {
        return UnableToGetSubscriptionsInformationForVendorFromConsul.ToErrorWithArguments(err, subscriptionsKey)
    } else {
        subscriptions := &subscriptions{}
        err = subscriptions.decode(keyValueStorage.Value)
        if err != nil {
            return UnableToDecodeSubscriptionsConfiguration.ToError(err)
        }
        config.subscriptions = subscriptions
    }
    return config.loadConfiguration()
}

func (config *config) loadConfiguration() error {
    config.Topic                    =   viper.GetString("broker.topic")
    config.Queue                    =   viper.GetString("broker.queue")
    config.DurableName              =   viper.GetString("broker.durable_name")
    config.Register.Interval        =   viper.GetDuration("register.interval")
    config.Register.TimeToLive      =   viper.GetDuration("register.ttl")
    config.MetaInformation          =   viper.GetStringMapString("meta")
    return nil
}

func (config *config) setDefaultValuesForBroker(application *Application) {
    viper.SetDefault("broker.topic", application.Name())
    viper.SetDefault("broker.queue", application.Name())
    viper.SetDefault("broker.durable_name", application.Name())
    viper.SetDefault("register.interval", time.Second * 5)
    viper.SetDefault("register.ttl", time.Second * 15)
}

func (config *config) Subscriptions() *subscriptions {
    return config.subscriptions
}

func (subscriptions *subscriptions) decode(bytesArray []byte) error {
    subscriptions.list = make(map[string]*subscription)
    err := json.Unmarshal(bytesArray, &subscriptions.list)
    return err
}

func (subscriptions *subscriptions) Get(name string) *subscription {
    return subscriptions.list[name]
}

func (subscriptions *subscriptions) Range() map[string]*subscription {
    return subscriptions.list
}
