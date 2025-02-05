package database

import (
	"ordine-api/config"
	"testing"
)

func TestGetConnector(t *testing.T) {
	tests := []struct {
		name         string
		wantError    bool
		setUpTest    func()
		dropDownTest func()
	}{
		{
			name:      "Should get connection successfully",
			wantError: false,
			setUpTest: func() {
				connection = nil
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name:      "Should not get connection",
			wantError: true,
			setUpTest: func() {
				connection = nil
				config.Envs.ENVIRONMENT = "Another environment"
			},
			dropDownTest: func() {
				config.Envs.ENVIRONMENT = config.Testing
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setUpTest()
			defer tt.dropDownTest()

			got := GetConnector()
			if (got.Error != nil) != tt.wantError {
				t.Errorf("GetConnector() = %v, want %v", (got != nil), tt.wantError)
			}
		})
	}
}
