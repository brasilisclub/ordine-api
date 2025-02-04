package database

import (
	"ordine-api/config"
	"testing"
)

func Test_getConnection(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		setUpTest    func()
		dropDownTest func()
	}{
		{
			name: "Should get connection successfully",
			args: args{
				dsn: ":memory:",
			},
			wantErr: false,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name: "Should error, mysql is not running",
			args: args{
				dsn: "fictional dsn",
			},
			wantErr: true,
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

			_, err := getConnection(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
