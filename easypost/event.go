package easypost

import "time"

type Event struct {
	Object              string            `json:"object,omitempty"`              //"Event"
	Id                  string            `json:"id,omitempty"`                  //Unique identifier, begins with "evt_"
	Mode                string            `json:"mode,omitempty"`                //"test" or "production"
	Description         string            `json:"description,omitempty"`         //Result type and event name, see the "Possible Event Types" section for more information
	Previous_Attributes map[string]string `json:"previous_attributes,omitempty"` //Previous values of relevant result attributes
	Result              map[string]string `json:"result,omitempty"`              //The object associated with the Event. See the "object" attribute on the result to determine its specific type
	Status              string            `json:"status,omitempty"`              //The current status of the event. Possible values are "completed", "failed", "in_queue", "retrying", or "pending" (deprecated)
	Pending_Urls        []string          `json:"pending_urls,omitempty"`        //Webhook URLs that have not yet been successfully notified as of the time this webhook event was sent. The URL receiving the Event will still be listed in pending_urls, as will any other URLs that receive the Event at the same time
	Completed_Urls      []string          `json:"completed_urls,omitempty"`      //Webhook URLs that have already been successfully notified as of the time this webhook was sent
	Created_At          time.Time         `json:"created_at,omitempty"`          //
	Updated_At          time.Time         `json:"updated_at,omitempty"`          //
}
