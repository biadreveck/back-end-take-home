package model

type Connection struct {
	Airline     *Airline `json:"airline"`
	Origin      *Airport `json:"origin"`
	Destination *Airport `json:"destination"`
}
