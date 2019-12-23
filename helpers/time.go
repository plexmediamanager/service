package helpers

import (
    "math"
    "strconv"
    "strings"
    "time"
)

// Convert seconds to duration
func SecondsToDuration(seconds time.Duration) time.Duration {
    return seconds * time.Second
}

// Convert minutes to duration
func MinutesToDuration(minutes time.Duration) time.Duration {
    return minutes * time.Minute
}

// Convert hours to duration
func HoursToDuration(hours time.Duration) time.Duration {
    return hours * time.Hour
}

func SecondsToHuman(input int) (result string) {
    years := math.Floor(float64(input) / 60 / 60 / 24 / 7 / 30 / 12)
    seconds := input % (60 * 60 * 24 * 7 * 30 * 12)
    months := math.Floor(float64(seconds) / 60 / 60 / 24 / 7 / 30)
    seconds = input % (60 * 60 * 24 * 7 * 30)
    weeks := math.Floor(float64(seconds) / 60 / 60 / 24 / 7)
    seconds = input % (60 * 60 * 24 * 7)
    days := math.Floor(float64(seconds) / 60 / 60 / 24)
    seconds = input % (60 * 60 * 24)
    hours := math.Floor(float64(seconds) / 60 / 60)
    seconds = input % (60 * 60)
    minutes := math.Floor(float64(seconds) / 60)
    seconds = input % 60

    timeParts := make([]string, 0)

    if seconds > 0 {
        timeParts = append(timeParts, Plural(int(seconds), "second"))
    }
    if minutes > 0 {
        timeParts = append(timeParts, Plural(int(minutes), "minute"))
    }
    if hours > 0 {
        timeParts = append(timeParts, Plural(int(hours), "hour"))
    }
    if days > 0 {
        timeParts = append(timeParts, Plural(int(days), "day"))
    }
    if weeks > 0 {
        timeParts = append(timeParts, Plural(int(weeks), "week"))
    }
    if months > 0 {
        timeParts = append(timeParts, Plural(int(months), "month"))
    }
    if years > 0 {
        timeParts = append(timeParts, Plural(int(years), "year"))
    }

    reverse(timeParts)
    result = strings.Join(timeParts, " ")
    return
}

func Plural(count int, singular string) (result string) {
    if (count == 1) || (count == 0) {
        result = strconv.Itoa(count) + " " + singular + " "
    } else {
        result = strconv.Itoa(count) + " " + singular + "s "
    }
    return
}

// Reverse array of strings
func reverse(ss []string) {
    last := len(ss) - 1
    for i := 0; i < len(ss)/2; i++ {
        ss[i], ss[last-i] = ss[last-i], ss[i]
    }
}