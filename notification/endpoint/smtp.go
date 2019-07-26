package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/influxdata/influxdb"
	"github.com/influxdata/influxdb/toml"
)

var _ influxdb.NotificationEndpoint = &SMTP{}

// SMTP is the notification endpoint config of smtp.
type SMTP struct {
	Base
	Host     string               `json:"host"`
	Port     uint                 `json:"port"`
	Username influxdb.SecretField `json:"username"`
	Password influxdb.SecretField `json:"password"`
	// Close connection to SMTP server after idle timeout has elapsed
	IdleTimeout toml.Duration `json:"idle-timeout"`
}

// SecretKeys returns a set of secret keys.
func (s SMTP) SecretKeys() []influxdb.SecretField {
	arr := make([]influxdb.SecretField, 0)
	if s.Username != "" {
		arr = append(arr, s.Username)
	}
	if s.Password != "" {
		arr = append(arr, s.Password)
	}
	return arr
}

// Valid returns error if some configuration is invalid
func (s SMTP) Valid() error {
	if err := s.Base.valid(); err != nil {
		return err
	}
	if s.Host == "" {
		return &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "smtp host is empty",
		}
	}
	if s.Port == 0 {
		return &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "smtp port can't be 0",
		}
	}
	if _, err := url.Parse(s.Host); err != nil {
		return &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  fmt.Sprintf("smtp host is invalid: %s", err.Error()),
		}
	}
	return nil
}

type smtpAlias SMTP

// MarshalJSON implement json.Marshaler interface.
func (s SMTP) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			smtpAlias
			Type string `json:"type"`
		}{
			smtpAlias: smtpAlias(s),
			Type:      s.Type(),
		})
}

// Type returns the
func (s SMTP) Type() string {
	return "smtp"
}

// ParseResponse will parse the http response from smtp.
func (s SMTP) ParseResponse(resp *http.Response) error {
	if resp.StatusCode != http.StatusResetContent {
		return &influxdb.Error{
			Msg: "stmp response status: " + resp.Status,
		}
	}
	return nil
}
