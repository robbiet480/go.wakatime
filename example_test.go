package wakatime_test

import (
	"fmt"
	"os"
	"time"

	"github.com/robbiet480/go.wakatime"
)

// This example gets the durations for the authenticated user.
func ExampleWakaTime_GetDurations() {
	client := wakatime.NewWakaTime(os.Getenv("WAKATIME_API_KEY"))
	yesterday := time.Now().AddDate(0, 0, -1)
	response, err := client.GetDurations(&wakatime.DurationParameters{Date: &yesterday})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the heartbeats for the authenticated user.
func ExampleWakaTime_GetHeartbeats() {
	client := wakatime.NewWakaTime(os.Getenv("WAKATIME_API_KEY"))
	yesterday := time.Now().AddDate(0, 0, -1)
	response, err := client.GetHeartbeats(&wakatime.HeartbeatParameters{
		Date: &yesterday,
		Show: []string{"time", "entity", "type", "project", "language", "branch", "is_write", "is_debugging"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the leaders for the authenticated user.
func ExampleWakaTime_GetLeaders() {
	client := wakatime.NewWakaTime(os.Getenv("WAKATIME_API_KEY"))
	response, err := client.GetLeaders("")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the stats for the authenticated user.
func ExampleWakaTime_GetStats() {
	client := wakatime.NewWakaTime(os.Getenv("WAKATIME_API_KEY"))
	response, err := client.GetStats(&wakatime.StatsParameters{})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the summaries for the authenticated user.
func ExampleWakaTime_GetSummaries() {
	client := wakatime.NewWakaTime(os.Getenv("WAKATIME_API_KEY"))
	now := time.Now()
	yesterday := time.Now().AddDate(0, 0, -1)
	response, err := client.GetSummaries(&wakatime.SummaryParameters{
		Start: &yesterday,
		End:   &now,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the user for the authenticated user.
func ExampleWakaTime_GetUser() {
	client := wakatime.NewWakaTime(os.Getenv("WAKATIME_API_KEY"))
	response, err := client.GetUser("")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

// This example gets the userAgents for the authenticated user.
func ExampleWakaTime_GetUserAgents() {
	client := wakatime.NewWakaTime(os.Getenv("WAKATIME_API_KEY"))
	response, err := client.GetUserAgents("")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
