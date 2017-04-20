package easypost

import "time"

type CustomsInfo struct {
	Id      string `json:"id,omitempty"`      //Unique, begins with 'cstinfo_'
	Object  string `json:"object,omitempty"`  //'CustomsInfo'
	Eel_Pfc string `json:"eel_pfc,omitempty"` //"EEL" or "PFC"
	//value less than $2500: "NOEEI 30.37(a)"; value greater than $2500: see Customs Guide
	ContentsType         string         `json:"contents_type,omitempty"`        //"documents", "gift", "merchandise", "returned_goods", "sample", or "other"
	ContentsExplanation  string         `json:"contents_explanation,omitempty"` //Human readable description of content. Required for certain carriers and always required if contents_type is "other"
	CustomsCertify       bool           `json:"customs_certify,omitempty"`      //Electronically certify the information provided
	CustomsSigner        string         `json:"customs_signer,omitempty"`       //Required if customs_certify is true
	Non_Delivery_Option  string         `json:"non_delivery_option,omitempty"`  //"abandon" or "return", defaults to "return"
	Restriction_Type     string         `json:"restriction_type,omitempty"`     //"none", "other", "quarantine", or "sanitary_phytosanitary_inspection"
	Restriction_Comments string         `json:"restriction_comments,omitempty"` //Required if restriction_type is not "none"
	CustomsItems         []*CustomsItem `json:"customs_items,omitempty"`        //Describes to products being shipped
	CreatedAt            time.Time      `json:"created_at,omitempty"`           //
	UpdatedAt            time.Time      `json:"updated_at,omitempty"`           //
}
