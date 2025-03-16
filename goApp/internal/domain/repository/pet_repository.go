package repository

import (
	"bufio"
	"fmt"
	"goApp/internal/model"
	"goApp/internal/model/enums"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type PetRepository interface {
	Insert(pet *model.Pet) error
	FindAll() ([]*model.Pet, error)
	FindByCriteria(criteria map[string]string) ([]*model.Pet, error)
	Update(pet *model.Pet, filename string) error
	Delete(filename string) error
}

type petRepository struct {
	baseDir string
}

func (p *petRepository) Insert(pet *model.Pet) error {
	filename := fmt.Sprintf("%sT%s-%s.txt",
		pet.CreatedAt.Format("20060102"),
		pet.CreatedAt.Format("1504"),
		strings.ToUpper(pet.Name+pet.Surname))
	filePath := filepath.Join(p.baseDir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		_ = fmt.Errorf("erro ao criar arquivo %s: %v", filePath, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	writer := bufio.NewWriter(file)
	_, err = fmt.Fprintf(writer,
		"%s %s\n%s\n%s\n%s, %s, %s\n%.1f\n%.1f\n%s",
		pet.Name, pet.Surname,
		pet.Type,
		pet.Sex,
		pet.Address.Street, pet.Address.Number, pet.Address.City,
		pet.Age,
		pet.Weight,
		pet.Breed,
	)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo %s: %v", filePath, err)
	}
	_ = writer.Flush()
	return nil
}

func (p *petRepository) FindAll() ([]*model.Pet, error) {
	files, err := filepath.Glob(filepath.Join(p.baseDir, "*.TXT"))
	if err != nil {
		return nil, fmt.Errorf("erro ao listar arquivos: %v", err)
	}

	var pets []*model.Pet
	for _, file := range files {
		pet, err := p.readPetFromFile(file)
		if err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	return pets, nil
}

func (p *petRepository) FindByCriteria(criteria map[string]string) ([]*model.Pet, error) {
	pets, err := p.FindAll()
	if err != nil {
		return nil, err
	}

	var filtered []*model.Pet
	for _, pet := range pets {
		matches := true
		for key, value := range criteria {
			switch strings.ToLower(key) {
			case "name":
				if !strings.Contains(strings.ToLower(pet.Name), strings.ToLower(value)) &&
					!strings.Contains(strings.ToLower(pet.Surname), strings.ToLower(value)) {
					matches = false
				}
			case "type":
				if strings.ToLower(string(pet.Type)) != strings.ToLower(value) {
					matches = false
				}
			case "sex":
				if strings.ToLower(string(pet.Sex)) != strings.ToLower(value) {
					matches = false
				}
			case "age":
				age, _ := strconv.ParseFloat(value, 64)
				if pet.Age != age {
					matches = false
				}
			case "weight":
				weight, _ := strconv.ParseFloat(value, 64)
				if pet.Weight != weight {
					matches = false
				}
			case "breed":
				if strings.ToLower(pet.Breed) != strings.ToLower(value) {
					matches = false
				}
			case "address":
				addr := strings.ToLower(fmt.Sprintf("%s %s %s", pet.Address.Street, pet.Address.Number, pet.Address.City))
				if !strings.Contains(addr, strings.ToLower(value)) {
					matches = false
				}
			}
			if !matches {
				break
			}
		}
		if matches {
			filtered = append(filtered, pet)
		}
	}
	return filtered, nil
}

func (p *petRepository) Update(pet *model.Pet, filename string) error {
	if err := p.Delete(filename); err != nil {
		return err
	}
	return p.Insert(pet)
}

func (p *petRepository) Delete(filename string) error {
	filePath := filepath.Join(p.baseDir, filename)
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("erro ao deletar arquivo %s: %v", filePath, err)
	}
	return nil
}

func (p *petRepository) readPetFromFile(filePath string) (*model.Pet, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo %s: %v", filePath, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 7)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo %s: %v", filePath, err)
	}
	if len(lines) != 7 {
		return nil, fmt.Errorf("formato inválido no arquivo %s", filePath)
	}

	// Extrai os dados
	nameParts := strings.SplitN(lines[0], " ", 2)
	petType := enums.PetType(lines[1])
	petSex := enums.PetSex(lines[2])
	addrParts := strings.Split(lines[3], ", ")
	age, _ := strconv.ParseFloat(lines[4], 64)
	weight, _ := strconv.ParseFloat(lines[5], 64)

	// Extrai a data do nome do arquivo
	filename := filepath.Base(filePath)
	timeStr := strings.Split(filename, "-")[0]
	createdAt, _ := time.Parse("20060102T1504", timeStr)

	pet := &model.Pet{
		Name:    nameParts[0],
		Surname: nameParts[1],
		Type:    petType,
		Sex:     petSex,
		Address: model.Address{
			Street: addrParts[0],
			Number: addrParts[1],
			City:   addrParts[2],
		},
		Age:       age,
		Weight:    weight,
		Breed:     lines[6],
		CreatedAt: createdAt,
	}
	return pet, nil
}

func NewPetRepository(baseDir string) PetRepository {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		panic(fmt.Sprintf("Erro ao criar o diretório %s: %v", baseDir, err))
	}
	return &petRepository{baseDir: baseDir}
}
