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

type ProbeStore interface {
	GetProbe(name string) Probe
	GetAll() []Probe
}

type ProbeConfigServer struct {
	store ProbeStore
}

func (p *ProbeConfigServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/probes/"):]

	if name != "" {
		probe := p.store.GetProbe(name)
		if (probe == Probe{}) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(probe)
		}
	} else {
		probes := p.store.GetAll()
		if len(probes) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		json.NewEncoder(w).Encode(probes)
	}
}

type InMemoryProbeStore struct {
	file   os.File
	probes map[string]Probe
}

func (i *InMemoryProbeStore) GetProbe(name string) Probe {
	return i.probes[name]
}
func (i *InMemoryProbeStore) GetAll() []Probe {
	var probes []Probe
	for _, probe := range i.probes {
		probes = append(probes, probe)
	}
	return probes
}

func main() {
	//read Config file
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatalf("Error reading config file %v", err)
	}
	defer file.Close()

	probes := readConfig(file)

	store := InMemoryProbeStore{
		file:   *file,
		probes: probes,
	}

	server := &ProbeConfigServer{&store}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 5000 %v", err)
	}
}

func readConfig(file *os.File) map[string]Probe {
	r := csv.NewReader(file)
	r.Comment = rune('#')

	probes := make(map[string]Probe)

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

		probes[probe.Name] = probe

		log.Printf("read probe config: %v", probe)
	}
	return probes
}
