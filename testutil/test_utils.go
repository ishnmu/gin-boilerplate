package testutil

import (
	"encoding/json"
)

const (
	panicMessageOnMarshallError = "panicked on must marshall"
)

//MustMarshal is a test util which panics on error. Strictly use it only for tests.
func MustMarshal(v interface{}) string {
	bytes, error := json.Marshal(v)
	if error != nil {
		panic(panicMessageOnMarshallError)
	}
	return string(bytes)
}
