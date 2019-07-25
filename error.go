package gconfig

import "encoding/json"

type configError map[string]string

func (err configError) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

func (err configError) add(key, msg string) {
	err[key] = msg
}
