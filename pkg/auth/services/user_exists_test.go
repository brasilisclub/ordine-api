package services

import (
	"ordine-api/pkg/auth"
	"ordine-api/tests"
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
				tests.MakeMigrationsForTests(&auth.User{})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&auth.User{})
			},
		},
		{
			name: "Should return true",
			args: args{
				userName: "test",
			},
			want: true,
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
			if got := UserExists(tt.args.userName); got != tt.want {
				t.Errorf("UserExists() = %v, want %v", got, tt.want)
			}
			tt.dropDownTest()
		})
	}
}
