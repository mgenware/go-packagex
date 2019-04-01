package jsonx

import "encoding/json"

// UnmarshalBytesToDict unmarshals the given bytes into a dictionary of type `map[string]interface{}`.
func UnmarshalBytesToDict(bytes []byte) (map[string]interface{}, error) {
	dict := make(map[string]interface{})
	err := json.Unmarshal(bytes, &dict)
	if err != nil {
		return nil, err
	}

	return dict, nil
}
