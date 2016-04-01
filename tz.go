// This function translates date/time information between time zones.
// Provide in input:
// - A Unix timestamp, i.e. number of seconds since 1970-01-01 00:00 UTC
// - A latitude and longitude
// The returned string contains wall time at the place of interest at
// the moment of the provided time stamp. To convert between (lat,lon) and
// an offset in seconds, a call is made to Google Map API to determine
// in which time zone the place is located, and if daylight saving time
// is observed at the moment of the requested time stamp.
// MIT License (c) nicolas314
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

type TZone struct {
    DSTOffset   int64   `json:"dstOffset"`
    RawOffset   int64   `json:"rawOffset"`
    TZId        string  `json:"timeZoneId"`
    TZName      string  `json:"timeZoneName"`
}

func GetDateTime(lat, lon float64, when int64) (string, error) {
    url:=fmt.Sprintf("https://maps.googleapis.com/maps/api/timezone/"+
                     "json"+
                     "?location=%.2f,%.2f&timestamp=%d",
                     lat, lon, when)
    resp, err := http.Get(url)
    if err!=nil {
        fmt.Println(err)
        return "", errors.New("cannot reach API")
    }
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    if err!=nil {
        fmt.Println(err)
        return "", errors.New("cannot read API response")
    }
    // fmt.Println(string(content))

    var ltz TZone
    err = json.Unmarshal(content, &ltz)
    if err!=nil {
        fmt.Println(err)
        return "", errors.New("cannot interpret API response")
    }
    at := time.Unix(when, 0)
    loc := time.FixedZone(ltz.TZId, int(ltz.DSTOffset+ltz.RawOffset))
    return fmt.Sprint(at.In(loc)), nil
}


func main() {
    now := time.Now().UTC().Unix()
    // Print current time in Paris
    at, _ := GetDateTime(48.0, 2.0, now)
    fmt.Println(at)

    // Print current time in Sydney
    at, _ = GetDateTime(-33.86, 151.20, now)
    fmt.Println(at)

    // Print current time in San Francisco
    at, _ = GetDateTime(37.77, -122.42, now)
    fmt.Println(at)

    return
}

