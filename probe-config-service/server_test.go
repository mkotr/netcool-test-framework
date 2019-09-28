package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubProbeStore struct {
	probes map[string]Probe
}

func (s *StubProbeStore) GetProbe(name string) Probe {
	probe := s.probes[name]
	return probe
}
func (s *StubProbeStore) GetAll() []Probe {
	var probes []Probe
	for _, probe := range s.probes {
		probes = append(probes, probe)
	}

	return probes
}
func TestGETProbes(t *testing.T) {
	store := StubProbeStore{
		map[string]Probe{
			"OMI": Probe{
				"OMI",
				"xxx",
				"lxapp6662.dc.corp.telstra.com",
				"4000",
			},
			"MessageBus": Probe{
				"MessageBus",
				"xxx",
				"lxapp6662.dc.corp.telstra.com",
				"4000",
			},
		},
	}

	server := &ProbeConfigServer{&store}

	t.Run("returns OMI probe details", func(t *testing.T) {
		request := newGetProbeRequest("OMI")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := Probe{
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
		want := Probe{
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

		var got []Probe
		json.NewDecoder(response.Body).Decode(&got)

		want := []Probe{
			Probe{
				"OMI",
				"xxx",
				"lxapp6662.dc.corp.telstra.com",
				"4000",
			},
			Probe{
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
		map[string]Probe{},
	}

	server := &ProbeConfigServer{&store}
	t.Run("returns 404  when getAll and there are no probes", func(t *testing.T) {
		request := newGetProbeRequest("")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Probe
		json.NewDecoder(response.Body).Decode(&got)

		//want := []Probe{}

		assertResponseCode(t, response.Code, http.StatusNotFound)

	})
}

func newGetProbeRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/probes/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, responseBody *bytes.Buffer, want Probe) {
	var got Probe
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
