package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    bool
	}{
		{
			name: "Valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey api_key"},
			},
			wantAPIKey: "api_key",
			wantErr:    false,
		},
		{
			name:       "Missing ApiKey",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErr:    true,
		},
		{
			name: "Malforemd request",
			headers: http.Header{
				"Authorization": []string{"InvalidApikey api_key"},
			},
			wantAPIKey: "api_key",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAPIKey != tt.wantAPIKey {
				t.Errorf("GetAPIKey() gotAPIKey = %v, want %v", gotAPIKey, tt.wantAPIKey)
			}
		})
	}
}
