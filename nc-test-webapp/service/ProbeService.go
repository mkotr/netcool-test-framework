package service

import (
	"net/http"
	"git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/nc-test-webapp/domain"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type service interface {
	GetAll() ([]domain.Probe, error)
}

//ProbeService struct TODO add client?
type ProbeService struct {}

//GetAll the probes from the config service
func (p *ProbeService) GetAll() ([]domain.Probe, error){
	//configurise this into the conf
	url := "http://localhost:5000/probes/"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var probes []domain.Probe

	err = json.Unmarshal(body, &probes)

	if err != nil {
		return nil, err
	}

	fmt.Println(probes)
	return probes, nil
}