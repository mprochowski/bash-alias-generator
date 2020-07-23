package decoder

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
)

var decoders = []func([]byte, interface{}) error {
	json.Unmarshal,
	yaml.Unmarshal,
}

func Decode(in []byte, out interface{}) error {

	for _, decoder := range decoders {
		err := decoder(in, out)

		if err == nil {
			return nil
		}
	}

	return errors.New(`unsupported data format`)
}
