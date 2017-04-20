package easypost

import "time"

type Pickup struct {
	Id                 string            `json:"id,omitempty"`                 //unique, begins with "pickup_"
	Object             string            `json:"object,omitempty"`             //"Pickup"
	Reference          string            `json:"reference,omitempty"`          //An optional field that may be used in place of ID in some API endpoints
	Mode               string            `json:"mode,omitempty"`               //"test" or "production"
	Status             string            `json:"status,omitempty"`             //One of: "unknown", "scheduled", or "canceled"
	Min_Datetime       time.Time         `json:"min_datetime,omitempty"`       //The earliest time at which the package is available to pick up
	Max_Datetime       time.Time         `json:"max_datetime,omitempty"`       //The latest time at which the package is available to pick up. Must be later than the min_datetime
	Is_Account_Address bool              `json:"is_account_address,omitempty"` //Is the pickup address the account's address?
	Instructions       string            `json:"instructions,omitempty"`       //Additional text to help the driver successfully obtain the package
	Messages           []*Message        `json:"messages,omitempty"`           //A list of messages containing carrier errors encountered during pickup rate generation
	Confirmation       string            `json:"confirmation,omitempty"`       //The confirmation number for a booked pickup from the carrier
	Shipment           *Shipment         `json:"shipment,omitempty"`           //The associated Shipment
	Address            *Address          `json:"address,omitempty"`            //The associated Address
	Carrier_Accounts   []*CarrierAccount `json:"carrier_accounts,omitempty"`   //The list of carriers (if empty, all carriers were used) used to generate pickup rates
	Pickup_Rates       []*PickupRate     `json:"pickup_rates,omitempty"`       //The list of different pickup rates across valid carrier accounts for the shipment
	Created_At         time.Time         `json:"created_at,omitempty"`         //
	Updated_At         time.Time         `json:"updated_at,omitempty"`         //
}

type PickupRate struct {
	Id         string    `json:"id,omitempty"`         //unique, begins with 'pickuprate_'
	Object     string    `json:"object,omitempty"`     //"PickupRate"
	Mode       string    `json:"mode,omitempty"`       //"test" or "production"
	Service    string    `json:"service,omitempty"`    //service name
	Carrier    string    `json:"carrier,omitempty"`    //name of carrier
	Rate       string    `json:"rate,omitempty"`       //the actual rate quote for this service
	Currency   string    `json:"currency,omitempty"`   //currency for the rate
	Pickup_Id  string    `json:"pickup_id,omitempty"`  //the ID of the pickup this is a quote for
	Created_At time.Time `json:"created_at,omitempty"` //
	Updated_At time.Time `json:"updated_at,omitempty"` //
}
