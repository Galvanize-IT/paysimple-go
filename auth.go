package paysimple

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

func CreateAuthorization(username, apiKey string) string {
	return createAuthorization(username, apiKey, time.Now().UTC())
}

func createSignature(apiKey string, now time.Time) string {
	hash := hmac.New(sha256.New, []byte(apiKey))
	hash.Write([]byte(now.Format(time.RFC3339Nano)))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func createAuthorization(username, apiKey string, now time.Time) string {
	return fmt.Sprintf(
		`PSSERVER accessid=%s; timestamp=%s; signature=%s`,
		username,
		now.Format(time.RFC3339Nano),
		createSignature(apiKey, now),
	)
}
