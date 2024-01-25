package main

import (
	"fmt"
	"github.com/guilherm5/Scraping-Jobs/controllers"
)

func main() {
	var vagaDesejada string
	fmt.Println("Digite o nome da vaga que voce deseja ocupar")
	fmt.Scanln(&vagaDesejada)

	resultado := controllers.GetJobs(vagaDesejada)
	println(resultado)
}
