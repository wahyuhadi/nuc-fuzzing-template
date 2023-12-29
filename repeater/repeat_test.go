package repeater

import (
	"bytes"
	"net/http"
	"testing"
	"text/template"
)

func TestConfig_gen_fuzzing_jbody(t *testing.T) {
	type fields struct {
		ProxyIP       string
		Payload       string
		ProxyPort     string
		Verbose       bool
		URI           []string
		AddQueryParam string
		Temp          *template.Template
		TypeAttack    string
	}
	type args struct {
		new_raw  http.Request
		req      http.Request
		bodysave bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := Config{
				ProxyIP:       tt.fields.ProxyIP,
				Payload:       tt.fields.Payload,
				ProxyPort:     tt.fields.ProxyPort,
				Verbose:       tt.fields.Verbose,
				URI:           tt.fields.URI,
				AddQueryParam: tt.fields.AddQueryParam,
				Temp:          tt.fields.Temp,
				TypeAttack:    tt.fields.TypeAttack,
			}
			config.gen_fuzzing_jbody(tt.args.new_raw, tt.args.req, tt.args.bodysave)
		})
	}
}
