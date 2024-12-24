package util

import (
	"encoding/json"
	"io"
)

func ReadBodyAndUnmarshal[T any](val T, body io.ReadCloser) (T, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return val, err
	}

	if err = json.Unmarshal(b, &val); err != nil {
		return val, err
	}

	return val, nil
}
