package webutil

import (
	"encoding/json"
	"io"
)

type OK interface {
	OK() error
}

func Decode(body io.ReadCloser, v interface{}) error {
	if err := json.NewDecoder(body).Decode(v); err != nil {
		return err
	}
	if validatable, ok := v.(OK); ok {
		return validatable.OK()
	}
	return nil
}
