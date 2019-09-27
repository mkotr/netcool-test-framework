package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	FILENAME = "config.csv"
)

type Probe struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

func ProbeConfigServer(w http.ResponseWriter, r *http.Request) {
	probeName := r.URL.Path[len("/probes/"):]

	probe := Probe{
		probeName,
		"xxx",
		"lxapp6662.dc.corp.telstra.com",
		"4000",
	}

	json.NewEncoder(w).Encode(probe)
	//w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("hello world!")

	//read Config file
	readConfig()

	handler := http.HandlerFunc(ProbeConfigServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("Could not listen on port 5000 %v", err)
	}
}

func readConfig() {

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatalf("Error occured when opening the config file: %v\n", err)
	}
	r := csv.NewReader(file)
	r.Comment = rune('#')

	for {
		line, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v,%v,%v,%v\n", line[0], line[1], line[2], line[3])

		probe := Probe{
			line[0],
			line[1],
			line[2],
			line[3],
		}

		log.Println(probe)
	}

}
