package wakatime

import (
	"fmt"
	"time"
)

// User is a single user.
type User struct {
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	Email                string     `json:"email,omitempty"`
	EmailPublic          bool       `json:"email_public,omitempty"`
	FullName             string     `json:"full_name,omitempty"`
	HumanReadableWebsite string     `json:"human_readable_website,omitempty"`
	ID                   string     `json:"id,omitempty"`
	LastHeartbeat        *time.Time `json:"last_heartbeat,omitempty"`
	LastPlugin           string     `json:"last_plugin,omitempty"`
	LastPluginName       string     `json:"last_plugin_name,omitempty"`
	LastProject          string     `json:"last_project,omitempty"`
	Location             string     `json:"location,omitempty"`
	LoggedTimePublic     bool       `json:"logged_time_public,omitempty"`
	ModifiedAt           *time.Time `json:"modified_at,omitempty"`
	Photo                string     `json:"photo,omitempty"`
	PhotoPublic          bool       `json:"photo_public,omitempty"`
	Plan                 string     `json:"plan,omitempty"`
	Timezone             string     `json:"timezone,omitempty"`
	Username             string     `json:"username,omitempty"`
	Website              string     `json:"website,omitempty"`
}

// A UserAgent is a plugins which has sent data for this user.
type UserAgent struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Editor    string     `json:"editor,omitempty"`
	ID        string     `json:"id,omitempty"`
	LastSeen  *time.Time `json:"last_seen,omitempty"`
	OS        string     `json:"os,omitempty"`
	Value     string     `json:"value,omitempty"`
	Version   string     `json:"version,omitempty"`
}

// GetUser returns the user for the given username.
func (waka *WakaTime) GetUser(username string) (User, error) {
	if username == "" {
		username = "current"
	}
	var user User
	err := waka.getURL(fmt.Sprintf("users/%s", username), true, &user)
	return user, err
}

// GetUserAgents returns the user agents for the given username.
func (waka *WakaTime) GetUserAgents(username string) ([]UserAgent, error) {
	if username == "" {
		username = "current"
	}
	var userAgent []UserAgent
	err := waka.getURL(fmt.Sprintf("users/%s/user_agents", username), true, &userAgent)
	return userAgent, err
}
