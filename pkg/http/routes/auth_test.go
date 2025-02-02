package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthRoutes(t *testing.T) {
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should get all auth routes",
			args: args{
				r: gin.Default(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AuthRoutes(tt.args.r)
			if len(tt.args.r.Routes()) != 2 {
				t.Errorf("Invalid routes number, expected 2 but got %v ", len(tt.args.r.Routes()))
			}
		})
	}
}
