package models

// Car is the type on which all the core layer's functionality is implemented on
type Car struct {
	Name string `json:"name"`
	Price  int    `json:"Price"`
	model string `json:"model"`
}
