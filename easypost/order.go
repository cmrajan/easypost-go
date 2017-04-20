package easypost

import "time"

type Order struct {
	Id             string      `json:"id,omitempty"`             //Unique, begins with "order_"
	Object         string      `json:"object,omitempty"`         //"Order"
	Reference      string      `json:"reference,omitempty"`      //An optional field that may be used in place of id in other API endpoints
	Mode           string      `json:"mode,omitempty"`           //"test" or "production"
	To_Address     *Address    `json:"to_address,omitempty"`     //The destination address
	From_Address   *Address    `json:"from_address,omitempty"`   //The origin address
	Return_Address *Address    `json:"return_address,omitempty"` //The shipper's address, defaults to from_address
	Buyer_Address  *Address    `json:"buyer_address,omitempty"`  //The buyer's address, defaults to to_address
	Shipments      []*Shipment `json:"shipments,omitempty"`      //All associated Shipment objects
	Rates          []*Rate     `json:"rates,omitempty"`          //All associated Rate objects
	Messages       []*Message  `json:"messages,omitempty"`       //Any carrier errors encountered during rating
	Is_Return      bool        `json:"is_return,omitempty"`      //Set true to create as a return
	Created_At     time.Time   `json:"created_at,omitempty"`     //
	Updated_At     time.Time   `json:"updated_at,omitempty"`     //
}
