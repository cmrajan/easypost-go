package easypost

type CarrierType struct {
	Object string            `json:"object,omitempty"` //"CarrierType"
	Type   string            `json:"type,omitempty"`   //Specifies the CarrierAccount type.
	Fields *CarrierTypeField `json:"fields,omitempty"` //Contains at least one of the following keys: "auto_link", "credentials", "test_credentials", and "custom_workflow"

}

type CarrierTypeField struct {
	Auto_Link        bool                   `json:"auto_link,omitempty"`        //If this key is present with the value of true, no other attributes are needed for CarrierAccount creation
	Custom_Workflow  bool                   `json:"custom_workflow,omitempty"`  //If this key is present with the value of true, CarrierAccount creation of this type requires extra work not handled by the CarrierAccount standard API
	Credentials      *CarrierTypeCredential `json:"credentials,omitempty"`      //If this object is present, required attribute names and their metadata are presented inside
	Test_Credentials *CarrierTypeCredential `json:"test_credentials,omitempty"` //If this object is present, it contains attribute names and metadata just as the credentials object. It is not required for CarrierAccount creation if you do not plan on using the carrier account for test mode
}

type CarrierTypeCredential struct {
	Name       string `json:"name,omitempty"`       //The key for each attribute sub-object within credentials is the name of the attribute for submission on CarrierAccounts
	Visibility string `json:"visibility,omitempty"` //There are five possible values, which encode the security need and storage type for each attribute: "visible", "checkbox", "fake", "password", and "masked"
	Label      string `json:"label,omitempty"`      //Most attributes have generic names, so for clarity a "label" value is provided for clearer representation when rendering forms
}
