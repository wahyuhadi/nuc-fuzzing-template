package repeater

import (
	"bytes"
	"text/template"
)

type Data struct {
	Raw      string
	Name     string
	Payload  string
	TempName string
}

type BodySave struct {
	Body bytes.Buffer
}

type Config struct {
	ProxyIP       string
	Payload       string
	ProxyPort     string
	Verbose       bool
	URI           []string
	AddQueryParam string
	Temp          *template.Template
	TypeAttack    string
}
