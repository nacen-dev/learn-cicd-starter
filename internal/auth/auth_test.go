package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headerKey     string
		headerValue   string
		expectedErr   string
		expectedValue string
	}{
		{name: "Get API Key", expectedErr: "no error", expectedValue: "test", headerKey: "Authorization", headerValue: "ApiKey -"},
		{name: "No authorization", expectedErr: "test"},
		{name: "Malformed Authorization header ApiKey but no actual value", expectedErr: "test", headerKey: "Authorization"},
		{name: "Malformed Authorization header with no ApiKey format", expectedErr: "malformed authorization header", headerKey: "Authorization", headerValue: "ApiKey"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			header := http.Header{}
			header.Add(test.headerKey, test.headerValue)
			actualValue, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey() expected error to be a substring of actual error: %v and actual error: %v\n", test.expectedErr, err)
				return
			}

			if actualValue != test.expectedValue {
				t.Errorf("Unexpected: TestGetAPIKey() expected value: %s and actual value: %s", test.expectedValue, actualValue)
				return
			}
		})
	}
}
