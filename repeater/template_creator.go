package repeater

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func new_dump(dump []byte) []byte {
	newDumpRequest := []byte{}
	for index, tmpDump := range bytes.Split(dump, []byte("\r\n")) {
		if index == 0 {
			newDumpRequest = append(newDumpRequest, append(tmpDump, "\r\n"...)...)
			continue
		}
		appendSpace := append([]byte("        "), tmpDump...)
		newDumpRequest = append(newDumpRequest, append(appendSpace, "\r\n"...)...)
	}
	return newDumpRequest
}

func create_template(new_raw http.Request, newDumpRequest []byte, config Config, index string) error {
	if new_raw.URL.Path == "/socket.io/" || strings.Contains(new_raw.URL.Path, "/_next/static/") || strings.Contains(new_raw.URL.Path, "/static/image") {
		return nil
	}

	var templates []string
	var temp *template.Template
	// temp := template.Must(template.ParseFiles("sqli"))
	files, err := os.ReadDir(config.Builder)

	if err != nil {
		fmt.Println("Error when read builder directory")
		return nil
	}
	for _, file := range files {
		templates = append(templates, file.Name())

	}

	if config.Templates != "" {
		templates = strings.Split(config.Templates, ",")
	}

	for _, file := range templates {
		if !is_exist_temp(file, config) {
			fmt.Println("Your template not exist in builder folder : your input is  ", file)
			continue
		}
		temp = template.Must(template.ParseFiles(fmt.Sprintf("%s/%s", config.Builder, file)))
		config.Temp = temp
		path := fmt.Sprintf("%s/%s/%s", config.Templates, config.URI[0], file)
		unique := fmt.Sprintf("%s-%s", new_raw.Method, new_raw.URL.Path)
		hash := sha256.Sum256([]byte(unique))
		nucleiPATH := fmt.Sprintf("%s-fuzzing-%x", file, hash[:])
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			errDir := os.MkdirAll(path, 0755)
			if errDir != nil {
				log.Fatal(err)
			}
		}
		file, err := os.Create(path + nucleiPATH + ".yaml")
		if err != nil {
			log.Fatal(err)
			os.Exit(3)
		}
		data_temp := Data{TempName: nucleiPATH, Payload: *&config.Payload, Raw: fmt.Sprintf("%s", newDumpRequest), Name: nucleiPATH}
		err = config.Temp.Execute(file, data_temp)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Creating template in ", nucleiPATH)
	}
	return nil
}

func is_exist_temp(t string, config Config) bool {
	files, err := os.ReadDir(config.Builder)

	if err != nil {
		fmt.Println("Error when read builder directory")
		return false
	}

	for _, file := range files {
		if t == file.Name() {
			return true
		}

	}
	return false
}
