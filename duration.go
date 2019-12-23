package service

import (
    "encoding/json"
    "time"
)

type Duration struct {
    time.Duration
}

func (duration Duration) MarshalJSON() ([]byte, error) {
    return []byte(`"` + duration.String() + `"`), nil
}

func (duration *Duration) UnmarshalJSON(data []byte) error {
    var durationAsString string
    if err := json.Unmarshal(data, &durationAsString); err != nil {
        return err
    }

    processedDuration, err := time.ParseDuration(durationAsString)
    if err != nil {
        return err
    }
    duration.Duration = processedDuration
    return nil
}