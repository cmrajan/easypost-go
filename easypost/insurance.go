package easypost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
)

type Insurances struct {
	Insurances []*Insurance `json:"insurances"`
	api        *API
}
type Insurance struct {
	ID           string    `json:"id,omitempty"`            //Unique identifier, begins with "ins_"
	Object       string    `json:"object,omitempty"`        //"Insurance"
	Mode         string    `json:"mode,omitempty"`          //"test" or "production"
	Reference    string    `json:"reference,omitempty"`     //The unique reference for this Insurance, if any
	Amount       string    `json:"amount,omitempty"`        //USD value of insured goods with sub-cent precision
	Provider     string    `json:"provider,omitempty"`      //The insurance provider used by EasyPost
	ProviderId   string    `json:"provider_id,omitempty"`   //An identifying number for some insurance providers used by EasyPost
	ShipmentId   string    `json:"shipment_id,omitempty"`   //The ID of the Shipment in EasyPost, if postage was purchased via EasyPost
	TrackingCode string    `json:"tracking_code,omitempty"` //The tracking code of either the shipment within EasyPost, or provided by you during creation
	Status       string    `json:"status,omitempty"`        //The current status of the insurance, possible values are "new", "pending", "purchased", "failed", or "cancelled"
	Tracker      *Tracker  `json:"tracker,omitempty"`       //The associated Tracker object
	ToAddress    *Address  `json:"to_address,omitempty"`    //The associated Address object for destination
	FromAddress  *Address  `json:"from_address,omitempty"`  //The associated Address object for origin
	Fee          Fee       `json:"fee,omitempty"`           //The associated InsuranceFee object if any
	Messages     []string  `json:"messages,omitempty"`      //The list of errors encountered during attempted purchase of the insurance
	CreatedAt    time.Time `json:"created_at,omitempty"`    //
	UpdatedAt    time.Time `json:"updated_at,omitempty"`    //
	api          *API
}

func (api *API) Insurances(before_id string, after_id string, start_datetime time.Time, end_datetime time.Time, page_size int) (*Insurances, error) {
	endpoint := "/insurances"

	params := make(map[string]interface{})
	params["before_id"] = before_id
	params["after_id"] = after_id
	params["start_datetime"] = start_datetime
	params["end_datetime"] = end_datetime
	params["page_size"] = page_size

	res, status, err := api.request(endpoint, "GET", params, nil)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("Status returned: %d", status)
	}

	r := Insurances{}
	err = json.NewDecoder(res).Decode(&r)

	result := r

	if err != nil {
		return nil, err
	}

	result.api = api

	return &result, nil

}

func (api *API) Insurance(id string) (*Insurance, error) {
	endpoint := fmt.Sprintf("/insurances/%s", id)

	res, status, err := api.request(endpoint, "GET", nil, nil)

	if err != nil {
		return nil, err
	}

	if status != 200 {

		return nil, fmt.Errorf("Status returned: %d", status)
	}

	r := Insurance{}
	err = json.NewDecoder(res).Decode(&r)

	result := r

	if err != nil {
		return nil, err
	}

	result.api = api

	return &result, nil
}

func (api *API) NewInsurance() *Insurance {
	return &Insurance{api: api}
}

func (obj *Insurance) Save() error {
	endpoint := "insurances"

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

	r := Insurance{}
	err = json.NewDecoder(res).Decode(&r)
	if err != nil {

		return err
	}
	*obj = r

	return nil
}
