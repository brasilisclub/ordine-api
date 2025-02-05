package database

import (
	"ordine-api/config"
	"testing"
)

func Test_getDsn(t *testing.T) {
	tests := []struct {
		name         string
		want         string
		setUpTest    func()
		dropDownTest func()
	}{
		{
			name: "Should get dsn from database in memory",
			want: ":memory:",
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name: "Should get dsn from mysql",
			want: "root:root@tcp(database:3306)/ordine?charset=utf8mb4&parseTime=True&loc=Local",
			setUpTest: func() {
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

			if got := getDsn(); got != tt.want {
				t.Errorf("getDsn() = %v, want %v", got, tt.want)
			}
		})
	}
}
