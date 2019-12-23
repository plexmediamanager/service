package service

import (
    "context"
    "crypto/rand"
    format "fmt"
    "github.com/plexmediamanager/service/ctx"
    "github.com/plexmediamanager/service/helpers"
    "github.com/plexmediamanager/service/log"
    "github.com/google/uuid"
    "github.com/micro/go-micro"
    "github.com/sirupsen/logrus"
    "os"
    "os/signal"
    "runtime"
    "strconv"
    "strings"
    "syscall"
    "time"
)

type ApplicationKey struct {}

type Application struct {
    uuid                uuid.UUID
    vendor              string
    name                string
    version             VersionDescriptor
    environment         ApplicationEnvironment
    originalVendor      string
    originalName        string

    config              *config
    micro               *Micro
    ctx                 context.Context
    cancel              context.CancelFunc
    beforeStop          []func() error
}

// Create new application
func CreateApplication() *Application {
    currentEnvironment := GetApplicationEnvironment()
    preloadEnvironment(currentEnvironment)

    application := &Application {
        vendor:         normalizeVendorName(getApplicationVendorFromEnvironment()),
        name:           normalizeApplicationName(getApplicationNameFromEnvironment()),
        environment:    currentEnvironment,
        uuid:           createApplicationUUID(),
        originalVendor: getApplicationVendorFromEnvironment(),
        originalName:   getApplicationNameFromEnvironment(),
    }
    application.version.Initialize()
    application.PrettyVersionCLI()

    application.ctx, application.cancel = context.WithCancel(ctx.Context())
    ctx.WithValue("cancel", application.ctx)
    application.InitializeLogger()
    ctx.WithValue(ApplicationKey{}, application)
    return application
}

// Application Version Information for CLI
func (application *Application) PrettyVersionCLI() {
    format.Println(format.Sprintf("%s Revision: %s", application.originalName, application.version.StringPretty()))
    format.Println(format.Sprintf("%s Built With: %s", application.originalName, strings.Replace(runtime.Version(), "go", "", -1)))
    format.Println(format.Sprintf("%s Environment: %s", application.originalName, application.Environment()))
    format.Println(format.Sprintf("%s Service Name: %s-%s\n", application.originalName, application.ConsulConfigurationKey(), application.Environment()))
}

// Application information for API
func (application *Application) PrettyVersionForAPI() string {
    return format.Sprintf("%s %s\n",
        application.originalName,
        application.version.StringPretty(),
    )
}

// Get application container name
func (application *Application) ContainerName() string {
    if application.EnvironmentDescriptor().Type == 1 {
        return strings.ReplaceAll(application.Name(), ".", "_") + ".1.xyznonexistingtask"
    } else {
        hostname, err := os.Hostname()
        if err != nil {
            return strings.ReplaceAll(application.Name(), ".", "_") + ".1.xyznonexistingtask"
        } else {
            return hostname
        }
    }
}

// Get service name for manager
func (application *Application) ManagerServiceName() string {
    containerName := application.ContainerName()
    partedName := strings.Split(containerName, ".")
    serviceName := strings.Replace(containerName, "." + partedName[len(partedName) - 1], "", -1)
    return strings.TrimSpace(serviceName)
}

// Get application vendor name
func (application *Application) Vendor() string {
    return application.vendor
}

// Get application original vendor name
func (application *Application) OriginalVendorName() string {
    return application.originalVendor
}

// Get application name
func (application *Application) Name() string {
    return application.name
}

// Get application original name
func (application *Application) OriginalName() string {
    return application.originalName
}

// Get application version descriptor
func (application *Application) VersionDescriptor() VersionDescriptor {
    return application.version
}

// Get application version as string
func (application *Application) Version() string {
    return application.VersionDescriptor().String()
}

// Get application UUID
func (application *Application) UUID() string {
    return application.uuid.String()
}

// Get application environment descriptor
func (application *Application) EnvironmentDescriptor() ApplicationEnvironment {
    return application.environment
}

// Get application environment as string
func (application *Application) Environment() string {
    return application.EnvironmentDescriptor().Name
}

// Get application environment prefix
func (application *Application) EnvironmentPrefix() string {
    return strings.ToUpper(format.Sprintf("%s_%s", application.Environment(), strings.ReplaceAll(application.ConsulConfigurationKey(), ".", "_")))
}

