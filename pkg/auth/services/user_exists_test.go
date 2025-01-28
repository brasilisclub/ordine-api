package services

import (
	"ordine-api/pkg/auth"
	"testing"
)

func TestUserExists(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name         string
		args         args
		want         bool
		setUpTest    func()
		dropDownTest func()
	}{
		{
			name: "Should execute with error, because table do not exist",
			args: args{
				userName: "test",
			},
			want: false,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name: "Should return false",
			args: args{
				userName: "test",
			},
			want: false,
			setUpTest: func() {
				c.AutoMigrate(&auth.User{})
			},
			dropDownTest: func() {
				c.Migrator().DropTable(&auth.User{})
			},
		},
		{
			name: "Should return true",
			args: args{
				userName: "test",
			},
			want: true,
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
			if got := UserExists(tt.args.userName); got != tt.want {
				t.Errorf("UserExists() = %v, want %v", got, tt.want)
			}
			tt.dropDownTest()
		})
	}
}
