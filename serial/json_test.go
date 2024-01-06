package serial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonSToObj(t *testing.T) {
	type Foo struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	expectedObj := Foo{A: 1, B: "two"}
	expectedStr := `{"a":1,"b":"two"}`

	assert.Equal(t, expectedStr, ObjToJsonS(&expectedObj))
	assert.Equal(t, expectedObj, JsonSToObj[Foo](expectedStr))
}
