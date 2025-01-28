package services

import (
	"ordine-api/pkg/auth"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	type args struct {
		user *auth.User
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should generate JWT",
			args: args{
				user: &auth.User{
					Username: "test",
					Password: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateJWT(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
