package tools

import (
	"encoding/json"
	"io"
)

func ShortUnmarshal(reader io.Reader, obj interface{}) error {

	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	return nil
}
