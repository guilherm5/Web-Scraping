package controllers

import (
	"fmt"
)

func Choice() string {
	var input string

	fmt.Println("Você deseja ver as vagas de qual site ?")
	fmt.Println("1 - Catho\n2 - Empregare")
	fmt.Scanln(&input)

	if input == "1" {
		fmt.Println("Qual vaga você deseja buscar ?")
		fmt.Scanln(&input)
		input = CathoJobs(input)

	} else if input == "2" {
		fmt.Println("Qual vaga você deseja buscar ?")
		fmt.Scanln(&input)
		input = EmpregareJobs(input)
	}

	return input
}
