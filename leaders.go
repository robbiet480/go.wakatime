package wakatime

import "fmt"

// RunningTotal is the total time logged for this user or leader.
type RunningTotal struct {
	DailyAverage              int        `json:"daily_average,omitempty"`
	HumanReadableDailyAverage string     `json:"human_readable_daily_average,omitempty"`
	HumanReadableTotal        string     `json:"human_readable_total,omitempty"`
	Languages                 []Language `json:"languages,omitempty"`
	TotalSeconds              int        `json:"total_seconds,omitempty"`
}

// CurrentUser is the user making the leaders request.
type CurrentUser struct {
	RunningTotal RunningTotal `json:"running_total,omitempty"`
	User         User         `json:"user,omitempty"`
}

// Leader is a user on the leaderboard.
type Leader struct {
	Rank         int          `json:"rank,omitempty"`
	RunningTotal RunningTotal `json:"running_total,omitempty"`
	User         User         `json:"user,omitempty"`
}

// Leaders is a list of users ranked by logged time in descending order. Same as the public leaderboards.
type Leaders struct {
	CurrentUser CurrentUser `json:"current_user"`
	Data        []Leader    `json:"data"`
	Language    string      `json:"language"`
	ModifiedAt  string      `json:"modified_at"`
	Range       string      `json:"range"`
}

// GetLeaders returns the leaderboard.
// If a language is given, it returns the leaderboard for only that language.
func (waka *WakaTime) GetLeaders(language string) (Leaders, error) {
	if language != "" {
		language = fmt.Sprintf("?language=%s", language)
	}
	var leaders Leaders
	err := waka.getURL(fmt.Sprintf("leaders/%s", language), false, &leaders)
	return leaders, err
}
