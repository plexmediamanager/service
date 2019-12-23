package service

import "github.com/plexmediamanager/service/errors"

const (
    ServiceID       errors.Service      =   0
)

var (
    InvalidBrokerConfiguration = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeUndefined,
            ErrorNumber:    1,
        },
        Message:    "Invalid broken configuration",
    }
    BrokerKeyNotFound = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeUndefined,
            ErrorNumber:    2,
        },
        Message:    "Broker key could not be found",
    }
    BrokerNotConfigured = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeUndefined,
            ErrorNumber:    3,
        },
        Message:    "Broker is not configured yet",
    }
    BrokerNotSupported = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeUndefined,
            ErrorNumber:    4,
        },
        Message:    "Specified broker is not supported: %s",
    }
    BrokerInitializationError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeUndefined,
            ErrorNumber:    5,
        },
        Message:    "Broker failed to complete initialization process",
    }
    BrokerConnectionError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeUndefined,
            ErrorNumber:    6,
        },
        Message:    "Failed to connect to broker",
    }
    MicroServiceNotInitialized = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeUndefined,
            ErrorNumber:    7,
        },
        Message:    "You have to initialize micro service first",
    }
    BeforeStopCallbackError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeUndefined,
            ErrorNumber:    8,
        },
        Message:    "Failed to execute `beforeStop` callback as it has errored",
    }
    UnableToConnectToConsul = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeNetwork,
            ErrorNumber:    1,
        },
        Message:    "Unable to connect to remote Consul instance",
    }
    UnableToGetKeyFromConsulKV = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    1,
        },
        Message:    "Unable to retrieve requested key from Consul Key/Value storage: %s",
    }
    UnableToGetServiceConfigurationFromConsulKV = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    2,
        },
        Message:    "Unable to retrieve microservice configuration from Consul using service key: %s",
    }
    UnableToGetSubscriptionsInformationForVendorFromConsul = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    3,
        },
        Message:    "Unable to retrieve microservice configuration from Consul using service key: %s",
    }
    UnableToDecodeSubscriptionsConfiguration = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    4,
        },
        Message:    "Unable to decode subscriptions configuration retrieved from Consul",
    }
    ViperUnableToMergeConfigs = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    5,
        },
        Message:    "Viper is not able to merge provided configurations",
    }
)
