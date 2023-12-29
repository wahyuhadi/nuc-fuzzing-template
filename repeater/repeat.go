package repeater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"gopkg.in/elazarl/goproxy.v1"
)

func Repeater(config Config) func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		var b bytes.Buffer
		b.ReadFrom(req.Body)

		var bodysave BodySave
		bodysave.Body = b
		newBody := io.NopCloser(&b)

		//Modified Request@!
		if Contains(config.URI, req.Host) && req.Method != "OPTIONS" && req.Method != "HEAD" {
			var new_raw http.Request

			is_host := strings.Split(req.Host, ":")
			// log.Println(is_host, "is host")
			new_raw.Method = req.Method
			new_raw.Proto = req.Proto
			new_raw.ProtoMajor = req.ProtoMajor
			new_raw.ProtoMinor = req.ProtoMinor
			new_raw.Host = is_host[0]
			new_raw.Header = req.Header
			new_raw.URL = req.URL
			lenQuery := len(req.URL.Query())

			if config.AddQueryParam != "" {
				lenQuery = len(config.AddQueryParam)
			}
			//Generate fuzzing in query param
			if lenQuery != 0 && req.Method == "GET" {
				config.gen_fuzzing_query(&new_raw, req)
			}

			uri, _ := url.Parse(req.URL.String())
			new_raw.URL = uri
			log.Println(uri)

			if req.Method == "POST" || req.Method == "PATCH" || req.Method == "PUT" {
				type_body := check_body_type(bodysave.Body.Bytes(), req)
				log.Println(type_body)
				if type_body == "json" {
					// json body
					config.gen_fuzzing_jbody(new_raw, *req, bodysave.Body)
				}

				if type_body == "form" {
					config.gen_fuzzing_form(&new_raw, req, bodysave.Body)
				}
			}
		}

		req.Body = newBody
		return req, nil
	}
}

func (config Config) gen_fuzzing_form(new_raw, req *http.Request, bodysave bytes.Buffer) {
	forms := create_index_form(config, req, bodysave)
	for l, form := range forms {
		content, _ := strconv.ParseInt(string(form), 10, 64)
		new_raw.ContentLength = content
		new_raw.Body = io.NopCloser(strings.NewReader(string(form)))
		dump, _ := httputil.DumpRequest(new_raw, true)
		newDumpRequest := new_dump(dump)
		create_template(*new_raw, newDumpRequest, config, fmt.Sprintf("post-form-%v", l))
	}

}

func check_body_type(s []byte, req *http.Request) string {
	log.Println("Validating request body")
	var js map[string]interface{}
	if json.Unmarshal(s, &js) == nil {
		log.Println("Body type is json")
		return "json"
	}

	if strings.Contains(req.Header.Get("Content-Type"), "form") {
		log.Println("Body type is post form ")
		return "form"
	}
	log.Println("undifined body format")
	return "undifined"
}

func (config Config) gen_fuzzing_query(new_raw, req *http.Request) {
	uri, _ := url.Parse(req.URL.String())
	new_raw.URL = uri
	array_query := create_index_query(config, req)
	for l, query := range array_query {
		uri.RawQuery = query
		dump, _ := httputil.DumpRequest(new_raw, true)
		newDumpRequest := new_dump(dump)
		create_template(*new_raw, newDumpRequest, config, fmt.Sprintf("query-%v", l))
	}
}

// json body
func (config Config) gen_fuzzing_jbody(new_raw, req http.Request, bodysave bytes.Buffer) {
	var body map[string]interface{}
	json.NewDecoder(&bodysave).Decode(&body)
	objs := create_index_bjson(config, &req, body)
	for l, o := range objs {
		fuzzedJSON, _ := json.Marshal(o)
		content, _ := strconv.ParseInt(string(fuzzedJSON), 10, 64)
		new_raw.Body = io.NopCloser(strings.NewReader(string(fuzzedJSON)))
		new_raw.ContentLength = content
		dump, _ := httputil.DumpRequest(&new_raw, true)
		newDumpRequest := new_dump(dump)
		create_template(new_raw, newDumpRequest, config, fmt.Sprintf("body-%v", l))
	}
}
