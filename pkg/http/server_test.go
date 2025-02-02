package http

import (
	"net/http"
	"testing"
)

func TestGetServer(t *testing.T) {
	tests := []struct {
		name             string
		wantedStatusCode int
	}{
		{
			name:             "Should get the server successfully",
			wantedStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetServer()
			if r == nil {
				t.Error("Error getting the server, pointer is nil")
			}
		})
	}
}
