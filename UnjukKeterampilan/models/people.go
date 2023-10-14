package models

type People struct {
	Id        uint   `gorm:"primary key" json:"id" form:"id"`
	Name      string `gorm:"not null" json:"name"`
	Height    int    `json:"height"`
	Mass      int    `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
}