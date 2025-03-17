package controller

import (
	"goApp/internal/domain/service"
	"goApp/internal/model"
)

type PetController struct {
	service service.PetService
}

func NewPetController(s service.PetService) PetController {
	return PetController{service: s}
}

func Index(pc *PetController) ([]*model.Pet, error) {
	return pc.service.ListPets()
}

func GetPet(pc *PetController, criteria string) ([]*model.Pet, error) {
	return pc.service.FindPets(criteria)
}

func CreatePet(pc *PetController, pet *model.Pet) error {
	return pc.service.RegisterPet(pet)
}

func UpdatePet(pc *PetController, pet *model.Pet, param string) error {
	return pc.service.UpdatePet(pet, param)
}

func RemovePet(pc *PetController, pet string, confirm string) error {
	return pc.service.RemovePet(pet, confirm)
}
