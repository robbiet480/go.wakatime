package wakatime

import (
	"errors"
	"fmt"
	"time"
)

// Duration is a user's logged time for the given day as an array of duration blocks.
type Duration struct {
	Dependencies string     `json:"dependencies,omitempty"`
	Duration     float64    `json:"duration,omitempty"`
	IsDebugging  bool       `json:"is_debugging,omitempty"`
	Project      string     `json:"project,omitempty"`
	Time         *timestamp `json:"time,omitempty"`
}

// DurationParameters are the query parameters used for GetDurations.
type DurationParameters struct {
	Date     *time.Time
	Project  string
	Branches string

	User string
}

// GetDurations returns durations for the given user and parameters.
func (waka *WakaTime) GetDurations(durationParameters *DurationParameters) ([]Duration, error) {
	var durations []Duration
	if durationParameters.User == "" {
		durationParameters.User = "current"
	}
	if durationParameters.Date == nil {
		return durations, errors.New("You must provide a date!")
	}
	params := structToMap(durationParameters)

	encodedParams := params.Encode()

	if encodedParams != "" {
		encodedParams = "?" + encodedParams
	}
	err := waka.getURL(fmt.Sprintf("users/%s/durations%s", durationParameters.User, encodedParams), true, &durations)
	return durations, err
}
