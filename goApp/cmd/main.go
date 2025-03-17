package cmd

import (
	"goApp/internal/controller"
	"goApp/internal/domain/repository"
	"goApp/internal/domain/service"
)

func main() {
	repo := repository.NewPetRepository("petsCadastrados/")
	svc := service.NewPetService(repo)
	ctrl := controller.NewPetController(svc)
}
