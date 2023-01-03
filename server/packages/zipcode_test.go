package packages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchZipCode(t *testing.T) {
	obj, err := SearchZipCode("")
	assert.Nil(t, obj)
	assert.Error(t, err, "ZipCode is empty")
	obj, err = SearchZipCode("17523-440")
	assert.Nil(t, err)
	assert.Equal(t, obj.Cep, "17523-440")
}
