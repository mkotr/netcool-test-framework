package domain

type Probe struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

type Config struct {
	Probes map[string]Probe `json:"probes"`
}
