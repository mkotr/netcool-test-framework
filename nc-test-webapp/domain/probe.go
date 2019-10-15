package domain

//Probe struct
type Probe struct{
	Name string `json:"name"`
	Desc string `json:"desc"`
	Hostname string `json:"hostname"`
	Port string `json:"port"`
}

