package services

import (
	"ordine-api/pkg/auth"
	"ordine-api/tests"
	"testing"
)

func TestLogin(t *testing.T) {
	type args struct {
		user *auth.User
	}
	tests := []struct {
		name         string
		args         args
		want         string
		wantErr      bool
		setUpTest    func()
		dropDownTest func()
	}{
		{
			name: "Should error because user do not exist",
			args: args{
				user: &auth.User{
					Username: "test",
					Password: "test",
				},
			},
			wantErr: true,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&auth.User{})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&auth.User{})
			},
		},
		{
			name: "Should error because password is invalid",
			args: args{
				user: &auth.User{
					Username: "test",
					Password: "",
				},
			},
			wantErr: true,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&auth.User{})
				hp, _ := HashPassword("test")
				tests.CreateInsertValueForTests(&auth.User{Username: "test", Password: hp})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&auth.User{})
			},
		},
		{
			name: "Should execute without errors",
			args: args{
				user: &auth.User{
					Username: "test",
					Password: "test",
				},
			},
			wantErr: false,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&auth.User{})
				hp, _ := HashPassword("test")
				tests.CreateInsertValueForTests(&auth.User{Username: "test", Password: hp})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&auth.User{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUpTest()
			_, err := Login(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.dropDownTest()
		})
	}
}
