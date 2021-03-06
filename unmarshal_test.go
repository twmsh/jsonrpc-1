package jsonrpc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {

	err := Unmarshal(nil, nil)
	require.IsType(t, &Error{}, err)
	assert.Equal(t, ErrorCodeInvalidParams, err.Code)

	src := json.RawMessage([]byte(`{"name":"john"}`))

	err = Unmarshal(&src, nil)
	require.IsType(t, &Error{}, err)
	assert.Equal(t, ErrorCodeInvalidParams, err.Code)

	dst := struct {
		Name string `json:"name"`
	}{}

	err = Unmarshal(&src, &dst)
	require.Nil(t, err)
	assert.Equal(t, "john", dst.Name)
}
