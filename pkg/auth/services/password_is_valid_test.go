package services

import (
	"ordine-api/pkg/auth"
	"testing"
)

func TestPasswordIsValid(t *testing.T) {
	type args struct {
		user *auth.User
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		setUpTest    func()
		dropDownTest func()
	}{
		{
			name: "Should execute with error, because table do not exist",
			args: args{
				user: &auth.User{
					Username: "test",
					Password: "test",
				},
			},
			wantErr: true,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
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
			if err := PasswordIsValid(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("PasswordIsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.dropDownTest()
		})
	}
}
