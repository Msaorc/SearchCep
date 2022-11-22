package packages

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SearchZipCode(zipCode string) (*ObjectZipCode, error) {
	response, error := http.Get(fmt.Sprintf(Url, zipCode))
	if error != nil {
		return nil, error
	}
	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	var objectZipCode ObjectZipCode
	json.Unmarshal(body, &objectZipCode)
	return &objectZipCode, nil
}
