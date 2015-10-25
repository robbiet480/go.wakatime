package wakatime

import (
	"fmt"
	"time"
)

// BestDay is the user's best day for the given time period.
type BestDay struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	Date         string     `json:"date,omitempty"`
	ID           string     `json:"id,omitempty"`
	ModifiedAt   *time.Time `json:"modified_at,omitempty"`
	TotalSeconds int        `json:"total_seconds,omitempty"`
}

// Stats is a user's logged time for the given time range.
type Stats struct {
	BestDay                   BestDay           `json:"best_day,omitempty"`
	CreatedAt                 string            `json:"created_at,omitempty"`
	DailyAverage              int               `json:"daily_average,omitempty"`
	Editors                   []Editor          `json:"editors,omitempty"`
	End                       *timestamp        `json:"end,omitempty"`
	Holidays                  int               `json:"holidays,omitempty"`
	HumanReadableDailyAverage string            `json:"human_readable_daily_average,omitempty"`
	HumanReadableTotal        string            `json:"human_readable_total,omitempty"`
	ID                        string            `json:"id,omitempty"`
	IsUpToDate                bool              `json:"is_up_to_date,omitempty"`
	Languages                 []Language        `json:"languages,omitempty"`
	ModifiedAt                string            `json:"modified_at,omitempty"`
	OperatingSystems          []OperatingSystem `json:"operating_systems,omitempty"`
	Project                   Project           `json:"project,omitempty"`
	Projects                  []Project         `json:"projects,omitempty"`
	Range                     string            `json:"range,omitempty"`
	Start                     *timestamp        `json:"start,omitempty"`
	Status                    string            `json:"status,omitempty"`
	Timeout                   int               `json:"timeout,omitempty"`
	Timezone                  string            `json:"timezone,omitempty"`
	TotalSeconds              int               `json:"total_seconds,omitempty"`
	UserID                    string            `json:"user_id,omitempty"`
	Username                  string            `json:"username,omitempty"`
	WritesOnly                bool              `json:"writes_only,omitempty"`
}

// StatsParameters are the query parameters used for GetStats.
type StatsParameters struct {
	Timeout    int
	WritesOnly bool `parameter_name:"writes_only"`
	Project    string

	Range string
	User  string
}

// GetStats returns stats for the given user and parameters.
func (waka *WakaTime) GetStats(statsParameters *StatsParameters) (Stats, error) {
	if statsParameters.User == "" {
		statsParameters.User = "current"
	}
	if statsParameters.Range == "" {
		statsParameters.Range = "last_7_days"
	}
	timeRange := statsParameters.Range
	statsParameters.Range = ""
	params := structToMap(statsParameters)

	encodedParams := params.Encode()

	if encodedParams != "" {
		encodedParams = "?" + encodedParams
	}

	var stats Stats
	err := waka.getURL(fmt.Sprintf("users/%s/stats/%s%s", statsParameters.User, timeRange, encodedParams), true, &stats)
	return stats, err
}
