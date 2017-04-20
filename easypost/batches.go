package easypost

import "time"

type Batches struct {
	Id            string            `json:"id,omitempty"`            //Unique, begins with "batch_"
	Reference     string            `json:"reference,omitempty"`     //An optional field that may be used in place of ID in some API endpoints
	Object        string            `json:"object,omitempty"`        //"Batch"
	Mode          string            `json:"mode,omitempty"`          //"test" or "production"
	State         string            `json:"state,omitempty"`         //The overall state. Possible values are "creating", "creation_failed", "created", "purchasing", "purchase_failed", "purchased", "label_generating", and "label_generated"
	Num_Shipments int               `json:"num_shipments,omitempty"` //The number of shipments added
	Shipments     []*BatchShipment  `json:"shipments,omitempty"`     //
	Status        map[string]string `json:"status,omitempty"`        //A map of BatchShipment statuses to the count of BatchShipments with that status. Valid statuses are "postage_purchased", "postage_purchase_failed", "queued_for_purchase", and "creation_failed"
	Label_Url     string            `json:"label_url,omitempty"`     //The label image url
	Scan_Form     ScanForm          `json:"scan_form,omitempty"`     //The created ScanForm
	Pickup        Pickup            `json:"pickup,omitempty"`        //The created Pickup
	Created_At    time.Time         `json:"created_at,omitempty"`    //
	Updated_At    time.Time         `json:"updated_at,omitempty"`    //

}

type BatchShipment struct {
	Id            string `json:"id,omitempty"`            //The id of the Shipment. Unique, begins with "shp_"
	Reference     string `json:"reference,omitempty"`     //An optional field that may be used in place of ID in some API endpoints
	Batch_Status  string `json:"batch_status,omitempty"`  //The current status. Possible values are "postage_purchased", "postage_purchase_failed", "queued_for_purchase", and "creation_failed"
	Batch_Message string `json:"batch_message,omitempty"` //A human readable message for any errors that occurred during the Batch's life cycle

}
