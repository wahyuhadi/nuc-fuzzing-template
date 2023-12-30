package repeater

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/google/uuid"
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
	if new_raw.URL.Path == "/socket.io/" {
		return nil
	}

	var templates []string
	var temp *template.Template
	// temp := template.Must(template.ParseFiles("sqli"))
	files, err := os.ReadDir("builder")

	if err != nil {
		log.Println("Error when read bulder directory")
		return nil
	}
	for _, file := range files {
		templates = append(templates, file.Name())

	}

	if config.Templates != "" {
		templates = strings.Split(config.Templates, ",")
	}

	for _, file := range templates {
		if !is_exist_temp(file) {
			log.Println("Your template not exist in builder folder : your input is  ", file)
			continue
		}
		temp = template.Must(template.ParseFiles(fmt.Sprintf("builder/%s", file)))
		config.Temp = temp

		nucleiPATH := fmt.Sprintf("%s-fuzzing-%v", file, uuid.New())
		_, err := os.Stat("template")
		if os.IsNotExist(err) {
			errDir := os.MkdirAll("template", 0755)
			if errDir != nil {
				log.Fatal(err)
			}
		}
		file, err := os.Create("template/" + nucleiPATH + ".yaml")
		if err != nil {
			log.Fatal(err)
			os.Exit(3)
		}
		data_temp := Data{TempName: nucleiPATH, Payload: *&config.Payload, Raw: fmt.Sprintf("%s", newDumpRequest), Name: nucleiPATH}
		err = config.Temp.Execute(file, data_temp)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Creating template in ", nucleiPATH)
	}
	return nil
}

func is_exist_temp(t string) bool {
	files, err := os.ReadDir("builder")

	if err != nil {
		log.Println("Error when read bulder directory")
		return false
	}

	for _, file := range files {
		if t == file.Name() {
			return true
		}

	}
	return false
}
