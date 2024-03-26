package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID               int
	Username         string
	Name             string
	Rank             int
	Score            float64
	Country          string
	Country_code     string
	Country_rank     int
	Subdivision      string
	Subdivision_code string
	Subdivision_rank int
	University       string
	University_code  string
	University_rank  int
}
