package easypost

import "time"

type Shipment struct {
	Id             string        `json:"id,omitempty"`             //Unique, begins with "shp_"
	Object         string        `json:"object,omitempty"`         //"Shipment"
	Reference      string        `json:"reference,omitempty"`      //An optional field that may be used in place of id in other API endpoints
	Mode           string        `json:"mode,omitempty"`           //"test" or "production"
	To_Address     *Address      `json:"to_address,omitempty"`     //The destination address
	From_Address   *Address      `json:"from_address,omitempty"`   //The origin address
	Return_Address *Address      `json:"return_address,omitempty"` //The shipper's address, defaults to from_address
	Buyer_Address  *Address      `json:"buyer_address,omitempty"`  //The buyer's address, defaults to to_address
	Parcel         *Parcel       `json:"parcel,omitempty"`         //The dimensions and weight of the package
	Customs_Info   *CustomsInfo  `json:"customs_info,omitempty"`   //Information for the processing of customs
	Scan_Form      *ScanForm     `json:"scan_form,omitempty"`      //Document created to manifest and scan multiple shipments
	Forms          []*Form       `json:"forms,omitempty"`          //All associated Form objects
	Insurance      *Insurance    `json:"insurance,omitempty"`      //The associated Insurance object
	Rates          []*Rate       `json:"rates,omitempty"`          //All associated Rate objects
	Selected_Rate  *Rate         `json:"selected_rate,omitempty"`  //The specific rate purchased for the shipment, or null if unpurchased or purchased through another mechanism
	Postage_Label  *PostageLabel `json:"postage_label,omitempty"`  //The associated PostageLabel object
	Messages       []*Message    `json:"messages,omitempty"`       //Any carrier errors encountered during rating, discussed in more depth below
	Options        *Options      `json:"options,omitempty"`        //All of the options passed to the shipment, discussed in more depth below
	Is_Return      bool          `json:"is_return,omitempty"`      //Set true to create as a return, discussed in more depth below
	Tracking_Code  string        `json:"tracking_code,omitempty"`  //If purchased, the tracking code will appear here as well as within the Tracker object
	Usps_Zone      string        `json:"usps_zone,omitempty"`      //The USPS zone of the shipment, if purchased with USPS
	Status         string        `json:"status,omitempty"`         //The current tracking status of the shipment
	Tracker        *Tracker      `json:"tracker,omitempty"`        //The associated Tracker object
	Fees           []*Fee        `json:"fees,omitempty"`           //The associated Fee objects charged to the billing user account
	Refund_Status  string        `json:"refund_status,omitempty"`  //The current status of the shipment refund process. Possible values are "submitted", "refunded", "rejected".
	Batch_Id       string        `json:"batch_id,omitempty"`       //The ID of the batch that contains this shipment, if any
	Batch_Status   string        `json:"batch_status,omitempty"`   //The current state of the associated BatchShipment
	Batch_Message  string        `json:"batch_message,omitempty"`  //The current message of the associated BatchShipment
	Created_At     time.Time     `json:"created_at,omitempty"`     //
	Updated_At     time.Time     `json:"updated_at,omitempty"`     //
}
