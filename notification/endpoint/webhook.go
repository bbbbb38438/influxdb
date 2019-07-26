package endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/influxdata/influxdb"
)

var _ influxdb.NotificationEndpoint = &WebHook{}

// WebHook is the notification endpoint config of webhook.
type WebHook struct {
	Base
	// Path is the API path of WebHook
	URL string `json:"url"`
	// Token is the bearer token for authorization
	Token           influxdb.SecretField `json:"token"`
	Username        influxdb.SecretField `json:"username"`
	Password        influxdb.SecretField `json:"password"`
	AuthMethod      string               `json:"authmethod"`
	Method          string               `json:"method"`
	ContentTemplate string               `json:"contentTemplate"`
}

// SecretKeys returns a set of secret keys.
func (s WebHook) SecretKeys() []influxdb.SecretField {
	arr := make([]influxdb.SecretField, 0)
	if s.Token != "" {
		arr = append(arr, s.Token)
	}
	if s.Username != "" {
		arr = append(arr, s.Username)
	}
	if s.Password != "" {
		arr = append(arr, s.Password)
	}
	return arr
}

var goodWebHookAuthMethod = map[string]bool{
	"none":   false,
	"basic":  false,
	"bearer": false,
}

// Valid returns error if some configuration is invalid
func (s WebHook) Valid() error {
	if err := s.Base.valid(); err != nil {
		return err
	}
	if s.URL == "" {
		return &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "webhook endpoint URL is empty",
		}
	}
	if _, err := url.Parse(s.URL); err != nil {
		return &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  fmt.Sprintf("webhook endpoint URL is invalid: %s", err.Error()),
		}
	}
	if _, ok := goodWebHookAuthMethod[s.AuthMethod]; !ok {
		return &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "invalid webhook auth method",
		}
	}
	if s.AuthMethod == "basic" && s.Username == "" {
		return &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "invalid webhook username for basic auth",
		}
	}
	if s.AuthMethod == "bearer" && s.Token == "" {
		return &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "invalid webhook token for bearer auth",
		}
	}

	return nil
}

type webhookAlias WebHook

// MarshalJSON implement json.Marshaler interface.
func (s WebHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			webhookAlias
			Type string `json:"type"`
		}{
			webhookAlias: webhookAlias(s),
			Type:         s.Type(),
		})
}

// Type returns the
func (s WebHook) Type() string {
	return "webhook"
}

// ParseResponse will parse the http response from webhook.
func (s WebHook) ParseResponse(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return &influxdb.Error{
			Msg: string(body),
		}
	}
	return nil
}
