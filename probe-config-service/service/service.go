package service

import (
	"git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/probe-config-service/domain"
)

//InMemoryProbeStore - using a config json file.

type service interface {
	GetProbe(name string) domain.Probe
	GetAll() []domain.Probe
}

type ProbeService struct {
	Config *domain.Config
}

func (i *ProbeService) GetProbe(name string) domain.Probe {
	return i.Config.Probes[name]
}
func (i *ProbeService) GetAll() []domain.Probe {
	var probes []domain.Probe
	for _, probe := range i.Config.Probes {
		probes = append(probes, probe)
	}
	return probes
}
