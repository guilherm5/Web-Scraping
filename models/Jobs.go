package models

type Jobs struct {
	Props struct {
		PageProps struct {
			JobSearch struct {
				JobSearchResult struct {
					Data struct {
						Jobs []struct {
							JobCustomizedData struct {
								FaixaSalarial string   `json:"faixaSalarial"`
								Titulo        string   `json:"titulo"`
								Descricao     string   `json:"descricao"`
								Beneficios    []string `json:"benef"`
							} `json:"job_customized_data"`
						} `json:"jobs"`
					} `json:"data"`
				} `json:"jobSearchResult"`
			} `json:"jobSearch"`
		} `json:"pageProps"`
	} `json:"props"`
}
