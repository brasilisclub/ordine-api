package services

import (
	"ordine-api/pkg/auth"
	"reflect"
	"testing"
	"time"
)

var now = time.Now()

func Test_getUserFromDb(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name         string
		args         args
		want         auth.User
		wantErr      bool
		setUpTest    func()
		dropDownTest func()
	}{
		{
			name: "Should execute without errors",
			args: args{
				userName: "test",
			},
			want:    auth.User{ID: 1, Username: "test", Password: "test"},
			wantErr: false,
			setUpTest: func() {
				c.AutoMigrate(&auth.User{})
				c.Create(&auth.User{Username: "test", Password: "test"})
			},
			dropDownTest: func() {
				c.Migrator().DropTable(&auth.User{})
			},
		},
		{
			name: "Should error table do not exists",
			args: args{
				userName: "test",
			},
			want:    auth.User{},
			wantErr: true,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUpTest()
			got, err := getUserFromDb(tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserFromDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Username, tt.want.Username) {
				t.Errorf("Username getUserFromDb() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Password, tt.want.Username) {
				t.Errorf("Password getUserFromDb() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("ID getUserFromDb() = %v, want %v", got, tt.want)
			}
			tt.dropDownTest()
		})
	}
}
