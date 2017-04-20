package easypost

import (
	"os"
)

var api API
var remoteEnabled = false
var TestParcelID string
var TestAddressID1 string
var TestAddressID2 string
var TestAddressID3 string

var TestInsuranceID string

func init() {

	if os.Getenv("EASYPOST_TOKEN") != "" && os.Getenv("EASYPOST_HOST") != "" {
		remoteEnabled = true
		api = API{URI: os.Getenv("EASYPOST_HOST"),
			Token: os.Getenv("EASYPOST_API")}
	} else {
		log.Printf("Tests disabled. set Env variable EASYPOST_API and EASYPOST_HOST")

	}
}
