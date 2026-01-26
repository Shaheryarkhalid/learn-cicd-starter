package auth

import (
	"net/http"
	"testing"
)


func TestGetAPIKey(t *testing.T){  
	req, err := http.NewRequest("GET", "http//google.com", nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	testApiKey :=  "test_this_key"
	req.Header.Set("Authorization", "ApiKey "+testApiKey)
	apiKey, err := GetAPIKey(req.Header)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if apiKey != testApiKey{
		t.Error("not valid key")
		t.Fail()
	}
}
