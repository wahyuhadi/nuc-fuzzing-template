package repeater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
		newBody := ioutil.NopCloser(&b)

		//Modified Request@!
		if req.Host == config.URI && req.Method != "OPTIONS" && req.Method != "HEAD" {
			var new_raw http.Request

			new_raw.Method = req.Method
			new_raw.Proto = req.Proto
			new_raw.ProtoMajor = req.ProtoMajor
			new_raw.ProtoMinor = req.ProtoMinor
			new_raw.Host = "{{ Hostname }}"
			new_raw.Header = req.Header
			new_raw.URL = req.URL

			//Generate fuzzing in query param
			if len(req.URL.Query()) != 0 && req.Method == "GET" {
				uri, _ := url.Parse(req.URL.String())
				new_raw.URL = uri
				array_query := create_index_query(config, req)
				for l, query := range array_query {
					uri.RawQuery = query
					dump, _ := httputil.DumpRequest(&new_raw, true)
					newDumpRequest := new_dump(dump)
					create_template(new_raw, newDumpRequest, config, fmt.Sprintf("query-%v", l))
				}
			}

			uri, _ := url.Parse(req.URL.String())
			new_raw.URL = uri
			var body map[string]interface{}
			json.NewDecoder(&bodysave.Body).Decode(&body)

			objs := create_index_bjson(config, req, body)
			if req.Method == "POST" || req.Method == "PATCH" || req.Method == "PUT" {
				for l, o := range objs {
					fmt.Println(o)
					fuzzedJSON, _ := json.Marshal(o)
					content, _ := strconv.ParseInt(string(fuzzedJSON), 10, 64)
					new_raw.Body = io.NopCloser(strings.NewReader(string(fuzzedJSON)))
					new_raw.ContentLength = content
					dump, _ := httputil.DumpRequest(&new_raw, true)
					newDumpRequest := new_dump(dump)
					create_template(new_raw, newDumpRequest, config, fmt.Sprintf("body-%v", l))
				}
			}
		}
		req.Body = newBody
		return req, nil
	}
}
