package easypost

type User struct {
	Id                        string  `json:"id,omitempty"`                        //Unique, begins with "user_"
	Object                    string  `json:"object,omitempty"`                    //"User"
	Parent_Id                 string  `json:"parent_id,omitempty"`                 //The ID of the parent user object. Top-level users are defined as users with no parent
	Name                      string  `json:"name,omitempty"`                      //First and last name required
	Email                     string  `json:"email,omitempty"`                     //Required
	Phone_Number              string  `json:"phone_number,omitempty"`              //Optional
	Balance                   string  `json:"balance,omitempty"`                   //Formatted as string "XX.XXXXX"
	Recharge_Amount           string  `json:"recharge_amount,omitempty"`           //USD formatted dollars and cents, delimited by a decimal point
	Secondary_Recharge_Amount string  `json:"secondary_recharge_amount,omitempty"` //USD formatted dollars and cents, delimited by a decimal point
	Recharge_Threshold        string  `json:"recharge_threshold,omitempty"`        //Number of cents USD that when your balance drops below, we automatically recharge your account with your primary payment method.
	Children                  []*User `json:"children,omitempty"`                  //All associated children

}
