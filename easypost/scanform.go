package easypost

import "time"

type ScanForm struct {
	Id             string    `json:"id,omitempty"`             //Unique, begins with "sf_"
	Object         string    `json:"object,omitempty"`         //"ScanForm"
	Status         string    `json:"status,omitempty"`         //Current status. Possible values are "creating", "created" and "failed"
	Message        string    `json:"message,omitempty"`        //Human readable message explaining any failures
	Address        *Address  `json:"address,omitempty"`        //Address the will be Shipments shipped from
	Tracking_Codes []string  `json:"tracking_codes,omitempty"` //Tracking codes included on the ScanForm
	Form_Url       string    `json:"form_url,omitempty"`       //Url of the document
	Form_File_Type string    `json:"form_file_type,omitempty"` //File format of the document
	Batch_Id       string    `json:"batch_id,omitempty"`       //The id of the associated Batch. Unique, starts with "batch_"
	Created_At     time.Time `json:"created_at,omitempty"`     //
	Updated_At     time.Time `json:"updated_at,omitempty"`     //
}
