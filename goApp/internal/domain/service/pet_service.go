package service

import (
	"errors"
	"goApp/internal/domain/repository"
	"goApp/internal/model"
	"goApp/internal/utils"
	"strings"
	"time"
)

type PetService struct {
	repo repository.PetRepository
}

func (s *PetService) RegisterPet(pet *model.Pet) error {
	pet.Name, pet.Surname, _ = utils.ValidateName(pet)
	pet.Type, _ = utils.ValideteType(pet)
	pet.Sex, _ = utils.ValideteGender(pet)
	pet.Address.Number, _ = utils.ValidateAddress(pet)
	pet.Age, _ = utils.ValidateAge(pet)
	pet.Weight, _ = utils.ValidateWeight(pet)
	pet.Breed, _ = utils.ValidateBreed(pet)
	pet.CreatedAt = time.Now()
	return s.repo.Insert(pet)
}

func (s *PetService) ListPets() ([]*model.Pet, error) {
	return s.repo.FindAll()
}

func (s *PetService) FindPets(criteria map[string]string) ([]*model.Pet, error) {
	criteria, _ = utils.ValidateCriteria(criteria)
	return s.repo.FindByCriteria(criteria)
}

func (s *PetService) UpdatePet(pet *model.Pet, filename string) error {
	existingPets, err := s.repo.FindByCriteria(map[string]string{"name": pet.Name, "surname": pet.Surname})
	if err != nil || len(existingPets) == 0 {
		return errors.New("pet não encontrado para atualização")
	}
	original := existingPets[0]
	pet.Type = original.Type
	pet.Sex = original.Sex

	if err := s.RegisterPet(pet); err != nil {
		return err
	}
	return s.repo.Update(pet, filename)
}

func (s *PetService) RemovePet(filename string, confirm string) error {
	if strings.ToUpper(confirm) != "SIM" {
		return errors.New("deleção cancelada: confirmação deve ser 'SIM'")
	}
	return s.repo.Delete(filename)
}

func NewPetService(repo repository.PetRepository) *PetService {
	return &PetService{repo: repo}
}
