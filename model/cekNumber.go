package model

// model number
type PhoneNumber struct {
	ID    int    `json:"id,omitempty"`
	Phone string `json:"phone"`
}

// untuk request
type PhoneRequest struct {
	Phones []PhoneNumber `json:"phones"`
}
