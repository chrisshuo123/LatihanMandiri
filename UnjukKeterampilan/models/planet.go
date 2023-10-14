package models

type Planet struct {
	Id             uint   `gorm:"primary key" json:"id" form:"id"`
	Name           string `gorm:"not null" json:"name"`
	RotationPeriod int    `json:"rotation_period"`
	OrbitalPeriod  int    `json:"orbital_period"`
	Diameter       int    `json:"diameter"`
	Climate        string `json:"climate"`
	Gravity        string `json:"gravity"`
	Terrain        string `json:"terrain"`
	SurfaceWater   int    `json:"surface_water"`
	Population     int    `json:"population"`
}