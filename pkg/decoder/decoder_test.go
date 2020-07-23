package decoder

import (
	"testing"
)

type TestStruct struct {
	Name string `json:"name" yaml:"name"`
}

func TestDecodeJson(t *testing.T) {
	var s TestStruct
	e := TestStruct{"Test"}
	d := `{name: "Test"}`

	err := Decode([]byte(d), &s)
	if err != nil {
		t.Errorf("Decode JSON ended with error %s.", err)
	}

	if e != s {
		t.Errorf("Expected %v got %v", e, s)
	}
}

func TestDecodeYaml(t *testing.T) {
	var s TestStruct
	e := TestStruct{"Test"}
	d := `name: "Test"`

	err := Decode([]byte(d), &s)
	if err != nil {
		t.Errorf("Decode YAML ended with error %s.", err)
	}

	if e != s {
		t.Errorf("Expected %v got %v", e, s)
	}
}

func TestDecodeInvalidFormat(t *testing.T) {
	var s TestStruct
	d := `"Test"`

	err := Decode([]byte(d), &s)
	if err == nil {
		t.Errorf("Expected error got %s.", s)
	}
}
