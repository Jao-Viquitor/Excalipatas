package useCases

import (
	"bufio"
	"fmt"
	"goApp/internal/controller"
	"os"
	"strconv"
	"strings"
)

func MainUseCases(ctrl controller.PetController) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler entrada:", err)
		return
	}

	input = strings.TrimSpace(input)
	option, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Por favor, digite um número válido.")
		return
	}

	switch option {
	case 1:
		usecases.CreatePetUsecase()
	case 2:
		usecases.UpdatePetUsecase()
	case 3:
		usecases.DeletePetUsecase()
	case 4:
		usecases.ListAllPetsUsecase()
	case 5:
		usecases.ListPetsByCriteriaUsecase()
	case 6:
		fmt.Println("Encerrando a aplicação. Até mais!")
		return
	default:
		fmt.Println("Opção inválida. Tente novamente.")
	}

}
