package services

import (
	"ordine-api/pkg/auth"
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
				c.AutoMigrate(&auth.User{})
			},
			dropDownTest: func() {
				c.Migrator().DropTable(&auth.User{})
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
				c.AutoMigrate(&auth.User{})
				hp, _ := HashPassword("test")
				c.Create(&auth.User{Username: "test", Password: hp})
			},
			dropDownTest: func() {
				c.Migrator().DropTable(&auth.User{})
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
				c.AutoMigrate(&auth.User{})
				hp, _ := HashPassword("test")
				c.Create(&auth.User{Username: "test", Password: hp})
			},
			dropDownTest: func() {
				c.Migrator().DropTable(&auth.User{})
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
