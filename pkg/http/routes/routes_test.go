package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLoad(t *testing.T) {
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should get application routes",
			args: args{
				r: gin.Default(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Load(tt.args.r)
			if len(tt.args.r.Routes()) != 13 {
				t.Errorf("Expected to get 13 routes, but got %v", len(tt.args.r.Routes()))
			}
		})
	}
}
