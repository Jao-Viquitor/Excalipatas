package model

import (
	"goApp/internal/model/enums"
	"time"
)

type Pet struct {
	Name      string
	Surname   string
	Type      enums.PetType
	Sex       enums.PetSex
	Address   Address
	Age       float64
	Weight    float64
	Breed     string
	CreatedAt time.Time
}
