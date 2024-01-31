package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func InterfaceHTTP(metodo string, url string, headerKey string, headerValue string) string {
	client := &http.Client{}

	req, err := http.NewRequest(metodo, url, nil)
	if err != nil {
		fmt.Println("Erro ao enviar requisição para o servidor desejado.", err)
	}
	if headerKey != "" && headerValue != "" {
		req.Header.Add(headerKey, headerValue)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao estabelecer conexão com o servidor desejado.", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler body", err)

	}

	return string(body)
}
