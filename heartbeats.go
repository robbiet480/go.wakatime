package wakatime

import (
	"fmt"
	"time"
)

// Heartbeat is a user's heartbeats sent from plugins for the given day.
type Heartbeat struct {
	Branch      string     `json:"branch,omitempty"`
	Entity      string     `json:"entity,omitempty"`
	ID          string     `json:"id,omitempty"`
	IsDebugging bool       `json:"is_debugging,omitempty"`
	IsWrite     bool       `json:"is_write,omitempty"`
	Language    string     `json:"language,omitempty"`
	Project     string     `json:"project,omitempty"`
	Time        *timestamp `json:"time,omitempty"`
	Type        string     `json:"type,omitempty"`
}

// HeartbeatParameters are the query parameters used for GetHeartbeats.
type HeartbeatParameters struct {
	Date *time.Time
	Show []string // Possible options: time,entity,type,project,language,branch,is_write,is_debugging

	User string
}

// GetHeartbeats returns heartbeats for the given user and parameters.
func (waka *WakaTime) GetHeartbeats(heartbeatParameters *HeartbeatParameters) ([]Heartbeat, error) {
	var heartbeat []Heartbeat
	if heartbeatParameters.User == "" {
		heartbeatParameters.User = "current"
	}
	params := structToMap(heartbeatParameters)

	encodedParams := params.Encode()

	if encodedParams != "" {
		encodedParams = "?" + encodedParams
	}

	err := waka.getURL(fmt.Sprintf("users/%s/heartbeats%s", heartbeatParameters.User, encodedParams), true, &heartbeat)
	return heartbeat, err
}
