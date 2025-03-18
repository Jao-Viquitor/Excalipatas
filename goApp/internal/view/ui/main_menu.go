package ui

import (
	"bufio"
	"fmt"
	"goApp/internal/controller"
	"os"
	"strconv"
	"strings"
)

func MainMenu(ctrl controller.PetController) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("====================================")
		fmt.Println("          MENU PRINCIPAL")
		fmt.Println("====================================")
		fmt.Println("1. Cadastrar um novo pet")
		fmt.Println("2. Alterar os dados do pet cadastrado")
		fmt.Println("3. Deletar um pet cadastrado")
		fmt.Println("4. Listar todos os pets cadastrados")
		fmt.Println("5. Listar pets por critério (idade, nome, raça)")
		fmt.Println("6. Sair")
		fmt.Print("Digite o número da opção desejada: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			continue
		}

		input = strings.TrimSpace(input)
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Por favor, digite um número válido.")
			continue
		}

		switch option {
		case 1:
			CreatePet(ctrl)
		case 2:
			UpdatePet(ctrl)
		case 3:
			DeletePet(ctrl)
		case 4:
			ShowPet(ctrl)
		case 5:
			FilterPet(ctrl)
		case 6:
			fmt.Println("Encerrando a aplicação. Até mais!")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
