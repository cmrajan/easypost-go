package easypost

import "time"

type Shipment struct {
	Id             string        `json:"id,omitempty"`             //Unique, begins with "shp_"
	Object         string        `json:"object,omitempty"`         //"Shipment"
	Reference      string        `json:"reference,omitempty"`      //An optional field that may be used in place of id in other API endpoints
	Mode           string        `json:"mode,omitempty"`           //"test" or "production"
	To_Address     *Address      `json:"to_address,omitempty"`     //The destination address
	From_Address   *Address      `json:"from_address,omitempty"`   //The origin address
	Return_Address *Address      `json:"return_address,omitempty"` //The shipper's address, defaults to from_address
	Buyer_Address  *Address      `json:"buyer_address,omitempty"`  //The buyer's address, defaults to to_address
	Parcel         *Parcel       `json:"parcel,omitempty"`         //The dimensions and weight of the package
	Customs_Info   *CustomsInfo  `json:"customs_info,omitempty"`   //Information for the processing of customs
	Scan_Form      *ScanForm     `json:"scan_form,omitempty"`      //Document created to manifest and scan multiple shipments
	Forms          []*Form       `json:"forms,omitempty"`          //All associated Form objects
	Insurance      *Insurance    `json:"insurance,omitempty"`      //The associated Insurance object
	Rates          []*Rate       `json:"rates,omitempty"`          //All associated Rate objects
	Selected_Rate  *Rate         `json:"selected_rate,omitempty"`  //The specific rate purchased for the shipment, or null if unpurchased or purchased through another mechanism
	Postage_Label  *PostageLabel `json:"postage_label,omitempty"`  //The associated PostageLabel object
	Messages       []*Message    `json:"messages,omitempty"`       //Any carrier errors encountered during rating, discussed in more depth below
	Options        *Options      `json:"options,omitempty"`        //All of the options passed to the shipment, discussed in more depth below
	Is_Return      bool          `json:"is_return,omitempty"`      //Set true to create as a return, discussed in more depth below
	Tracking_Code  string        `json:"tracking_code,omitempty"`  //If purchased, the tracking code will appear here as well as within the Tracker object
	Usps_Zone      string        `json:"usps_zone,omitempty"`      //The USPS zone of the shipment, if purchased with USPS
	Status         string        `json:"status,omitempty"`         //The current tracking status of the shipment
	Tracker        *Tracker      `json:"tracker,omitempty"`        //The associated Tracker object
	Fees           []*Fee        `json:"fees,omitempty"`           //The associated Fee objects charged to the billing user account
	Refund_Status  string        `json:"refund_status,omitempty"`  //The current status of the shipment refund process. Possible values are "submitted", "refunded", "rejected".
	Batch_Id       string        `json:"batch_id,omitempty"`       //The ID of the batch that contains this shipment, if any
	Batch_Status   string        `json:"batch_status,omitempty"`   //The current state of the associated BatchShipment
	Batch_Message  string        `json:"batch_message,omitempty"`  //The current message of the associated BatchShipment
	Created_At     time.Time     `json:"created_at,omitempty"`     //
	Updated_At     time.Time     `json:"updated_at,omitempty"`     //
}

