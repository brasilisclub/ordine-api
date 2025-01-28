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
			gotToken := GenerateJWT(tt.args.user)
			if len(gotToken) != 139 {
				t.Errorf("Invalid token length, %s has %v characters, must be 139", gotToken, len(gotToken))
				return
			}

		})
	}
}
