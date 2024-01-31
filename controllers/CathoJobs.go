package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/guilherm5/Scraping-Jobs/models"
)

func CathoJobs(input string) string {
	var jobs models.CathoObject
	url := fmt.Sprintf(`https://www.catho.com.br/vagas/%s/`, input)
	headerKey := "user-agent"
	headerValue := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	metodo := "GET"

	Body := InterfaceHTTP(metodo, url, headerKey, headerValue)

	document, err := goquery.NewDocumentFromReader(strings.NewReader(Body))
	if err != nil {
		fmt.Println("Erro ao criar o documento GoQuery.", err)
	}

	// goquery(mesma sintaxe que o jquery para achar css[mas nunca usei jquery ksks])
	scriptContent := document.Find("script#__NEXT_DATA__").Text()

	err = json.Unmarshal([]byte(scriptContent), &jobs)
	if err != nil {
		fmt.Println("Erro ao ler json", err)
	}

	file, err := os.Create("Catho-" + time.Now().Format("02-01-2006") + ".csv")
	if err != nil {
		log.Println("Erro ao criar arquivo .csv", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Vaga", "Salario", "Beneficios"}
	writer.Write(header)

	for _, job := range jobs.Props.PageProps.JobSearch.JobSearchResult.Data.Jobs {
		fmt.Println("Vaga: ", job.JobCustomizedData.Titulo)
		fmt.Println("Salario: ", job.JobCustomizedData.FaixaSalarial)
		fmt.Println("Beneficios: ", job.JobCustomizedData.Beneficios)
		fmt.Println("===============")
		header = append(header, job.JobCustomizedData.Titulo)
		header = append(header, job.JobCustomizedData.FaixaSalarial)
		header = append(header, job.JobCustomizedData.Beneficios...)
		writer.Write(header)
	}
	qtdeVagas, _ := fmt.Println("Quantidade de vagas encontradas: ", len(jobs.Props.PageProps.JobSearch.JobSearchResult.Data.Jobs))
	return string(qtdeVagas)
}
