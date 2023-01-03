package packages

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func SearchZipCode(zipCode string) (*ObjectZipCode, error) {
	if zipCode == "" {
		return nil, errors.New("ZipCode is empty")
	}
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
