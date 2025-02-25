package migrations

import (
	"bytes"
	"encoding/json"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/order"
	"ordine-api/pkg/product"
	"ordine-api/tests"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestMigrate(t *testing.T) {
	tests := []struct {
		name         string
		expectedLog  string
		dropDownTest func()
	}{
		{
			name:        "Should execute migrations successfully",
			expectedLog: "Migrations successfully executed!",
			dropDownTest: func() {

				tests.DropTablesForTests(&ordine.Ordine{}, &product.Product{}, &order.OrderProducts{}, &auth.User{})
				logrus.SetOutput(os.Stdout)
				logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var logBuffer bytes.Buffer
			logBuffer.Reset()
			logrus.SetOutput(&logBuffer)
			logrus.SetFormatter(&logrus.JSONFormatter{})

			Migrate()

			var log struct {
				Msg string `json:"msg"`
			}

			err := json.Unmarshal(logBuffer.Bytes(), &log)

			if err != nil {
				t.Errorf("Error trying to unmarshal log %v", err.Error())
			}

			if log.Msg != tt.expectedLog {
				t.Errorf("Invalid log when run the migrations, expected %v got %v", tt.expectedLog, log.Msg)
			}
		})
	}
}
