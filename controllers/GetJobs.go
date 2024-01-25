package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/guilherm5/Scraping-Jobs/models"
)

func GetJobs(input string) string {
	var Jobs models.Jobs
	Emprego := fmt.Sprintf(`https://www.catho.com.br/vagas/%s/`, input)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, Emprego, nil)
	if err != nil {
		fmt.Println("Erro ao enviar requisição para o servidor desejado.", err)
	}

	// Único header necessário
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao estabelecer conexão com o servidor desejado.", err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Erro ao criar o documento GoQuery.", err)
	}

	// goquery(mesma sintaxe que o jquery para achar css[mas nunca usei jquery ksks])
	scriptContent := document.Find("script#__NEXT_DATA__").Text()

	err = json.Unmarshal([]byte(scriptContent), &Jobs)
	if err != nil {
		fmt.Println("Erro ao ler json", err)
	}

	for _, job := range Jobs.Props.PageProps.JobSearch.JobSearchResult.Data.Jobs {
		fmt.Println("Vaga: ", job.JobCustomizedData.Titulo)

		fmt.Println("Salario: ", job.JobCustomizedData.FaixaSalarial)

		fmt.Println("Beneficios: ", job.JobCustomizedData.Beneficios)
		fmt.Println("===============")
	}

	return ""
}