// Get application service name
func (application *Application) ServiceName() string {
    return helpers.ToLowerAndTrim(format.Sprintf("%s-%s",
        application.ConsulConfigurationKey(),
        application.Environment(),
    ))
}

// Generate consul configuration key for microservice
func (application *Application) ConsulConfigurationKey() string {
    return format.Sprintf("%s.%s", application.Vendor(), application.Name())
}

// Get application configuration
func (application *Application) Configuration() *config {
    return application.config
}

// Initialize application configuration
func (application *Application) InitializeConfiguration() error {
    if application.config == nil {
        application.config = &config{}
    }
    return application.config.init(application)
}

// Initialize go-micro service
func (application *Application) InitializeMicroService() error {
    microInstance := &Micro {}
    err := microInstance.Initialize(application)
    if err != nil {
        return err
    }
    application.micro = microInstance
    return nil
}

// Initialize logger
func (application *Application) InitializeLogger() {
    logrusInstance := logrus.New()
    if application.EnvironmentDescriptor().Type == 4 {
        logrusInstance.SetFormatter(&logrus.JSONFormatter{})
    }
    log.SetLogger(logrusInstance.WithFields(logrus.Fields {
        "uuid":         application.uuid.String(),
        "service":      application.ConsulConfigurationKey(),
        "version":      application.Version(),
        "environment":  application.Environment(),
    }))
}

// Start microservice
func (application *Application) StartMicroService(registrations ...Registration) {
    application.Service().Init()
    for _, registration := range registrations {
        application.RegisterSubscriptionServer(registration.Topic, registration.Event, registration.Options)
    }
    if err := application.Service().Run(); err != nil {
        log.Panic(err)
    }
}

// Get micro service instance
func (application *Application) Service() micro.Service {
    if application.micro != nil {
        return application.micro.Service
    }
    log.Panic(MicroServiceNotInitialized.ToError(nil))
    return nil
}

// Get application context
func (application *Application) Context() context.Context {
    return application.ctx
}

// Get port application is running on
func (application *Application) Port() uint64 {
    partedString := strings.Split(application.Service().Server().Options().Address, ":")
    servicePortAsString := partedString[len(partedString) - 1]

    if servicePortAsString == "" {
        servicePortAsString = "0"
    }
    port, err := strconv.ParseUint(servicePortAsString, 10, 64)
    if err != nil {
        panic(err)
    }
    return port
}

// Add new callback which will be called before application is stopped
func (application *Application) BeforeStop(callback func() error) {
    application.beforeStop = append(application.beforeStop, callback)
}

// Get application instance from context
func FromContext() (*Application, bool) {
    application, done := ctx.Value(ApplicationKey{}).(*Application)
    return application, done
}

// Global method to add `beforeStop` callback
func BeforeStop(callback func() error) {
    application := ctx.Value(ApplicationKey{}).(*Application)
    application.BeforeStop(callback)
}

// Wait for the OS signal, run callback and then terminate application
func WaitForOSSignal(timeout time.Duration) {
    application := ctx.Value(ApplicationKey{}).(*Application)
    signalChannel := make(chan os.Signal)
    signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
    <- signalChannel
    application.cancel()
    for _, callback := range application.beforeStop {
        err := callback()
        if err != nil {
            log.Print(BeforeStopCallbackError.ToError(err))
        }
    }
    <-time.After(time.Second * timeout)
    os.Exit(1)
}

// Create UUID for the application
func createApplicationUUID() uuid.UUID {
    var possibleUUID    uuid.UUID
    var err             error

    possibleUUID, err = uuid.NewRandom()
    if err != nil {
        // All UUID generation process just went south, fallback to manually generated one
        bytesArray := make([]byte, 16)
        _, err = rand.Read(bytesArray)
        if err != nil {
            // Well, this is it, there is nothing we can do here, just panic and kill the application
            log.Panic("Failed to generate application UUID", err)
        }
        possibleUUID, err = uuid.FromBytes(bytesArray)
        if err != nil {
            log.Panic("This is the second time `github.com/google/uuid` fails to generate/process UUID", err)
        }
    }
    return possibleUUID
}

// Normalize vendor name
func normalizeVendorName(vendorName string) string {
    return helpers.ToLowerAndReplace(vendorName, " ", "")
}

// Normalize application name
func normalizeApplicationName(applicationName string) string {
    return helpers.ToLowerAndReplace(applicationName, " ", ".")
}
