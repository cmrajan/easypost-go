package easypost

import "time"

type Report struct {
	Id               string    `json:"id,omitempty"`               //Unique, begins with "plrep_" (Payment Log Report), "refrep_" (Refund Report), "shprep_" (Shipment Report), or "trkrep_" (Tracker Report)
	Object           string    `json:"object,omitempty"`           //"PaymentLogReport", "RefundReport", "ShipmentReport", or "TrackerReport"
	Mode             string    `json:"mode,omitempty"`             //"test" or "production"
	Status           string    `json:"status,omitempty"`           //"new", "available", "failed", or null
	Start_Date       string    `json:"start_date,omitempty"`       //A date string in YYYY-MM-DD form eg: "2016-02-02"
	End_Date         string    `json:"end_date,omitempty"`         //A date string in YYYY-MM-DD form eg: "2016-02-03"
	Include_Children bool      `json:"include_children,omitempty"` //Set true if you would like to include Refunds /% include docs/api/type-link.html type_name="Shipments" type_href="shipment-object"%}/Trackers created by child users
	Url              string    `json:"url,omitempty"`              //A url that contains a link to the Report. Expires 30 seconds after retrieving this object
	Url_Expires_At   time.Time `json:"url_expires_at,omitempty"`   //Url expiring time
	Send_Email       bool      `json:"send_email,omitempty"`       //Set true if you would like to send an email containing the Report
	Created_At       time.Time `json:"created_at,omitempty"`       //
	Updated_At       time.Time `json:"updated_at,omitempty"`       //
}
