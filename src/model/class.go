package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Credit struct {
	credit float64 // provide utility methods without exposing basic type
}

func (c Credit) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c Credit) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &c.credit)
}

func (c Credit) String() string {
	return fmt.Sprintf("%f", c.credit)
}

func NewCredit(credit float64) Credit {
	return Credit{credit}
}

type Class struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Credit   Credit    `json:"credit"`
	Location Address   `json:"location"`
	Time     time.Time `json:"time"`
}
