package service

import (
    "context"
    "github.com/plexmediamanager/service/log"
    "github.com/micro/go-micro"
    microClient "github.com/micro/go-micro/client"
    microServer "github.com/micro/go-micro/server"
)

type Registration struct {
    Topic               string
    Event               interface{}
    Options             func(options *microServer.SubscriberOptions)
}

// Register subscriptions server
func (application *Application) RegisterSubscriptionServer(subscriptionName string, processor interface{}, options func(options *microServer.SubscriberOptions)) {
    server := application.Service().Server()
    subscriptions := application.Configuration().Subscriptions()
    err := server.Subscribe(
        server.NewSubscriber(
            subscriptions.Get(subscriptionName).Topic,
            processor,
            options,
        ),
    )
    if err != nil {
        log.Panic(err)
    }
}

// Get default subscription options
func (application *Application) DefaultSubscriptionServerOptions() func(options *microServer.SubscriberOptions) {
    return func(options *microServer.SubscriberOptions) {
        options.Context = context.Background()
    }
}

// Register subscriptions client
func (application *Application) SendMessageToTopic(subscriptionName string, message interface{}, options func(options *microClient.PublishOptions), panicOnError bool) {
    client := application.Service().Client()
    subscriptions := application.Configuration().Subscriptions()

    publisher := micro.NewPublisher(subscriptions.Get(subscriptionName).Topic, client)
    err := publisher.Publish(context.TODO(), message, options)
    if err != nil {
        if panicOnError {
            log.Panic(err)
        } else {
            log.Print(err)
        }
    }
}

// Get default publisher options
func (application *Application) DefaultSubscriptionClientOptions() func(options *microClient.PublishOptions) {
    return func(options *microClient.PublishOptions) {}
}

// Register application
func (application *Application) RegisterApplication(topic string, message interface{}) {
    application.sendServiceDescriptorMessage(topic, message)
}

// Deregister Application
func (application *Application) DeregisterApplication(topic string, message interface{}) {
    application.sendServiceDescriptorMessage(topic, message)
}

// Send data to the manager
func (application *Application) RunServiceHealthCheck(topic string, message interface{}) {
    application.sendServiceDescriptorMessage(topic, message)
}

// Actually send message to the service manager
func (application *Application) sendServiceDescriptorMessage(topic string, message interface{}) {
    application.SendMessageToTopic(topic, message, application.DefaultSubscriptionClientOptions(), false)
}