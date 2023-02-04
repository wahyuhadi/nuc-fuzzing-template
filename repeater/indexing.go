package repeater

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func create_index_bjson(config Config, req *http.Request, body map[string]interface{}) []map[string]interface{} {

	var objects = []map[string]interface{}{}
	log.Println("Creating body json index for fuzzing")
	if config.TypeAttack == "snipper" {
		var bodyo []string
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
	}
	return objects
}

func create_index_query(config Config, req *http.Request) []string {
	log.Println("Creating query param index for fuzzing")
	var array_string []string
	if config.TypeAttack == "snipper" {
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
					query_raw += fmt.Sprintf("%s=%s&", k, i[0])
				}
				count = count + 1
			}
			array_string = append(array_string, query_raw)
		}
	}
	return array_string
}
