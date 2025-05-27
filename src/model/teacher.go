package model

type Teacher struct {
	Identities []Identity `json:"identities"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
}
