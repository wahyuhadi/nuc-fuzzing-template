package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/wahyuhadi/nuc-fuzzing-template/repeater"
	"gopkg.in/elazarl/goproxy.v1"
)

var (
	template_file   = flag.String("tf", "", "Location template file ")
	type_attack     = flag.String("t", "snipper", "Attack type")
	proxy_ip        = flag.String("ip", "0.0.0.0", "Set Proxy IP")
	payload         = flag.String("p", "payload/default.txt", "payload location")
	proxy_port      = flag.String("port", "9191", "Set Proxy Port")
	verbose         = flag.Bool("v", false, "Log every request to stdout")
	URI             = flag.String("url", "", "[Mandatory] Set URL to watch on proxy (Default : api.example.com)")
	URIFILE         = flag.String("ufile", "", "Location whiteliest file")
	add_query_param = flag.String("q", "", "Add query params example -q sort,direction,user")
)

// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parse_config() repeater.Config {
	flag.Parse()
	if *URI == "" && *URIFILE == "" {
		os.Exit(1)
	}
	var config repeater.Config
	list_host := []string{}
	list_host = append(list_host, *URI)

	if *URIFILE != "" {
		list, err := readLines(*URIFILE)
		if err != nil {
			log.Println(err.Error())
		}

		log.Println("add URI ", list)
		list_host = list
	}

	// temp := template.Must(template.ParseFiles(*template_file))
	config.Templates = *template_file
	config.Payload = *payload
	config.ProxyIP = *proxy_ip
	config.ProxyPort = *proxy_port
	config.Verbose = *verbose
	config.URI = list_host
	config.AddQueryParam = *add_query_param
	// config.Temp = temp
	config.TypeAttack = *type_attack

	return config
}

func main() {
	config := parse_config()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose
	log.Println("Service proxy running on port", *proxy_port)
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(repeater.Repeater(config))
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case sig := <-c:
			log.Fatal("Got %s signal. Aborting...\n", sig)
			os.Exit(1)
		}
	}()
	log.Fatal(http.ListenAndServe(*proxy_ip+":"+*proxy_port, proxy))
}
