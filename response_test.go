package paysimple

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleResponse = `{
    "Meta": {
        "Errors": null,
        "HttpStatus": "OK",
        "HttpStatusCode": 200,
        "PagingDetails": {
            "TotalItems": 3,
            "Page": 1,
            "ItemsPerPage": 200
        }
    },
    "Response": []
}
`

var exampleErrors = ``

type testResponse struct {
	Meta     Meta
	Response []interface{}
}

func TestResponse(t *testing.T) {
	assert := assert.New(t)

	// Test without errors
	var response testResponse
	require.Nil(t, json.Unmarshal([]byte(exampleResponse), &response))

	assert.Nil(response.Meta.Errors)
	assert.Equal(200, response.Meta.HttpStatusCode)
	assert.Equal(3, response.Meta.PagingDetails.TotalItems)
	assert.Equal(1, response.Meta.PagingDetails.Page)
	assert.Equal(200, response.Meta.PagingDetails.ItemsPerPage)
}
