package utils

import (
	"encoding/json"
)

func CopyStructToStruct(input interface{}, output interface{}) error {
	if byteData, err := json.Marshal(input); err == nil {
		if err := json.Unmarshal(byteData, &output); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}
