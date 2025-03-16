package service

import (
	"goApp/internal/domain/repository"
	"goApp/internal/model"
)

type PetService interface {
	RegisterPet(pet *model.Pet) error
	ListPets() ([]*model.Pet, error)
	FindPets(criteria map[string]string) ([]*model.Pet, error)
	UpdatePet(pet *model.Pet, filename string) error
	RemovePet(filename string) error
}

type petService struct {
	repo repository.PetRepository
}

func (s *petService) RegisterPet(pet *model.Pet) error {
	return s.repo.Insert(pet)
}

func (s *petService) ListPets() ([]*model.Pet, error) {
	return s.repo.FindAll()
}

func (s *petService) FindPets(criteria map[string]string) ([]*model.Pet, error) {
	return s.repo.FindByCriteria(criteria)
}

func (s *petService) UpdatePet(pet *model.Pet, filename string) error {
	return s.repo.Update(pet, filename)
}

func (s *petService) RemovePet(filename string) error {
	return s.repo.Delete(filename)
}

func NewPetService(repo repository.PetRepository) PetService {
	return &petService{repo: repo}
}
