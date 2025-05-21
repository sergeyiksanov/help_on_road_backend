package models

type HelpCallStatus int

const (
	Pending  HelpCallStatus = iota
	Helping  HelpCallStatus = iota
	Helped   HelpCallStatus = iota
	Rejected HelpCallStatus = iota
)

type HelpCall struct {
	Service     string         `json:"service,omitempty"`
	Latitude    float64        `json:"latitude,omitempty"`
	Longitude   float64        `json:"longitude,omitempty"`
	Description string         `json:"description"`
	Caller      *User          `json:"caller,omitempty"`
	Status      HelpCallStatus `json:"status"`
	PayType     string         `json:"pay_type"`
}
