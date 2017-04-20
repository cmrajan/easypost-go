package easypost

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
)

type Address struct {
	ID              string         `json:"id,omitempty"`               //Unique identifier, begins with "adr_"
	Object          string         `json:"object,omitempty"`           //"Address"
	Mode            string         `json:"mode,omitempty"`             //Set based on which api-key you used, either "test" or "production"
	Street1         string         `json:"street1,omitempty"`          //First line of the address
	Street2         string         `json:"street2,omitempty"`          //Second line of the address
	City            string         `json:"city,omitempty"`             //City the address is located in
	State           string         `json:"state,omitempty"`            //State or province the address is located in
	Zip             string         `json:"zip,omitempty"`              //ZIP or postal code the address is located in
	Country         string         `json:"country,omitempty"`          //ISO 3166 country code for the country the address is located in
	Residential     bool           `json:"residential,omitempty"`      //Whether or not this address would be considered residential
	CarrierFacility string         `json:"carrier_facility,omitempty"` //The specific designation for the address (only relevant if the address is a carrier facility)
	Name            string         `json:"name,omitempty"`             //Name of the person. Both name and company can be included
	Company         string         `json:"company,omitempty"`          //Name of the organization. Both name and company can be included
	Phone           string         `json:"phone,omitempty"`            //Phone number to reach the person or organization
	Email           string         `json:"email,omitempty"`            //Email to reach the person or organization
	FederalTaxID    string         `json:"federal_tax_id,omitempty"`   //Federal tax identifier of the person or organization
	StateTaxID      string         `json:"state_tax_id,omitempty"`     //State tax identifier of the person or organization
	Verifications   *Verifications `json:"verifications,omitempty"`    //The result of any verifications requested
	Verify          []string       `json:"verify,omitempty"`
	VerifyStrict    []string       `json:"verify_strict,omitempty"`
	api             *API           `json:"-"`
}

type Verifications struct {
	Success bool                 `json:"sucess"`  //The success of the verification
	Errors  []*FieldError        `json:"errors"`  //FieldError array	//All errors that caused the verification to fail
	Details *VerificationDetails `json:"details"` //Extra data related to the verification

}
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type VerificationDetails struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (api *API) Address(id string) (*Address, error) {
	endpoint := fmt.Sprintf("/addresses/%s", id)

	res, status, err := api.request(endpoint, "GET", nil, nil)

	if err != nil {
		return nil, err
	}

	if status != 200 {

		return nil, fmt.Errorf("Status returned: %d", status)
	}

	r := Address{}
	err = json.NewDecoder(res).Decode(&r)

	result := r

	if err != nil {
		return nil, err
	}

	result.api = api

	return &result, nil
}

func (api *API) NewAddress() *Address {
	return &Address{api: api}
}

func (obj *Address) Save() error {
	endpoint := "addresses"

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(obj)

	log.WithFields(logrus.Fields{}).Debug(buf)

	if err != nil {
		log.WithFields(logrus.Fields{}).Debug("Error in JSON")
		return err
	}
	res, status, err := obj.api.request(endpoint, "POST", nil, buf)

	if err != nil {
		return err
	}

	if status != 201 { //not created status

		r := errorResponse{}
		err = json.NewDecoder(res).Decode(&r)
		if err == nil {
			log.WithFields(logrus.Fields{}).Debugf("Status %d: %+v", status, r.Errors)

			return fmt.Errorf("Status %d: %+v", status, r.Errors)
		}
		return err

	}

	r := Address{}
	err = json.NewDecoder(res).Decode(&r)
	if err != nil {

		return err
	}
	*obj = r

	return nil
}
