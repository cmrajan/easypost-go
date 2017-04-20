package easypost

import (
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestNewParcel(t *testing.T) {

	if !remoteEnabled {
		fmt.Println("Env not set")
		return
	}

	parcel := api.NewParcel()

	parcel.Length = 20.2
	parcel.Width = 10.9
	parcel.Height = 5
	parcel.Predefined_Package = "Parcel"
	parcel.Weight = 65.9

	err := parcel.Save()
	if err != nil {
		t.Errorf("Error saving parcel in api_test: %s", err)
	}
	fmt.Printf("New parcel ID Created: %s\n", parcel.Id)
	TestParcelID = parcel.Id
}

func TestParcel(t *testing.T) {

	if !remoteEnabled {
		return
	}

	parcel, err := api.Parcel(TestParcelID)

	if err != nil {
		t.Errorf("Error fetching parcels: %v", err)
	}
	log.WithFields(logrus.Fields{}).Debugf("Fetched Parcel is: %+v", parcel)

}
