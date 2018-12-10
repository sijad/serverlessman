package handler

import (
	"regexp"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	slugRegex  = regexp.MustCompile("^[a-z0-9]+(?:-[a-z0-9]+)*$")
	uuidRegex  = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")
	// numberRegex = regexp.MustCompile("^[0-9]*$")
)

// InputType type
type InputType string

// IsValid takes a string and check if it's valid
func (v InputType) IsValid(val string) bool {
	switch v {
	// case InputTypeNumber:
	// 	return numberRegex.MatchString(val)
	case InputTypeEmail:
		return emailRegex.MatchString(val)
	case InputTypeSlug:
		return slugRegex.MatchString(val)
	case InputTypeUUID:
		return uuidRegex.MatchString(val)
	}
	return true
}

// Input Types
const (
	InputTypeNumber      = "number"
	InputTypeEmail       = "email"
	InputTypeSlug        = "slug"
	InputTypeUUID        = "uuid"
	InputTypeRecaptchaV2 = "recaptcha_v2"
)
