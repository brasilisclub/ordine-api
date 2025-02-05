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
		name           string
		args           args
		expectedRoutes int
	}{
		{
			name: "Should get all auth routes",
			args: args{
				r: gin.Default(),
			},
			expectedRoutes: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AuthRoutes(tt.args.r)
			if len(tt.args.r.Routes()) != tt.expectedRoutes {
				t.Errorf("Invalid routes number, expected %v but got %v ", tt.expectedRoutes, len(tt.args.r.Routes()))
			}
		})
	}
}
