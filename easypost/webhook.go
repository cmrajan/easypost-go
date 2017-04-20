package easypost

import "time"

type Webhook struct {
	Id          string    `json:"id,omitempty"`          //Unique, begins with "hook_"
	Object      string    `json:"object,omitempty"`      //"Webhook"
	Mode        string    `json:"mode,omitempty"`        //"test" or "production"
	Url         string    `json:"url,omitempty"`         //http://example.com
	Disabled_At time.Time `json:"disabled_at,omitempty"` //the timestamp at which the webhook was most recently disabled (if any)
}
