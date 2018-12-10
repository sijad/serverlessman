package main

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
)

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	property := r.PostFormValue("property")
	config, ok := configs[property]
	if property == "" || !ok {
		w.WriteHeader(400)
		fmt.Fprintln(w, "peroperty not found")
		return
	}

	filePath := replaceConstPlaceholders(config.Output.Path)
	fields := make(map[string]string)
	for _, field := range config.Fields {
		key := field.Name
		val := field.DefaultValue

		if field.Input == nil || *field.Input {
			if v := r.PostFormValue(field.FormKey()); v != "" {
				val = v
			}
		} else {
			val = replaceConstPlaceholders(val)
		}

		if !field.IsValid(val) {
			w.WriteHeader(400)
			fmt.Fprintln(w, key+" is not valid")
			return
		}

		if field.Transformer != nil {
			val = field.Transformer.Transform(val)
		}

		if field.Save == nil || *field.Save {
			fields[key] = val
		}

		filePath = replacePlaceholder(filePath, key, val)
	}

	filePath = path.Clean(filePath)

	// newFile := Commit{
	// 	ID:     "123",
	// 	Path:   "commnets/123.yaml",
	// 	Fields: fields,
	// }

	b := new(bytes.Buffer)
	yaml.NewEncoder(b).Encode(fields)

	if err := config.Provider.CreateNewFile(filePath, b); err != nil {
		w.WriteHeader(500)
		fmt.Println("error", err)
		fmt.Fprintln(w, "somthing went wrong")
	}
}

var constPlaceholders = map[string]func() string{
	"TIME_UNIX_NANO": func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	},
	"ISO_DATE": func() string {
		return time.Now().Format(time.RFC3339)
	},
	"UUID": func() string {
		if uid, err := uuid.NewV4(); err == nil {
			return uid.String()
		}
		return ""
	},
}

func replaceConstPlaceholders(val string) string {
	for k, v := range constPlaceholders {
		key := "{" + k + "}"
		if strings.Contains(val, key) {
			val = strings.Replace(val, key, v(), -1)
		}
	}
	return val
}

func replacePlaceholder(s, key, val string) string {
	p := "{$" + key + "}"
	if strings.Contains(s, p) {
		return strings.Replace(s, p, val, -1)
	}
	return s
}
