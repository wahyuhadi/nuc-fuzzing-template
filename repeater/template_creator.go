package repeater

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

func create_template(new_raw http.Request, newDumpRequest []byte, config Config, index string) {
	name := strings.Replace(new_raw.URL.Path, "/", "-", -2)
	nucleiPATH := "fuzzing-" + strings.ToLower(new_raw.Method) + name + "-" + index
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
