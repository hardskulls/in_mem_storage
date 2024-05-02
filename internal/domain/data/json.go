package data

import "encoding/json"

type JSON struct {
	contents []byte
}

func ToJson[T any](v T) (JSON, error) {
	contents, err := json.Marshal(v)
	return JSON{contents: contents}, err
}

func FromJson[T any](j JSON) (T, error) {
	var v T
	err := json.Unmarshal(j.contents, &v)
	return v, err
}
