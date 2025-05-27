package model

import "encoding/json"

type Term int

func (t Term) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t Term) String() string {
	switch t {
	case Spring:
		return "Spring"
	case Summer:
		return "Summer"
	case Fall:
		return "Fall"
	default:
		return "Unknown Term"
	}
}

const (
	Spring Term = iota
	Summer
	Fall
)

type Course struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Term     Term      `json:"term"`
	Credit   Credit    `json:"credit"`
	Teachers []Teacher `json:"teachers"`
}
