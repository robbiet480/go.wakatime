package wakatime

import (
	"fmt"
	"strconv"
	"time"
)

type timestamp time.Time

func (t *timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}
func (t *timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}
	*t = timestamp(time.Unix(int64(ts), 0))
	return nil
}

func (t *timestamp) String() string {
	return time.Time(*t).String()
}

// Editor describes a text editor
type Editor struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	ID           string     `json:"id,omitempty"`
	ModifiedAt   *time.Time `json:"modified_at,omitempty"`
	Name         string     `json:"name,omitempty"`
	Percent      float64    `json:"percent,omitempty"`
	TotalSeconds int        `json:"total_seconds,omitempty"`
}

// Language describes a programming language
type Language struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	ID           string     `json:"id,omitempty"`
	ModifiedAt   *time.Time `json:"modified_at,omitempty"`
	Name         string     `json:"name,omitempty"`
	Percent      float64    `json:"percent,omitempty"`
	TotalSeconds int        `json:"total_seconds,omitempty"`
}

// OperatingSystem describes an operating system
type OperatingSystem struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	ID           string     `json:"id,omitempty"`
	ModifiedAt   *time.Time `json:"modified_at,omitempty"`
	Name         string     `json:"name,omitempty"`
	Percent      float64    `json:"percent,omitempty"`
	TotalSeconds int        `json:"total_seconds,omitempty"`
}

// Project describes a project
type Project struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	ID           string     `json:"id,omitempty"`
	ModifiedAt   *time.Time `json:"modified_at,omitempty"`
	Name         string     `json:"name,omitempty"`
	Percent      float64    `json:"percent,omitempty"`
	TotalSeconds int        `json:"total_seconds,omitempty"`
}
