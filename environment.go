package service

import (
    format "fmt"
    "github.com/plexmediamanager/service/helpers"
    goDotEnv "github.com/joho/godotenv"
    "log"
    "os"
    "strings"
)

type EnvironmentType uint8

const (
    Development         EnvironmentType     =   1
    Staging             EnvironmentType     =   2
    Debug               EnvironmentType     =   3
    Production          EnvironmentType     =   4
)

var EnvironmentLookup = map[EnvironmentType] string {
    Development:        "development",
    Staging:            "staging",
    Debug:              "debug",
    Production:         "production",
}

type ApplicationEnvironment struct {
    Name                string
    Type                EnvironmentType
}

// Get application environment
func GetApplicationEnvironment() ApplicationEnvironment {
    var environmentStringValue string = strings.TrimSpace(strings.ToLower(os.Getenv("ENVIRONMENT")))
    var currentEnvironment EnvironmentType = 0

    if strings.Contains(environmentStringValue, "prod") {
        currentEnvironment = Production
    } else if strings.Contains(environmentStringValue, "stag") {
        currentEnvironment = Staging
    } else if strings.Contains(environmentStringValue, "deb") {
        currentEnvironment = Debug
    } else {
        currentEnvironment = Development
    }

    return ApplicationEnvironment {
        Name: EnvironmentLookup[currentEnvironment],
        Type: currentEnvironment,
    }
}

// Preload environment variables
func preloadEnvironment(environment ApplicationEnvironment) {
    err := goDotEnv.Load("application.env")
    if err != nil {
        log.Panicln("application.env file not found in the root of the project")
    }
    err = goDotEnv.Load(format.Sprintf("%s.env", environment.Name))
    if err != nil {
        log.Println(format.Sprintf("I was unable to find conviguration for `%s (%s.env)` environment, i hope you know what you are doing", environment.Name, environment.Name))
    }
}

// Get application vendor name from the environment or use default if not set
func getApplicationVendorFromEnvironment() string {
    return helpers.GetEnvironmentVariableAsString("APPLICATION_VENDOR", "FreedomCore")
}

// Get application name from the environment or use default if not set
func getApplicationNameFromEnvironment() string {
    return helpers.GetEnvironmentVariableAsString("APPLICATION_NAME", "New Application")
}

// Get application version from the environment or use default if not set
func getApplicationVersionFromEnvironment() string {
    return helpers.GetEnvironmentVariableAsString("APPLICATION_VERSION", "1.0.0.0")
}