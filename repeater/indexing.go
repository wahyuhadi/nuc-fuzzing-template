package repeater

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func create_index_form(config Config, req *http.Request, bodysave bytes.Buffer) []string {
	var indexed []string
	log.Println("Creating body json index for fuzzing")
	if config.TypeAttack == "snipper" {
		indexed = snipper_form(config, req, bodysave)
	}
	return indexed
}

func snipper_form(config Config, req *http.Request, bodysave bytes.Buffer) []string {
	var array_string []string
	form_body, _ := url.ParseQuery(bodysave.String())
	fmt.Println(form_body)

	var param []string
	for k, _ := range form_body {
		param = append(param, k)
	}

	for _, parv := range param {
		count := 1
		var query_raw string
		for k, i := range form_body {
			if k == parv {
				log.Println("Indexing object param ", k)
				query_raw += fmt.Sprintf("%s=%s&", k, fmt.Sprintf("§%v§", "path"))
			} else {
				log.Println("Indexing object param ", k)
				t := &url.URL{Path: i[0]}
				query_raw += fmt.Sprintf("%s=%s&", k, t.String())
			}
			count = count + 1
		}
		array_string = append(array_string, query_raw)
	}

	return array_string
}

func create_index_bjson(config Config, req *http.Request, body map[string]interface{}) []map[string]interface{} {

	var indexed = []map[string]interface{}{}
	log.Println("Creating body json index for fuzzing")
	if config.TypeAttack == "snipper" {
		indexed = config.snipper_bjson(req, body)
	}
	return indexed
}

func (config Config) snipper_bjson(req *http.Request, body map[string]interface{}) []map[string]interface{} {
	var bodyo []string
	var objects = []map[string]interface{}{}

	for k, _ := range body {
		bodyo = append(bodyo, k)
	}
	for _, object := range bodyo {
		var new_body = make(map[string]interface{})
		for k, i := range body {
			if object == k {
				log.Println("Indexing object body ", k)
				body := fmt.Sprintf("§%v§", "path")
				new_body[k] = body
			} else {
				log.Println("Indexing objec body ", k)
				new_body[k] = i
			}
		}
		objects = append(objects, new_body)
	}

	return objects
}

func create_index_query(config Config, req *http.Request) []string {
	log.Println("Creating query param index for fuzzing")
	var indexed []string
	if config.TypeAttack == "snipper" {
		indexed = snipper_query(config, req)
	}

	return indexed
}

func snipper_query(config Config, req *http.Request) []string {
	var array_string []string
	url_query := req.URL.Query()
	if config.AddQueryParam != "" {
		params := strings.Split(config.AddQueryParam, ",")
		for _, newQ := range params {
			url_query.Add(newQ, "null")
		}
	}

	var param []string
	for k, _ := range url_query {
		param = append(param, k)
	}

	for _, parv := range param {
		count := 1
		var query_raw string
		for k, i := range url_query {
			if k == parv {
				log.Println("Indexing object param ", k)
				query_raw += fmt.Sprintf("%s=%s&", k, fmt.Sprintf("§%v§", "path"))
			} else {
				log.Println("Indexing object param ", k)
				t := &url.URL{Path: i[0]}
				query_raw += fmt.Sprintf("%s=%s&", k, t.String())
			}
			count = count + 1
		}
		array_string = append(array_string, query_raw)
	}

	return array_string
}
