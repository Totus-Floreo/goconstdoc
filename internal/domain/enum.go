package domain

import (
	"fmt"
	"strings"
)

type EnumValue struct {
	Allowed []string
	Value   string
}

func (e *EnumValue) String() string {
	return e.Value
}

func (e *EnumValue) Set(value string) error {
	for _, allowed := range e.Allowed {
		if strings.EqualFold(value, allowed) {
			e.Value = allowed
			return nil
		}
	}
	return fmt.Errorf("invalid value %q, allowed values are %q", value, strings.Join(e.Allowed, ", "))
}

func (e *EnumValue) Type() string {
	return "enum"
}
