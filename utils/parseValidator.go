package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ParseValidator(err error) map[string][]string {
	errors := err.(validator.ValidationErrors)
	result := make(map[string][]string)
	for _, e := range errors {
		var sb strings.Builder

		sb.WriteString("validation failed on field '" + e.Field() + "'")
		sb.WriteString(", condition: " + e.ActualTag())

		// Print condition parameters, e.g. oneof=red blue -> { red blue }
		if e.Param() != "" {
			sb.WriteString(" { " + e.Param() + " }")
		}

		if e.Value() != nil && e.Value() != "" {
			sb.WriteString(fmt.Sprintf(", actual: %v", e.Value()))
		}

		//return sb.String()
		result[e.Field()] = append(result[e.Field()], sb.String())
	}
	return result
}
