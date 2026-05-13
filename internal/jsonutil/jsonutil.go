package jsonutil

import "encoding/json"

// UnmarshalList handles Last.fm's polymorphic JSON (single object vs. array).
func UnmarshalList[T any](data []byte, slice *[]T) error {
	if len(data) == 0 {
		return nil
	}

	// If it's an object, it's a single item.
	if data[0] == '{' {
		var item T
		if err := json.Unmarshal(data, &item); err != nil {
			return err
		}
		*slice = []T{item}
		return nil
	}

	// Otherwise, it should be an array.
	// If it's an empty string or something else, it might be an empty list.
	// Last.fm sometimes returns an empty string "" instead of [] when empty.
	if string(data) == `""` {
		*slice = []T{}
		return nil
	}

	return json.Unmarshal(data, slice)
}
