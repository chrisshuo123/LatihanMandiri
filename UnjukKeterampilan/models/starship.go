package models

type Starship struct {
	Id                   uint   `gorm:"primary key" json:"id" form:"id"`
	Name                 string `gorm:"not null" json:"name"`
	Model                string `json:"model"`
	Manufacturer         string `json:"manufacturer"`
	CostInCredits        int    `json:"cost_in_credits"`
	Length               int    `json:"length"` //float64
	MaxAtmospheringSpeed int    `json:"max_atmosphering_speed"`
	Crew                 string `json:"crew"`
	Passengers           int    `json:"passengers"`     // int64
	CargoCapacity        int    `json:"cargo_capacity"` // int64
	Consumables          string `json:"consumables"`
	HyperdriveRating     int    `json:"hyperdrive_rating"` // float64
	MGLT                 int    `json:"mglt"`
	StarshipClass        string `json:"starship_class"`
}
