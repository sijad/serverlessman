package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// Transformer type
type Transformer string

// Transform convert input to coresponding value
func (t Transformer) Transform(val string) string {
	switch t {
	case TransformerMD5:
		sum := md5.Sum([]byte(val))
		return hex.EncodeToString(sum[:])
	case TransformerBASE64:
		return base64.StdEncoding.EncodeToString([]byte(val))
	default:
		return ""
	}
}

// Transforms
const (
	TransformerMD5    Transformer = "md5"
	TransformerBASE64             = "base64"
	// TransformerEncrypt             = "encrypt" // TODO
	// TransformerSlug                  = "slug" // TODO
)
