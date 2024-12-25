package util

import (
	"encoding/json"
	"io"
	"log"
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

func ReadBodyAndGetString(body io.ReadCloser) (string, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func ReadBodyAndPrintJSON[T any](body io.ReadCloser) error {
	var obj T
	obj, err := ReadBodyAndUnmarshal(obj, body)
	if err != nil {
		return err
	}
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	log.Print(string(b))
	return nil
}
