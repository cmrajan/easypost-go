package easypost

type CarrierAccount struct {
	Id               string                `json:"id,omitempty"`               //Unique, begins with "ca_"
	Object           string                `json:"object,omitempty"`           //"CarrierAccount"
	Type             string                `json:"type,omitempty"`             //The name of the carrier type.
	Fields           *CarrierAccountFields `json:"fields,omitempty"`           //Contains "credentials" and/or "test_credentials", or may be empty
	Clone            bool                  `json:"clone,omitempty"`            //If clone is true, only the reference and description are possible to update
	Description      string                `json:"description,omitempty"`      //An optional, user-readable field to help distinguish accounts
	Reference        string                `json:"reference,omitempty"`        //An optional field that may be used in place of carrier_account_id in other API endpoints
	Readable         string                `json:"readable,omitempty"`         //The name used when displaying a readable value for the type of the account
	Credentials      object                `json:"credentials,omitempty"`      //Unlike the "credentials" object contained in "fields", this nullable object contains just raw credential pairs for client library consumption
	Test_Credentials object                `json:"test_credentials,omitempty"` //Unlike the "test_credentials" object contained in "fields", this nullable object contains just raw test_credential pairs for client library consumption
	Created_At       datetime              `json:"created_at,omitempty"`       //
	Updated_At       datetime              `json:"updated_at,omitempty"`       //

}

type CarrierAccountFields struct {
	Credentials      *CarrierAccountField `json:"credentials,omitempty"`      //Credentials used in the production environment.
	Test_Credentials *CarrierAccountField `json:"test_credentials,omitempty"` //Credentials used in the test environment.
	Auto_Link        boolean              `json:"auto_link,omitempty"`        //For USPS this designates that no credentials are required.
	Custom_Workflow  boolean              `json:"custom_workflow,omitempty"`  //When present, a seperate authentication process will be required through the UI to link this account type.
}

type CarrierAccountField struct {
	Key        string `json:"key,omitempty"`        //Each key in the sub-objects of a CarrierAccount's fields is the name of a settable field
	Visibility string `json:"visibility,omitempty"` //The visibility value is used to control form field types, and is discussed in the CarrierType section
	Label      string `json:"label,omitempty"`      //The label value is used in form rendering to display a more precise field name
	Value      string `json:"value,omitempty"`      //Checkbox fields use "0" and "1" as False and True, all other field types present plaintext, partly-masked, or masked credential data for reference
}
