package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/guilherm5/Scraping-Jobs/models"
)

func EmpregareJobs(input string) string {
	var jobs models.EmpregareObject
	url := fmt.Sprintf(`https://www.empregare.com/api/pt-br/vagas/buscar-novo?pagina=1&itensPagina=50&query=%s`, input)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Erro ao enviar requisição para o site desejado", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("Erro ao estabelecer conexão com o servidor desejado", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler body", err)

	}

	err = json.Unmarshal([]byte(body), &jobs)
	if err != nil {
		fmt.Println("Erro ao ler json", err)
	}

	for _, job := range jobs.Model.Dados {
		if job.Salario == "" {
			job.Salario = "[Sálario não informado.]"
		}
		fmt.Println("Vaga: ", job.Titulo)
		fmt.Println("Sálario: ", job.Salario)
		fmt.Println("Modalidade: ", job.TrabalhoRemoto+"-"+job.TrabalhoRemotoTexto)
		fmt.Println("======================")
	}
	fmt.Println("Quantidade de vagas encontradas: ", len(jobs.Model.Dados))

	return ""
}
