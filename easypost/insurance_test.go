package easypost

import (
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
)

func TestInsurance(t *testing.T) {

	insurance := api.NewInsurance()
	insurance.FromAddress = &Address{ID: "adr_d318053cfe5746dfaeadbcd5c5e00c1a"}
	insurance.ToAddress = &Address{ID: "adr_5f3611f0ce9f423094164398896b2243"}
	insurance.TrackingCode = "9400110898825022579493"
	insurance.Reference = "insuranceRef1"
	insurance.Amount = "$100.00"

	err := insurance.Save()
	if err != nil {
		log.WithFields(logrus.Fields{}).Debugf("Error saving insurance in api_test: %s", err)
	}
	log.WithFields(logrus.Fields{}).Debugf("New insurance ID Created: %s\n", insurance.ID)

	TestInsuranceID = insurance.ID

}

func TestGetInsurances(t *testing.T) {
	var insurance *Insurances
	var err error
	insurance, err = api.Insurances("zzz", "aaa", time.Now(), time.Now(), 1)

	if err != nil {
		log.WithFields(logrus.Fields{}).Debugf("Error saving insurance in api_test: %s", err)
	}
	// log.WithFields(logrus.Fields{}).Debugf("New insurance ID Created: %s\n", insurance.ID)

	log.WithFields(logrus.Fields{}).Debugf("============Insurances: %+v\n", insurance)

}
func TestGetInsurance(t *testing.T) {
	var insurance *Insurance
	var err error
	insurance, err = api.Insurance(TestInsuranceID)

	if err != nil {
		log.WithFields(logrus.Fields{}).Debugf("Error saving insurance in api_test: %s", err)
	}

	log.WithFields(logrus.Fields{}).Debugf("============Insurances: %+v\n", insurance)

}
