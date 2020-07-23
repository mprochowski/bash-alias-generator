package processor

import (
	"fmt"
	"github.com/mprochowski/bash-alias-generator/pkg/decoder"
	"strings"
)

type Elements []Element

type Element struct {
	Name      string      `json:"name" yaml:"name"`
	Command   string      `json:"command" yaml:"command"`
	Instances []Instances `json:"instances" yaml:"instances"`
	Children  Elements    `json:"children" yaml:"children"`
}

type Instances struct {
	Name string `json:"name" yaml:"name"`
	Parameters map[string]string `json:"parameters" yaml:"parameters"`
}

func processElement(element Element, namespace []string, aliases []string, instanceParameters map[string]string) []string {
	if len(namespace) >= 1 {
		namespace = append(namespace, `-`)
	}

	instances := element.Instances
	if len(instances) == 0 {
		instances = []Instances{{Name: ``}}
	}

	namespace = append(namespace, element.Name)

	for _, i := range instances {

		for k, v := range i.Parameters {
			instanceParameters[k] = v
		}

		if len(element.Command) > 0 {
			command := element.Command

			for a, b := range instanceParameters {
				command = strings.ReplaceAll(command, a, b)
			}

			alias := fmt.Sprintf(`alias %s = '%s'`, strings.Join(namespace, ``) + i.Name, command)
			aliases = append(aliases, alias)
		}

		for _, child := range element.Children {
			aliases = processElement(child, []string{strings.Join(namespace, ``) + i.Name}, aliases, instanceParameters)
		}
	}

	return aliases
}


func Process(data string) ([]string, error) {
	var elements Elements
	var aliases []string

	err := decoder.Decode([]byte(data), &elements)

	if err != nil {
		return nil, err
	}

	for _, element := range elements {
		aliases = processElement(element, []string{}, aliases, map[string]string{})
	}

	return aliases, nil
}