package easypost

import "time"

type CustomsItem struct {
	Id               string    `json:"id,omitempty"`               //Unique, begins with 'cstitem_'
	Object           string    `json:"object,omitempty"`           //'CustomsItem'
	Description      string    `json:"description,omitempty"`      //Required, description of item being shipped
	Quantity         float32   `json:"quantity,omitempty"`         //Required, greater than zero
	Value            float32   `json:"value,omitempty"`            //Required, greater than zero, total value (unit value * quantity) USD
	Weight           float32   `json:"weight,omitempty"`           //Required, greater than zero, total weight (unit weight * quantity) oz
	Hs_Tariff_Number string    `json:"hs_tariff_number,omitempty"` //Harmonized Tariff Schedule, e.g. "6109.10.0012" for Men's T-shirts
	Code             string    `json:"code,omitempty"`             //SKU/UPC or other product identifier
	Origin_Country   string    `json:"origin_country,omitempty"`   //Required, 2 char country code
	Currency         string    `json:"currency,omitempty"`         //3 char currency code, default USD
	Created_At       time.Time `json:"created_at,omitempty"`       //
	Updated_At       time.Time `json:"updated_at,omitempty"`       //

}
