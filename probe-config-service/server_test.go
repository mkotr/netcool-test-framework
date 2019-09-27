package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETProbes(t *testing.T) {
	t.Run("returns OMI probe details", func(t *testing.T) {
		request := newGetProbeRequest("OMI")
		response := httptest.NewRecorder()

		ProbeConfigServer(response, request)

		want := Probe{
			"OMI",
			"xxx",
			"lxapp6662.dc.corp.telstra.com",
			"4000",
		}
		assertResponseBody(t, response.Body, want)
	})

	t.Run("returns MessageBus probe details", func(t *testing.T) {
		request := newGetProbeRequest("MessageBus")
		response := httptest.NewRecorder()

		ProbeConfigServer(response, request)
		want := Probe{
			"MessageBus",
			"xxx",
			"lxapp6662.dc.corp.telstra.com",
			"4000",
		}

		assertResponseBody(t, response.Body, want)
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
