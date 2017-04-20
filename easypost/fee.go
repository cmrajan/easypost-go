package easypost

type Fee struct {
	Object   string `json:"object,omitempty"`   //"Fee"
	Type     string `json:"type,omitempty"`     //The name of the category of fee. Possible types are "LabelFee", "PostageFee", "InsuranceFee", and "TrackerFee"
	Amount   string `json:"amount,omitempty"`   //USD value with sub-cent precision
	Charged  bool   `json:"charged,omitempty"`  //Whether EasyPost has successfully charged your account for the fee
	Refunded bool   `json:"refunded,omitempty"` //Whether the Fee has been refunded successfully

}
