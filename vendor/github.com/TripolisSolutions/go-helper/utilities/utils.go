package utilities

import "encoding/json"

// ToJSON marshals interface into json []byte
func ToJSON(results interface{}) []byte {
	b, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	return b
}
