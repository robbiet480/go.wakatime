// Package wakatime provides a Golang interface for accessing the WakaTime API
package wakatime

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// APIBaseURL is the Base API URL.
const APIBaseURL = "https://wakatime.com/api/v1/"

// The WakaTime struct, containing the APIKey.
type WakaTime struct {
	APIKey string
}

// Response is a response from WakaTime.
type Response struct {
	Result       interface{} `json:"data,omitempty"`
	Start        *timestamp  `json:"start,omitempty"`
	End          *timestamp  `json:"end,omitempty"`
	Timezone     string      `json:"timezone,omitempty"`
	Branches     []string    `json:"branches,omitempty"`
	ErrorMessage string      `json:"error,omitempty"`
	Range        string      `json:"range,omitempty"`
	Language     string      `json:"language,omitempty"`
	ModifiedAt   *time.Time  `json:"modified_at,omitempty"`
	CurrentUser  CurrentUser `json:"current_user,omitempty"`
}

// decodeResponse decodes the response from r writing the result into the struct
// pointed to by want. From: https://gist.github.com/kylelemons/2407845
func decodeResponse(r io.Reader, want interface{}) error {
	resp := &Response{Result: want}
	if err := json.NewDecoder(r).Decode(resp); err != nil {
		return err
	}
	if resp.ErrorMessage != "" {
		return errors.New(resp.ErrorMessage)
	}
	return nil
}

var httpClient = &http.Client{}

// NewWakaTime returns a WakaTime struct with the key filled in
func NewWakaTime(key string) *WakaTime {
	return &WakaTime{
		APIKey: key,
	}
}

// getURL is a handy wrapper function for interfacing with the WakaTime API
func (w *WakaTime) getURL(url string, decode bool, result interface{}) error {
	fullURL := fmt.Sprintf("%s%s", APIBaseURL, url)
	req, err := http.NewRequest("GET", fullURL, nil)
	b64edKey := base64.StdEncoding.EncodeToString([]byte(w.APIKey))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", b64edKey))
	req.Header.Add("User-Agent", "go.wakatime")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		err := fmt.Sprintf("Status code was not 200, was instead %d", resp.StatusCode)
		return errors.New(err)
	}
	defer resp.Body.Close()
	if decode {
		if err := decodeResponse(resp.Body, &result); err != nil {
			fmt.Printf("%q: error %q\n", resp.Body, err)
			return err
		}
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, result)
}

func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case time.Time:
			v = f.Interface().(time.Time).Format("01/02/2006")
		case *time.Time:
			v = f.Interface().(*time.Time).Format("01/02/2006")
		case string:
			v = f.String()
		case []string:
			v = strings.Join(f.Interface().([]string), ",")
		}
		if v == "" || v == "0" {
			continue
		}
		if typ.Field(i).Name == "User" {
			continue
		}
		key := strings.ToLower(typ.Field(i).Name)
		if typ.Field(i).Tag.Get("parameter_name") != "" {
			key = typ.Field(i).Tag.Get("parameter_name")
		}
		values.Set(key, v)
	}
	return
}
