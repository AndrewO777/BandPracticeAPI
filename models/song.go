package models

type Song struct {
	ID int `json:"id"`
	SongName string `json:"SongName"`
	BandName string `json:"BandName"`
	PlayCount uint16 `json:"PlayCount"`
	CurrentConfidence byte `json:"CurrentConfidence"`
	LastPlayed []uint8 `json:"LastPlayed"`
}
