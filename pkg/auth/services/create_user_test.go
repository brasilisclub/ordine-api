package services

import (
	"ordine-api/pkg/auth"

	"ordine-api/pkg/database"
	"testing"

	"github.com/sirupsen/logrus"
)

var c = database.GetConnector()

func TestCreateUser(t *testing.T) {
	type args struct {
		user *auth.AuthRequestBody
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		setUpTest    func()
		dropDownTest func()
	}{
		{
			name: "Should return error, user already exists",
			args: args{
				user: &auth.AuthRequestBody{Username: "test", Password: "test"},
			},
			wantErr: true,
			setUpTest: func() {
				c.AutoMigrate(&auth.User{})
				hp, _ := HashPassword("test")
				c.Save(&auth.User{Username: "test", Password: hp})
			},
			dropDownTest: func() {
				err := c.Migrator().DropTable(&auth.User{})
				if err != nil {
					logrus.Error(err.Error())
				}
			},
		},
		{
			name: "Should return error, error trying hash password(password > 72)",
			args: args{
				user: &auth.AuthRequestBody{Username: "test", Password: "Password longer than 72 characters......................................."},
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
			name: "Should create a new user",
			args: args{
				user: &auth.AuthRequestBody{Username: "test", Password: "test"},
			},
			wantErr: false,
			setUpTest: func() {
				c.AutoMigrate(&auth.User{})
			},
			dropDownTest: func() {
				c.Migrator().DropTable(&auth.User{})
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.setUpTest()
			defer tt.dropDownTest()
			if err := CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
