package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Msaorc/SearchZipCode/server/packages"
)

func main() {
	file, err := os.Create("CEP.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error of create file: %v", err)
	}
	defer file.Close()

	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf(packages.Url, cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error querying zip code, error: %v\n\n", err)
		}
		defer req.Body.Close()

		response, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading reply: %v\n\n", err)
		}

		var responseCep packages.ObjectZipCode
		err = json.Unmarshal(response, &responseCep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error of converting json to struct: %v\n\n", err)
		}

		file.WriteString(fmt.Sprintf("CEP: %s\nStreet: %s\nDistrict: %s\nCity: %s-%s\n\n",
			responseCep.Cep, responseCep.Logradouro, responseCep.Bairro, responseCep.Localidade, responseCep.Uf))
	}
}
