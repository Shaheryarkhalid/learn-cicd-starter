package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://localhost:8080/verify", http.NoBody)
	key := "81AFHBj8REeoQeYWv7Oazszm975Joak="
	keyString := "ApiKey " + key
	req.Header.Set("Authorization", keyString)
	APIKey, err := GetAPIKey(req.Header)
	if err != nil {
		t.Errorf("Error getting the api key: %v", err)
		t.FailNow()
	}
	if APIKey == "" || APIKey != key {
		t.Errorf("Wrong APIKey Returned by the method APIKey: %v", APIKey)
		t.FailNow()
	}
}
func TestGetAPIKeyWithoutApiKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://localhost:8080/verify", http.NoBody)
	req.Header.Set("Authorization", "")
	APIKey, err := GetAPIKey(req.Header)
	if err == nil {
		t.Errorf("error should not be nil.")
		t.FailNow()
	}
	if APIKey != "" {
		t.Errorf("ApiKey should be empty.")
		t.FailNow()
	}
}
func TestGetAPIKeyWithmalformedApiKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://localhost:8080/verify", http.NoBody)
	key := ""
	keyString := "Apitoken " + key
	req.Header.Set("Authorization", keyString)
	APIKey, err := GetAPIKey(req.Header)
	if err == nil {
		t.Errorf("error should not be nil.")
		t.FailNow()
	}
	if APIKey != "" {
		t.Errorf("ApiKey should be empty.")
		t.FailNow()
	}
}
