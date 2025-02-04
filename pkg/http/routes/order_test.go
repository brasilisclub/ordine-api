package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestOrderRoutes(t *testing.T) {
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name           string
		args           args
		expectedRoutes int
	}{
		{
			name: "Should get all order routes",
			args: args{
				r: gin.Default(),
			},
			expectedRoutes: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OrderRoutes(tt.args.r)
			if len(tt.args.r.Routes()) != tt.expectedRoutes {
				t.Errorf("Invalid routes number, expected %v but got %v ", tt.expectedRoutes, len(tt.args.r.Routes()))
			}
		})
	}
}
