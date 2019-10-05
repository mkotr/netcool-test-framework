package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/probe-config-service/domain"
	"git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/probe-config-service/service"
)

type StubProbeStore struct {
	probes map[string]domain.Probe
}

func (s *StubProbeStore) GetProbe(name string) domain.Probe {
	probe := s.probes[name]
	return probe
}
func (s *StubProbeStore) GetAll() []domain.Probe {
	var probes []domain.Probe
	for _, probe := range s.probes {
		probes = append(probes, probe)
	}

	return probes
}
func TestGETProbes(t *testing.T) {
	store := StubProbeStore{
		map[string]domain.Probe{
			"OMI": domain.Probe{
				"OMI",
				"xxx",
				"lxapp6662.dc.corp.telstra.com",
				"4000",
			},
			"MessageBus": domain.Probe{
				"MessageBus",
				"xxx",
				"lxapp6662.dc.corp.telstra.com",
				"4000",
			},
		},
	}

	server := NewProbeConfigServer(&store)

	t.Run("returns OMI probe details", func(t *testing.T) {
		request := newGetProbeRequest("OMI")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := domain.Probe{
			"OMI",
			"xxx",
			"lxapp6662.dc.corp.telstra.com",
			"4000",
		}
		assertResponseBody(t, response.Body, want)
		assertResponseCode(t, response.Code, http.StatusOK)

	})

	t.Run("returns MessageBus probe details", func(t *testing.T) {
		request := newGetProbeRequest("MessageBus")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		want := domain.Probe{
			"MessageBus",
			"xxx",
			"lxapp6662.dc.corp.telstra.com",
			"4000",
		}

		assertResponseBody(t, response.Body, want)
		assertResponseCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing probes", func(t *testing.T) {
		request := newGetProbeRequest("Nothing")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		assertResponseCode(t, got, want)
	})

	t.Run("returns 200 and all probes", func(t *testing.T) {
		request := newGetProbeRequest("")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []domain.Probe
		json.NewDecoder(response.Body).Decode(&got)

		want := []domain.Probe{
			domain.Probe{
				"OMI",
				"xxx",
				"lxapp6662.dc.corp.telstra.com",
				"4000",
			},
			domain.Probe{
				"MessageBus",
				"xxx",
				"lxapp6662.dc.corp.telstra.com",
				"4000",
			},
		}

		assertResponseCode(t, response.Code, http.StatusOK)

		if len(got) != 2 {
			t.Errorf("got %v, want %v", len(got), len(want))
		}

	})

}

func TestGetAll(t *testing.T) {
	store := StubProbeStore{
		map[string]domain.Probe{},
	}

	server := &ProbeConfigServer{&store}
	t.Run("returns 404  when getAll and there are no probes", func(t *testing.T) {
		request := newGetProbeRequest("")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []domain.Probe
		json.NewDecoder(response.Body).Decode(&got)

		//want := []domain.Probe{}

		assertResponseCode(t, response.Code, http.StatusNotFound)

	})
}

func newGetProbeRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/probes/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, responseBody *bytes.Buffer, want domain.Probe) {
	var got domain.Probe
	json.NewDecoder(responseBody).Decode(&got)
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func assertResponseCode(t *testing.T, got, want int) {
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
