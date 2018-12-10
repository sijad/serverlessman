package main

import "strconv"
import "strings"

// Filed defines config field structure
type Filed struct {
	Name string `json:"name"`

	DefaultValue string `json:"default"`
	InputName    string `json:"input_name"`

	Rquired bool `json:"required"`

	Input *bool `json:"input"`
	Save  *bool `json:"save"`

	Transformer *Transformer `json:"transformer"`

	InputType InputType `json:"type"`

	Min *int `json:"min"`
	Max *int `json:"max"`

	Options []string `json:"options"`
}

// IsValid takes a string and check if it's valid by Filed validators
func (f *Filed) IsValid(val string) bool {
	if f.Rquired && strings.TrimSpace(val) == "" {
		return false
	}

	if val != "" && !f.InputType.IsValid(val) {
		return false
	}

	if f.Min != nil || f.Max != nil {
		nval := len(val)

		if f.InputType == InputTypeNumber {
			v, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			nval = v
		}

		if f.Min != nil && nval < *f.Min {
			return false
		}

		if f.Max != nil && nval > *f.Max {
			return false
		}
	}

	if len(f.Options) > 0 {
		for _, o := range f.Options {
			if val == o {
				return true
			}
		}
		return false
	}

	return true
}

// FormKey retruns input name key
func (f *Filed) FormKey() string {
	if f.InputName != "" {
		return f.InputName
	}
	return f.Name
}
