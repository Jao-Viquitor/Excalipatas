package cmd

import (
	"goApp/internal/controller"
	"goApp/internal/domain/repository"
	"goApp/internal/domain/service"
	"goApp/internal/view/ui"
)

func main() {
	repo := repository.NewPetRepository("petsCadastrados/")
	svc := service.NewPetService(repo)
	ctrl := controller.NewPetController(svc)
	ui.MainMenu(ctrl)
}
