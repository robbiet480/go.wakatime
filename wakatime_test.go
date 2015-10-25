package wakatime

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var WakaTimeAPIKey = os.Getenv("WAKATIME_API_KEY")

func TestMain(m *testing.M) {
	if WakaTimeAPIKey == "" {
		fmt.Println("No API key provided in the WAKATIME_API_KEY environment variable!")
		os.Exit(1)
	} else {
		os.Exit(m.Run())
	}
}

func TestGetDurations(t *testing.T) {
	client := NewWakaTime(WakaTimeAPIKey)
	yesterday := time.Now().AddDate(0, 0, -1)
	response, err := client.GetDurations(&DurationParameters{Date: &yesterday})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestGetHeartbeats(t *testing.T) {
	client := NewWakaTime(WakaTimeAPIKey)
	yesterday := time.Now().AddDate(0, 0, -1)
	response, err := client.GetHeartbeats(&HeartbeatParameters{
		Date: &yesterday,
		Show: []string{"time", "entity", "type", "project", "language", "branch", "is_write", "is_debugging"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestGetLeaders(t *testing.T) {
	client := NewWakaTime(WakaTimeAPIKey)
	response, err := client.GetLeaders("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestGetStats(t *testing.T) {
	client := NewWakaTime(WakaTimeAPIKey)
	response, err := client.GetStats(&StatsParameters{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestGetSummaries(t *testing.T) {
	client := NewWakaTime(WakaTimeAPIKey)
	now := time.Now()
	yesterday := time.Now().AddDate(0, 0, -1)
	response, err := client.GetSummaries(&SummaryParameters{
		Start: &yesterday,
		End:   &now,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestGetUser(t *testing.T) {
	client := NewWakaTime(WakaTimeAPIKey)
	response, err := client.GetUser("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestGetUserAgents(t *testing.T) {
	client := NewWakaTime(WakaTimeAPIKey)
	response, err := client.GetUserAgents("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
