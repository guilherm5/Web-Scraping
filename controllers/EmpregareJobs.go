package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/guilherm5/Scraping-Jobs/models"
)

func EmpregareJobs(input string) string {
	var jobs models.EmpregareObject
	url := fmt.Sprintf(`https://www.empregare.com/api/pt-br/vagas/buscar-novo?pagina=1&itensPagina=50&query=%s`, input)
	metodo := "GET"
	headerKey := ""
	headerValue := ""

	Body := InterfaceHTTP(metodo, url, headerKey, headerValue)

	err := json.Unmarshal([]byte(Body), &jobs)
	if err != nil {
		fmt.Println("Erro ao ler json", err)
	}

	file, err := os.Create("Empregare-" + time.Now().Format("02-01-2006") + ".csv")
	if err != nil {
		log.Println("Erro ao criar arquivo .csv", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Vaga", "Salario", "Modalidade"}
	writer.Write(header)

	for _, job := range jobs.Model.Dados {
		if job.Salario == "" {
			job.Salario = "[Sálario não informado.]"
		}
		fmt.Println("Vaga: ", job.Titulo)
		fmt.Println("Sálario: ", job.Salario)
		fmt.Println("Modalidade: ", job.TrabalhoRemoto+"-"+job.TrabalhoRemotoTexto)
		fmt.Println("======================")

		// preenche arquivo .csv
		header = append(header, job.Titulo)
		header = append(header, job.Salario)
		header = append(header, job.TrabalhoRemoto)
		writer.Write(header)
	}
	qtdeVagas, _ := fmt.Println("Quantidade de vagas encontradas: ", len(jobs.Model.Dados))

	return string(qtdeVagas)
}
