package jsonrpc

import (
	"encoding/json"
	"github.com/json-iterator/go"
)

// Unmarshal decodes JSON-RPC params.
func Unmarshal(params *json.RawMessage, dst interface{}) *Error {
	if params == nil {
		return ErrInvalidParams()
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if err := json.Unmarshal(*params, dst); err != nil {
		return ErrInvalidParams()
	}
	return nil
}
