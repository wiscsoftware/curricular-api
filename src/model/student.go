package model

import (
	"fmt"
	"time"
)

type Identity struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Address struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zipCode"`
}

type ResidencyStatus int

func (r ResidencyStatus) String() string {
	switch r {
	case Resident:
		return "Resident"
	case NonResident:
		return "NonResident"
	case Undermined:
		return "Undermined"
	default:
		return "Unknown"
	}
}

const (
	Resident ResidencyStatus = iota
	NonResident
	Undermined
)

type Student struct {
	Identities []Identity      `json:"identities"`
	FirstName  string          `json:"firstName"`
	LastName   string          `json:"lastName"`
	Addresses  []Address       `json:"addresses"`
	Birthdate  time.Time       `json:"birthdate"`
	AgeInYears int             `json:"ageInYears"`
	Residency  ResidencyStatus `json:"residencyStatus"`
}

func (s Student) Print() {
	fmt.Printf("student: %+v\n", s)
}
