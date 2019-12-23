package helpers

import (
    "os"
    "strconv"
)

// Get environment variable as a string
func GetEnvironmentVariableAsString(key string, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        value = fallback
    }
    return value
}

// Get environment variable as a boolean
func GetEnvironmentVariableAsBool(key string, fallback bool) bool {
    var newValue bool
    value, exists := os.LookupEnv(key)
    if !exists {
        newValue = fallback
    } else {
        newValue = value == "true"
    }
    return newValue
}

// Get environment variable as an integer
func GetEnvironmentVariableAsInteger(key string, fallback int) int {
    var newValue int
    value, exists := os.LookupEnv(key)
    if !exists {
        newValue = fallback
    } else {
        value, err := strconv.Atoi(value)
        if err != nil {
            newValue = fallback
        }
        newValue = value
    }
    return newValue
}