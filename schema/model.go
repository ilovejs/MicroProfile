package schema

import "time"

// for create
type ProfileReq struct {
	Name     string `json:"name" validate:"required"`
	Gender   bool   `json:"gender"`
	DoB      string `json:"dob"`
	PostCode int    `json:"postcode"`
	PhoneNo  string `json:"phone_no"`
}

type ProfileResp struct {
	ID int `json:"id"`
}

// for select
type Profile struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Gender   bool      `json:"gender"`
	DoB      time.Time `json:"dob"`
	PostCode int       `json:"postcode"`
	PhoneNo  string    `json:"phone_no"`
}

type UpdateResponse struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Gender   bool      `json:"gender"`
	DoB      time.Time `json:"dob"`
	PostCode int       `json:"postcode"`
	PhoneNo  string    `json:"phone_no"`
}
