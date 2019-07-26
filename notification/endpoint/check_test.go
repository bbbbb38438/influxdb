package endpoint_test

import (
	"testing"

	"github.com/influxdata/influxdb"
	"github.com/influxdata/influxdb/notification/endpoint"
	influxTesting "github.com/influxdata/influxdb/testing"
)

const (
	id1 = "020f755c3c082000"
	id2 = "020f755c3c082001"
	id3 = "020f755c3c082002"
)

func numPtr(f float64) *float64 {
	p := new(float64)
	*p = f
	return p
}

var goodBase = endpoint.Base{
	ID:          influxTesting.MustIDBase16(id1),
	Name:        "name1",
	OrgID:       influxTesting.MustIDBase16(id3),
	Status:      influxdb.Active,
	Description: "desc1",
}

func TestValidEndpoint(t *testing.T) {
	cases := []struct {
		name string
		src  influxdb.NotificationEndpoint
		err  error
	}{
		{
			name: "invalid endpoint id",
			src:  &endpoint.Slack{},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "Notification Endpoint ID is invalid",
			},
		},
		{
			name: "empty name",
			src: &endpoint.SMTP{
				Base: endpoint.Base{
					ID: influxTesting.MustIDBase16(id1),
				},
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "Notification Endpoint Name can't be empty",
			},
		},
		{
			name: "invalid status",
			src: &endpoint.PagerDuty{
				Base: endpoint.Base{
					ID:    influxTesting.MustIDBase16(id1),
					Name:  "name1",
					OrgID: influxTesting.MustIDBase16(id3),
				},
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "invalid status",
			},
		},
		{
			name: "empty slack url",
			src: &endpoint.Slack{
				Base: goodBase,
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "slack endpoint URL is empty",
			},
		},
		{
			name: "invalid slack url",
			src: &endpoint.Slack{
				Base: goodBase,
				URL:  "posts://er:{DEf1=ghi@:5432/db?ssl",
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "slack endpoint URL is invalid: parse posts://er:{DEf1=ghi@:5432/db?ssl: net/url: invalid userinfo",
			},
		},
		{
			name: "empty slack token",
			src: &endpoint.Slack{
				Base: goodBase,
				URL:  "localhost",
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "slack endpoint token is empty",
			},
		},
		{
			name: "empty smtp host",
			src: &endpoint.SMTP{
				Base: goodBase,
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "smtp host is empty",
			},
		},
		{
			name: "empty smtp port",
			src: &endpoint.SMTP{
				Base: goodBase,
				Host: "localhost",
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "smtp port can't be 0",
			},
		},
		{
			name: "bad smtp host",
			src: &endpoint.SMTP{
				Base: goodBase,
				Port: 465,
				Host: "hoho://fdads@{=fdsa:",
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "smtp host is invalid: parse hoho://fdads@{=fdsa:: invalid character \"{\" in host name",
			},
		},
		{
			name: "empty pagerduty url",
			src: &endpoint.PagerDuty{
				Base: goodBase,
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "pagerduty endpoint URL is empty",
			},
		},
		{
			name: "invalid pagerduty url",
			src: &endpoint.PagerDuty{
				Base: goodBase,
				URL:  "posts://er:{DEf1=ghi@:5432/db?ssl",
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "pagerduty endpoint URL is invalid: parse posts://er:{DEf1=ghi@:5432/db?ssl: net/url: invalid userinfo",
			},
		},
		{
			name: "empty routine key",
			src: &endpoint.PagerDuty{
				Base: goodBase,
				URL:  "localhost",
			},
			err: &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  "pagerduty routing key is empty",
			},
		},
	}
	for _, c := range cases {
		got := c.src.Valid()
		influxTesting.ErrorsEqual(t, got, c.err)
	}
}

// var timeGen1 = mock.TimeGenerator{FakeValue: time.Date(2006, time.July, 13, 4, 19, 10, 0, time.UTC)}
// var timeGen2 = mock.TimeGenerator{FakeValue: time.Date(2006, time.July, 14, 5, 23, 53, 10, time.UTC)}

// func TestJSON(t *testing.T) {
// 	cases := []struct {
// 		name string
// 		src  influxdb.NotificationEndpoint
// 	}{
// 		{
// 			name: "simple Deadman",
// 			src: &endpoint.Deadman{
// 				Base: endpoint.Base{
// 					ID:              influxTesting.MustIDBase16(id1),
// 					AuthorizationID: influxTesting.MustIDBase16(id2),
// 					Name:            "name1",
// 					OrgID:           influxTesting.MustIDBase16(id3),
// 					Status:          influxdb.Active,
// 					Every:           influxdb.Duration{Duration: time.Hour},
// 					Tags: []notification.Tag{
// 						{
// 							Key:   "k1",
// 							Value: "v1",
// 						},
// 						{
// 							Key:   "k2",
// 							Value: "v2",
// 						},
// 					},
// 					CRUDLog: influxdb.CRUDLog{
// 						CreatedAt: timeGen1.Now(),
// 						UpdatedAt: timeGen2.Now(),
// 					},
// 				},
// 				TimeSince:  33,
// 				ReportZero: true,
// 				Level:      notification.Warn,
// 			},
// 		},
// 		{
// 			name: "simple threshold",
// 			src: &endpoint.SMTP{
// 				Base: endpoint.Base{
// 					ID:              influxTesting.MustIDBase16(id1),
// 					Name:            "name1",
// 					AuthorizationID: influxTesting.MustIDBase16(id2),
// 					OrgID:           influxTesting.MustIDBase16(id3),
// 					Status:          influxdb.Active,
// 					Every:           influxdb.Duration{Duration: time.Hour},
// 					Tags: []notification.Tag{
// 						{
// 							Key:   "k1",
// 							Value: "v1",
// 						},
// 						{
// 							Key:   "k2",
// 							Value: "v2",
// 						},
// 					},
// 					CRUDLog: influxdb.CRUDLog{
// 						CreatedAt: timeGen1.Now(),
// 						UpdatedAt: timeGen2.Now(),
// 					},
// 				},
// 				SMTPs: []endpoint.SMTPConfig{
// 					{AllValues: true, LowerBound: numPtr(-1.36)},
// 					{LowerBound: numPtr(10000), UpperBound: numPtr(500)},
// 				},
// 			},
// 		},
// 	}
// 	for _, c := range cases {
// 		b, err := json.Marshal(c.src)
// 		if err != nil {
// 			t.Fatalf("%s marshal failed, err: %s", c.name, err.Error())
// 		}
// 		got, err := endpoint.UnmarshalJSON(b)
// 		if err != nil {
// 			t.Fatalf("%s unmarshal failed, err: %s", c.name, err.Error())
// 		}
// 		if diff := cmp.Diff(got, c.src); diff != "" {
// 			t.Errorf("failed %s, NotificationEndpoint are different -got/+want\ndiff %s", c.name, diff)
// 		}
// 	}
// }
