package wakatime

import (
	"errors"
	"fmt"
	"time"
)

// GrandTotal is the total amount of time logged for this summary.
type GrandTotal struct {
	Digital      string `json:"digital,omitempty"`
	Hours        int    `json:"hours,omitempty"`
	Minutes      int    `json:"minutes,omitempty"`
	Text         string `json:"text,omitempty"`
	TotalSeconds int    `json:"total_seconds,omitempty"`
}

// Range is the time range this summary covers.
type Range struct {
	Date      string     `json:"date,omitempty"`
	DateHuman string     `json:"date_human,omitempty"`
	End       *timestamp `json:"end,omitempty"`
	Start     *timestamp `json:"start,omitempty"`
	Text      string     `json:"text,omitempty"`
	Timezone  string     `json:"timezone,omitempty"`
}

// SummaryItem is an object in a summary
type SummaryItem struct {
	Digital      string  `json:"digital,omitempty"`
	Hours        int     `json:"hours,omitempty"`
	Minutes      int     `json:"minutes,omitempty"`
	Name         string  `json:"name,omitempty"`
	Percent      float64 `json:"percent,omitempty"`
	Seconds      int     `json:"seconds,omitempty"`
	Text         string  `json:"text,omitempty"`
	TotalSeconds int     `json:"total_seconds,omitempty"`
	Type         string  `json:"type,omitempty"`
}

// Summary is a user's logged time for the given time range.
type Summary struct {
	Editors          []SummaryItem `json:"editors,omitempty"`
	Entities         []SummaryItem `json:"entities,omitempty"`
	GrandTotal       GrandTotal    `json:"grand_total,omitempty"`
	Languages        []SummaryItem `json:"languages,omitempty"`
	OperatingSystems []SummaryItem `json:"operating_systems,omitempty"`
	Projects         []SummaryItem `json:"projects,omitempty"`
	Range            Range         `json:"range,omitempty"`
}

// SummaryParameters are the query parameters used for GetSummaries.
type SummaryParameters struct {
	Start    *time.Time
	End      *time.Time
	Project  string
	Branches string

	User string
}

// GetSummaries returns summaries for the given user and parameters.
func (waka *WakaTime) GetSummaries(summaryParameters *SummaryParameters) ([]Summary, error) {
	var summary []Summary
	if summaryParameters.User == "" {
		summaryParameters.User = "current"
	}
	if summaryParameters.Start == nil {
		return summary, errors.New("You must provide a start time!")
	}
	if summaryParameters.End == nil {
		return summary, errors.New("You must provide a end time!")
	}
	params := structToMap(summaryParameters)

	encodedParams := params.Encode()

	if encodedParams != "" {
		encodedParams = "?" + encodedParams
	}

	err := waka.getURL(fmt.Sprintf("users/%s/summaries%s", summaryParameters.User, encodedParams), true, &summary)
	return summary, err
}
