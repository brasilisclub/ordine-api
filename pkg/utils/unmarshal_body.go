package utils

import (
	"encoding/json"
	"io"

	"github.com/sirupsen/logrus"
)

func UnmarshalBody(b io.ReadCloser, s any) error {
	body, err := io.ReadAll(b)
	if err != nil {
		logrus.Errorf("Error reading body: %s", err.Error())
	}

	err = json.Unmarshal(body, s)
	return err
}
