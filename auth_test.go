package paysimple

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAuthorization(t *testing.T) {
	assert := assert.New(t)

	then := time.Date(2012, 7, 20, 20, 45, 44, 97392800, time.UTC)
	assert.Equal("2012-07-20T20:45:44.0973928Z", then.Format(time.RFC3339Nano))

	// TODO Need the apiKey they used to generate this signature
	// assert.Equal(
	//  `PSSERVER accessid=APIUser1000; timestamp=; signature=WqV47Dddgc6XqBKnQASzZbNU/UZd1tzSrFJJFVv76dw=`,
	//  createAuthorization("APIUser1000", "", then),
	// )
}