type Options struct {
	Additional_Handling          bool    `json:"additional_handling,omitempty"`          //Setting this option to true, will add an additional handling charge. An
	Address_Validation_Level     string  `json:"address_validation_level,omitempty"`     //Setting this option to "0", will allow the minimum amount of address information to pass the validation check. Only for USPS postage.
	Alcohol                      bool    `json:"alcohol,omitempty"`                      //Set this option to true if your shipment contains alcohol.
	Bill_Receiver_Account        string  `json:"bill_receiver_account,omitempty"`        //Setting an account number of the receiver who is to receive and buy the postage.
	Bill_Receiver_Postal_Code    string  `json:"bill_receiver_postal_code,omitempty"`    //Setting a postal code of the receiver account you want to buy postage.
	Bill_Third_Party_Account     string  `json:"bill_third_party_account,omitempty"`     //Setting an account number of the third party account you want to buy postage.
	Bill_Third_Party_Country     string  `json:"bill_third_party_country,omitempty"`     //Setting a country of the third party account you want to buy postage.
	Bill_Third_Party_Postal_Code string  `json:"bill_third_party_postal_code,omitempty"` //Setting a postal code of the third party account you want to buy postage.
	By_Drone                     bool    `json:"by_drone,omitempty"`                     //Setting this option to true will indicate to the carrier to prefer delivery by drone, if the carrier supports drone delivery.
	Carbon_Neutral               bool    `json:"carbon_neutral,omitempty"`               //Setting this to true will add a charge to reduce carbon emissions.
	Cod_Amount                   string  `json:"cod_amount,omitempty"`                   //Adding an amount will have the carrier collect the specified amount from the recipient.
	Cod_Method                   string  `json:"cod_method,omitempty"`                   //Method for payment. "CASH", "CHECK", "MONEY_ORDER"
	Currency                     string  `json:"currency,omitempty"`                     //Which currency this shipment will show for rates if carrier allows.
	Delivery_Confirmation        string  `json:"delivery_confirmation,omitempty"`        //If you want to request a signature, you can pass "ADULT_SIGNATURE" or
	Dry_Ice                      bool    `json:"dry_ice,omitempty"`                      //Package contents contain dry ice.
	Dry_Ice_Medical              string  `json:"dry_ice_medical,omitempty"`              //If the dry ice is for medical use, set this option to true.
	Dry_Ice_Weight               string  `json:"dry_ice_weight,omitempty"`               //Weight of the dry ice in ounces.
	Endorsement                  string  `json:"endorsement,omitempty"`                  //Possible values "ADDRESS_SERVICE_REQUESTED",
	Freight_Charge               float32 `json:"freight_charge,omitempty"`               //Additional cost to be added to the invoice of this shipment. Only applies to UPS currently.
	Handling_Instructions        string  `json:"handling_instructions,omitempty"`        //This is to designate special instructions for the carrier like "Do not drop!".
	Hazmat                       string  `json:"hazmat,omitempty"`                       //Dangerous goods indicator. Possible values are "ORMD" and "LIMITED_QUANTITY". Applies to USPS, FedEx and DHL eCommerce.
	Hold_For_Pickup              bool    `json:"hold_for_pickup,omitempty"`              //Package will wait at carrier facility for pickup.
	Incoterm                     string  `json:"incoterm,omitempty"`                     //Incoterm negotiated for shipment. Supported values are "EXW", "FCA", "CPT",
	Invoice_Number               string  `json:"invoice_number,omitempty"`               //This will print an invoice number on the postage label.
	Label_Date                   string  `json:"label_date,omitempty"`                   //Set the date that will appear on the postage label. Accepts ISO 8601 formatted string including time zone offset. EasyPost stores all dates as UTC time.
	Label_Format                 string  `json:"label_format,omitempty"`                 //Supported label formats include "PNG", "PDF", "ZPL", and "EPL2". "PNG" is the only format that allows for conversion.
	Machinable                   bool    `json:"machinable,omitempty"`                   //Whether or not the parcel can be processed by the carriers equipment.
	Print_Custom_1               string  `json:"print_custom_1,omitempty"`               //You can optionally print custom messages on labels. The locations of these fields show up on different spots on the carrier's labels.
	Print_Custom_2               string  `json:"print_custom_2,omitempty"`               //An additional message on the label. Same restrictions as print_custom_1
	Print_Custom_3               string  `json:"print_custom_3,omitempty"`               //An additional message on the label. Same restrictions as print_custom_1
	Print_Custom_1_Barcode       bool    `json:"print_custom_1_barcode,omitempty"`       //Create a barcode for this custom reference if supported by carrier.
	Print_Custom_2_Barcode       bool    `json:"print_custom_2_barcode,omitempty"`       //Create a barcode for this custom reference if supported by carrier.
	Print_Custom_3_Barcode       bool    `json:"print_custom_3_barcode,omitempty"`       //Create a barcode for this custom reference if supported by carrier.
	Print_Custom_1_Code          string  `json:"print_custom_1_code,omitempty"`          //Specify the type of print_custom_1.
	Print_Custom_2_Code          string  `json:"print_custom_2_code,omitempty"`          //See print_custom_1_code.
	Print_Custom_3_Code          string  `json:"print_custom_3_code,omitempty"`          //See print_custom_1_code.
	Saturday_Delivery            bool    `json:"saturday_delivery,omitempty"`            //Set this value to true for delivery on Saturday.
	Special_Rates_Eligibility    string  `json:"special_rates_eligibility,omitempty"`    //This option allows you to request restrictive rates from USPS. Can set to 'USPS.MEDIAMAIL' or 'USPS.LIBRARYMAIL'.
	Smartpost_Hub                string  `json:"smartpost_hub,omitempty"`                //You can use this to override the hub ID you have on your account.
	Smartpost_Manifest           string  `json:"smartpost_manifest,omitempty"`           //The manifest ID is used to group SmartPost packages onto a manifest for each trailer.

}

type Rates struct {
	Id                       string    `json:"id,omitempty"`                       //unique, begins with 'rate_'
	Object                   string    `json:"object,omitempty"`                   //"Rate"
	Mode                     string    `json:"mode,omitempty"`                     //"test" or "production"
	Service                  string    `json:"service,omitempty"`                  //service level/name
	Carrier                  string    `json:"carrier,omitempty"`                  //name of carrier
	Carrier_Account_Id       string    `json:"carrier_account_id,omitempty"`       //ID of the CarrierAccount record used to generate this rate
	Shipment_Id              string    `json:"shipment_id,omitempty"`              //ID of the Shipment this rate belongs to
	Rate                     string    `json:"rate,omitempty"`                     //the actual rate quote for this service
	Currency                 string    `json:"currency,omitempty"`                 //currency for the rate
	Retail_Rate              string    `json:"retail_rate,omitempty"`              //the retail rate is the in-store rate given with no account
	Retail_Currency          string    `json:"retail_currency,omitempty"`          //currency for the retail rate
	List_Rate                string    `json:"list_rate,omitempty"`                //the list rate is the non-negotiated rate given for having an account with the carrier
	List_Currency            string    `json:"list_currency,omitempty"`            //currency for the list rate
	Delivery_Days            int       `json:"delivery_days,omitempty"`            //delivery days for this service
	Delivery_Date            string    `json:"delivery_date,omitempty"`            //date for delivery
	Delivery_Date_Guaranteed bool      `json:"delivery_date_guaranteed,omitempty"` //indicates if delivery window is guaranteed (true) or not (false)
	Est_Delivery_Days        int       `json:"est_delivery_days*,omitempty"`       //*This field is deprecated and should be ignored.
	Created_At               time.Time `json:"created_at,omitempty"`               //
	Updated_At               time.Time `json:"updated_at,omitempty"`               //
}

type Messages struct {
	Carrier string `json:"carrier,omitempty"` //the name of the carrier generating the error, e.g. "UPS"
	Type    string `json:"type,omitempty"`    //the category of error that occurred. Most frequently "rate_error"
	Message string `json:"message,omitempty"` //the string from the carrier explaining the problem.

}
