package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestOrdineRoutes(t *testing.T) {
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should get all ordine routes",
			args: args{
				r: gin.Default(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OrdineRoutes(tt.args.r)
			if len(tt.args.r.Routes()) != 5 {
				t.Errorf("Expected get 5 routes, but got only %v", len(tt.args.r.Routes()))
			}
		})
	}
}
