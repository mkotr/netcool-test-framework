package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/probe-config-service/domain"
	"git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/probe-config-service/service"
	"github.com/micro/go-micro/config"
)

const (
	CONFIG_FILE = "config.json"
)

//Config xD

type ProbeConfigServer struct {
	service *service.ProbeService
	router  *http.ServeMux
}

func NewProbeConfigServer(service *service.ProbeService) *ProbeConfigServer {
	p := &ProbeConfigServer{
		service,
		http.NewServeMux(),
	}

	return p
}

func (p *ProbeConfigServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router := http.NewServeMux()

	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	router.Handle("/probes/", http.HandlerFunc(p.probesHandler))

	router.ServeHTTP(w, r)
}

func (p *ProbeConfigServer) probesHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/probes/"):]

	if name != "" {
		probe := p.service.GetProbe(name)
		if (probe == domain.Probe{}) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(probe)
		}
	} else {
		probes := p.service.GetAll()
		if len(probes) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		json.NewEncoder(w).Encode(probes)
	}
}

//pollConfig - https://micro.mu/docs/go-config.html
func pollConfig(conf *domain.Config) {
	for {
		w, err := config.Watch()
		if err != nil {
			fmt.Printf("Cannot watch file path 'probes' %v \n", err)
		}

		v, err := w.Next()
		if err != nil {
			fmt.Printf("Cannot get updated file '%v' : %v \n", CONFIG_FILE, err)
		}

		var tempConf domain.Config
		v.Scan(&tempConf)

		fmt.Println("After update")
		*conf = tempConf
	}

}

func main() {
	//read Config file

	config.LoadFile(CONFIG_FILE)

	var conf domain.Config
	config.Scan(&conf)

	go pollConfig(&conf)

	service := service.ProbeService{
		Config: &conf,
	}

	server := NewProbeConfigServer(&service)

	log.Println("Running on port :5000")

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 5000 %v", err)
	}
}
