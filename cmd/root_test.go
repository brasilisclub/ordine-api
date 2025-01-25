package cmd

import (
	"bytes"
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "should exit 0",
			args: []string{"--help"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				commandBuffer bytes.Buffer
				exitCode      int
			)
			defer commandBuffer.Reset()
			rootCmd.SetOut(&commandBuffer)
			rootCmd.SetArgs(tt.args)
			osExit = func(code int) {
				exitCode = code
			}
			Execute()
			if exitCode != tt.want {
				t.Errorf("command exit code expect %v got %v", tt.want, exitCode)
			}
		})
	}
}
