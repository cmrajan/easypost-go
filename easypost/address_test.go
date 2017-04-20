package easypost

import (
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestNewAddress(t *testing.T) {
	fmt.Println("\nCreating New Address.............")
	if !remoteEnabled {
		fmt.Println("Env not set")
		return
	}

	address := api.NewAddress()
	address.Verify = []string{"delivery"}
	address.Street1 = "5050 Hacienda Drive"
	address.Street2 = "1112"
	address.City = "DUBLIN"
	address.State = "CA"
	address.Zip = "94568"
	address.Country = "US"

	address.Phone = "415-123-4567"

	err := address.Save()
	if err != nil {
		log.WithFields(logrus.Fields{}).Debugf("Error saving parcel in api_test: %s", err)
	}
	log.WithFields(logrus.Fields{}).Debugf("New address ID Created: %s\n", address.ID)
	verifications := address.Verifications
	log.WithFields(logrus.Fields{}).Debugf("address verification is %+v", verifications)

	TestAddressID1 = address.ID
}
func TestNewAddress2(t *testing.T) {
	fmt.Println("\nCreating New Address2.............")
	if !remoteEnabled {
		fmt.Println("Env not set")
		return
	}

	address := api.NewAddress()

	address.Verify = []string{"delivery"}

	address.Street1 = "417 MONTGOMERY ST"
	address.Street2 = "FLOOR 5"
	address.City = "SAN FRANCISCO"
	address.State = "CA"
	address.Zip = "94104"
	address.Country = "US"
	address.Company = "EasyPost"
	address.Phone = "415-123-4567"

	err := address.Save()
	if err != nil {
		log.WithFields(logrus.Fields{}).Debugf("Error saving parcel in api_test: %s", err)
	}
	log.WithFields(logrus.Fields{}).Debugf("New address ID Created: %s\n", address.ID)
	verifications := address.Verifications
	log.WithFields(logrus.Fields{}).Debugf("address verification is %+v", verifications)

	TestAddressID2 = address.ID
}
func TestNewAddress3(t *testing.T) {
	fmt.Println("\nCreating New Address3.............")
	if !remoteEnabled {
		fmt.Println("Env not set")
		return
	}

	address := api.NewAddress()

	address.Street1 = "UNDELIVERABLE ST"

	address.City = "SAN FRANCISCO"
	address.State = "CA"
	address.Zip = "94104"
	address.Country = "US"
	address.Company = "EasyPost"
	address.Phone = "415-123-4567"

	err := address.Save()
	if err != nil {
		log.WithFields(logrus.Fields{}).Debugf("Error saving parcel in api_test: %s", err)
	}
	log.WithFields(logrus.Fields{}).Debugf("New address ID Created: %s\n", address.ID)

	TestAddressID3 = address.ID
}

func TestGetAddress(t *testing.T) {

	if !remoteEnabled {
		return
	}

	address, err := api.Address(TestAddressID1)

	if err != nil {
		t.Errorf("Error fetching address: %v", err)
	}
	verifications := address.Verifications

	log.WithFields(logrus.Fields{}).Debugf("Fetched address is %+v", address)
	log.WithFields(logrus.Fields{}).Debugf("address verification is %+v", verifications)

}
