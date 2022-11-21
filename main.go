package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "http://viacep.com.br/ws/%s/json/"

type ObjectCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	file, err := os.Create("CEP.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error of create file: %v", err)
	}
	defer file.Close()

	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf(url, cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error querying zip code, error: %v\n\n", err)
		}
		defer req.Body.Close()

		response, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading reply: %v\n\n", err)
		}

		var responseCep ObjectCep
		err = json.Unmarshal(response, &responseCep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error of converting json to struct: %v\n\n", err)
		}

		file.WriteString(fmt.Sprintf("CEP: %s\nStreet: %s\nDistrict: %s\nCity: %s-%s\n\n",
			responseCep.Cep, responseCep.Logradouro, responseCep.Bairro, responseCep.Localidade, responseCep.Uf))
	}
}
