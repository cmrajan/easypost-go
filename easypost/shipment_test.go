package easypost

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

var TestShipmentID string

func TestShipment(t *testing.T) {

	shipment := api.NewShipment()

	shipment.To_Address = &Address{Name: "Dr. Steve Brule", Street1: "179 N Harbor Dr", City: "Redondo Beach", State: "CA",
		Zip: "90277", Country: "US", Phone: "8573875756"}
	shipment.From_Address = &Address{Name: "EasyPost", Street1: "417 Montgomery Street", Street2: "5th Floor", City: "San Francisco", State: "CA",
		Zip: "94104", Country: "US", Phone: "4153334445", Email: "support@easypost.com"}
	shipment.Parcel = &Parcel{Length: 20.2, Width: 10.9, Height: 5, Weight: 65.9}

	//   -d 'shipment[customs_info][id]=cstinfo_...'

	err := shipment.Save()
	if err != nil {
		log.WithFields(logrus.Fields{}).Debugf("Error saving insurance in api_test: %s", err)
	}
	log.WithFields(logrus.Fields{}).Debugf("New insurance ID Created: %s\n", shipment.Id)

	TestShipmentID = shipment.Id

}

// func TestGetInsurances(t *testing.T) {
// 	var insurance *Insurances
// 	var err error
// 	insurance, err = api.Insurances("zzz", "aaa", time.Now(), time.Now(), 1)

// 	if err != nil {
// 		log.WithFields(logrus.Fields{}).Debugf("Error saving insurance in api_test: %s", err)
// 	}
// 	// log.WithFields(logrus.Fields{}).Debugf("New insurance ID Created: %s\n", insurance.ID)

// 	log.WithFields(logrus.Fields{}).Debugf("============Insurances: %+v\n", insurance)

// }

func TestGetShipment(t *testing.T) {
	var shipment *Shipment
	var err error
	shipment, err = api.Shipment(TestShipmentID)

	if err != nil {
		log.WithFields(logrus.Fields{}).Debugf("Error saving shipment in api_test: %s", err)
	}

	log.WithFields(logrus.Fields{}).Debugf("============Insurances: %+v\n", shipment)

}

func TestShipmentBuy(t *testing.T) {

	shipment := api.NewShipment()

	shipment.Rates = &Rate{}
	shipment.Insurance = "100.00"

	//   -d 'shipment[customs_info][id]=cstinfo_...'

	err := shipment.Save()
	if err != nil {
		log.WithFields(logrus.Fields{}).Debugf("Error saving insurance in api_test: %s", err)
	}
	log.WithFields(logrus.Fields{}).Debugf("New insurance ID Created: %s\n", shipment.Id)

	TestShipmentID = shipment.Id

}
