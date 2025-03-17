package utils

import (
	"errors"
	"goApp/internal/model"
	"goApp/internal/model/enums"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func ValidateName(pet *model.Pet) (string, string, error) {
	if strings.TrimSpace(pet.Name) == "" || strings.TrimSpace(pet.Surname) == "" {
		return "", "", errors.New("nome e sobrenome são obrigatórios")
	}
	nameRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !nameRegex.MatchString(pet.Name) || !nameRegex.MatchString(pet.Surname) {
		return "", "", errors.New("nome e sobrenome devem conter apenas letras")
	}
	return pet.Name, pet.Surname, nil
}

func ValideteType(pet *model.Pet) (enums.PetType, error) {
	if pet.Type != enums.TypeDog && pet.Type != enums.TypeCat {
		return "", errors.New("tipo deve ser 'Cachorro' ou 'Gato'")
	}
	return pet.Type, nil
}

func ValideteGender(pet *model.Pet) (enums.PetSex, error) {
	if pet.Sex != enums.SexMale && pet.Sex != enums.SexFemale {
		return "", errors.New("sexo deve ser 'Macho' ou 'Fêmea'")
	}
	return pet.Sex, nil
}

func ValidateAddress(pet *model.Pet) (string, error) {
	if strings.TrimSpace(pet.Address.Number) == "" {
		pet.Address.Number = NotInformed
	}
	return pet.Address.Number, nil
}

func ValidateAge(pet *model.Pet) (float64, error) {
	if pet.Age < 0 {
		return 0, errors.New("idade não pode ser negativa")
	}
	if pet.Age > 20 {
		return 0, errors.New("idade não pode ser maior que 20 anos")
	}
	if pet.Age < 1 {
		months := math.Round(pet.Age * 12)
		return months, nil
	}
	return pet.Age, nil
}

func ValidateWeight(pet *model.Pet) (float64, error) {
	if pet.Weight < 0.5 || pet.Weight > 60 {
		return -1, errors.New("peso deve estar entre 0.5 kg e 60 kg")
	}
	return pet.Weight, nil
}

func ValidateBreed(pet *model.Pet) (string, error) {
	if strings.TrimSpace(pet.Breed) == "" {
		pet.Breed = NotInformed
	}
	breedRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !breedRegex.MatchString(pet.Breed) {
		return "", errors.New("raça deve conter apenas letras")
	}
	return pet.Breed, nil
}

func ParseAndValidateCriteria(criteriaStr string) (map[string]string, error) {
	criteriaStr = strings.TrimSpace(criteriaStr)
	if criteriaStr == "" {
		return nil, errors.New("critérios de busca não podem ser vazios")
	}

	pairs := strings.Split(criteriaStr, " ")
	criteria := make(map[string]string)

	for _, pair := range pairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			return nil, errors.New("formato inválido: use 'chave=valor'")
		}
		key, value := strings.ToLower(parts[0]), strings.TrimSpace(parts[1])
		if value == "" {
			return nil, errors.New("valor não pode ser vazio")
		}
		criteria[key] = value
	}

	if _, ok := criteria["type"]; !ok {
		return nil, errors.New("o critério 'type' é obrigatório")
	}
	if criteria["type"] != strings.ToLower(string(enums.TypeDog)) && criteria["type"] != strings.ToLower(string(enums.TypeCat)) {
		return nil, errors.New("tipo deve ser 'Cachorro' ou 'Gato'")
	}
	if len(criteria) > 3 {
		return nil, errors.New("máximo de 2 critérios adicionais além de 'type'")
	}

	return criteria, nil
}

func ValidateNumericInput(input string, fieldName string) (float64, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, errors.New(fieldName + " não pode ser vazio")
	}

	numRegex := regexp.MustCompile(`^[0-9]+(\.[0-9]+)?$`)
	if !numRegex.MatchString(input) {
		return 0, errors.New(fieldName + " deve conter apenas números")
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New(fieldName + " inválido: não pôde ser convertido em número")
	}

	return value, nil
}
