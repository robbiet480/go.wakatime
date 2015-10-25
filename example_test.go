package wakatime

import (
	"fmt"
	"time"
)

// This example gets the durations for the authenticated user.
func ExampleGetDurations() {
	client := NewWakaTime(WakaTimeAPIKey)
	yesterday := time.Now().AddDate(0, 0, -1)
	response, err := client.GetDurations(&DurationParameters{Date: &yesterday})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the heartbeats for the authenticated user.
func ExampleGetHeartbeats() {
	client := NewWakaTime(WakaTimeAPIKey)
	yesterday := time.Now().AddDate(0, 0, -1)
	response, err := client.GetHeartbeats(&HeartbeatParameters{
		Date: &yesterday,
		Show: []string{"time", "entity", "type", "project", "language", "branch", "is_write", "is_debugging"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the leaders for the authenticated user.
func ExampleGetLeaders() {
	client := NewWakaTime(WakaTimeAPIKey)
	response, err := client.GetLeaders("")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the stats for the authenticated user.
func ExampleGetStats() {
	client := NewWakaTime(WakaTimeAPIKey)
	response, err := client.GetStats(&StatsParameters{})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the summaries for the authenticated user.
func ExampleGetSummaries() {
	client := NewWakaTime(WakaTimeAPIKey)
	now := time.Now()
	yesterday := time.Now().AddDate(0, 0, -1)
	response, err := client.GetSummaries(&SummaryParameters{
		Start: &yesterday,
		End:   &now,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the user for the authenticated user.
func ExampleGetUser() {
	client := NewWakaTime(WakaTimeAPIKey)
	response, err := client.GetUser("")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the userAgents for the authenticated user.
func ExampleGetUserAgents() {
	client := NewWakaTime(WakaTimeAPIKey)
	response, err := client.GetUserAgents("")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
