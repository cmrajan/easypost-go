package easypost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
)

type Parcel struct {
	Id                 string    `json:"id,omitempty"`                 //	Unique, begins with "prcl_"
	Object             string    `json:"object,omitempty"`             //	"Parcel"
	Mode               string    `json:"mode,omitempty"`               //	"test" or "production"
	Length             float32   `json:"length"`                       //	inches Required if width and/or height are present
	Width              float32   `json:"width"`                        //	inches Required if length and/or height are present
	Height             float32   `json:"height"`                       //	inches Required if length and/or width are present
	Predefined_Package string    `json:"predefined_package,omitempty"` //	Optional, one of our predefined_packages
	Weight             float32   `json:"weight"`                       //oz	Always required
	Created_At         time.Time `json:"created_at,omitempty"`         //
	Updated_At         time.Time `json:"updated_at,omitempty"`         //
	api                *API      `json:"-"`
}

func (api *API) Parcel(id string) (*Parcel, error) {
	endpoint := fmt.Sprintf("parcels/%s", id)
	res, status, err := api.request(endpoint, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, fmt.Errorf("Status returned: %d", status)
	}
	r := Parcel{}
	err = json.NewDecoder(res).Decode(&r)

	result := r

	if err != nil {
		return nil, err
	}

	result.api = api
	return &result, nil
}

func (api *API) NewParcel() *Parcel {
	return &Parcel{api: api}
}

func (obj *Parcel) Save() error {
	endpoint := "parcels"

	// body := map[string]*Parcel{}
	// body = obj
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(obj)
	// fmt.Printf("passed data is %+v\n", buf)

	if err != nil {

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
			return err
		}
		return err

	}

	r := Parcel{}
	err = json.NewDecoder(res).Decode(&r)
	if err != nil {
		return err
	}
	*obj = r

	return nil
}
