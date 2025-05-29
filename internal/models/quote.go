package models

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
