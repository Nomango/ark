package pretty_test

import (
	"fmt"
	"testing"

	"github.com/Nomango/ark/pretty"
	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {
	type S struct {
		I    int64 `json:"i"`
		Data struct {
			S string `json:"s"`
		} `json:"data"`
	}
	var s S
	s.I = 1
	s.Data.S = "test"

	result := fmt.Sprint(pretty.JSON(s))
	require.Equal(t, result, `{"i":1,"data":{"s":"test"}}`)

	result = fmt.Sprintf("%s", pretty.JSON(s))
	require.Equal(t, result, `{"i":1,"data":{"s":"test"}}`)

	result = fmt.Sprintf("%v", pretty.JSON(s))
	require.Equal(t, result, `{"i":1,"data":{"s":"test"}}`)

	result = fmt.Sprintf("%+v", pretty.JSON(s))
	require.Equal(t, result, `{"i":1,"data":{"s":"test"}}`)

	result = fmt.Sprintf("%#v", pretty.JSON(s))
	require.Equal(t, result, `{
  "i": 1,
  "data": {
    "s": "test"
  }
}`)
}
