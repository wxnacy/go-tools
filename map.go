package gotool

import "encoding/json"

// map 转换为 interface
func MapConverForInterface(m map[string]interface{}, i interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, i)
}
