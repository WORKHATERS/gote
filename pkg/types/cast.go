package types

import (
	"encoding/json"
)

func CastTo[T any](data any) (*T, error) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	result := new(T)
	if err := json.Unmarshal(jsonbody, &result); err != nil {
		return nil, err
	}

	return result, err
}
