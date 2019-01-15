package DTO


type Customer struct {
	 Id int32 `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`

}
type CustomerList []Customer



