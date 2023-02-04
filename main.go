package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"text/template"

	"github.com/wahyuhadi/nuc-fuzzing-template/repeater"
	"gopkg.in/elazarl/goproxy.v1"
)

var (
	type_attack     = flag.String("t", "snipper", "Attack type")
	proxy_ip        = flag.String("ip", "0.0.0.0", "Set Proxy IP")
	payload         = flag.String("p", "payload/default.txt", "payload location")
	proxy_port      = flag.String("port", "9191", "Set Proxy Port")
	verbose         = flag.Bool("v", false, "Log every request to stdout")
	URI             = flag.String("url", "", "[Mandatory] Set URL to watch on proxy (Default : api.example.com)")
	add_query_param = flag.String("q", "", "Add query params example -q sort,direction,user")
)

func parse_config() repeater.Config {
	flag.Parse()
	if *URI == "" {
		os.Exit(1)
	}

	var config repeater.Config
	temp := template.Must(template.ParseFiles("template.txt"))
	config.Payload = *payload
	config.ProxyIP = *proxy_ip
	config.ProxyPort = *proxy_port
	config.Verbose = *verbose
	config.URI = *URI
	config.AddQueryParam = *add_query_param
	config.Temp = temp
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
