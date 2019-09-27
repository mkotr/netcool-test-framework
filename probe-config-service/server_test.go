package main

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestGETProbes(t *testing.T) {
	t.Run("returns OMI probe details", func(t *testing.T){
		request, _ := http.NewRequest(http.MethodGet, "/probes/OMI", nil)
		response := httptest.NewRecorder()

		ProbeConfigServer(response, request)

		var got Probe
		json.NewDecoder(response.Body).Decode(&got)
		want := Probe{
			"OMI",
			"xxx",
			"lxapp6662.dc.corp.telstra.com",
			"4000",
		}

		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("returns MessageBus probe details", func(t *testing.T){
		request, _ := http.NewRequest(http.MethodGet, "/probes/OMI", nil)
		response := httptest.NewRecorder()

		ProbeConfigServer(response, request)

		var got Probe
		json.NewDecoder(response.Body).Decode(&got)
		want := Probe{
			"MessageBus",
			"xxx",
			"lxapp6662.dc.corp.telstra.com",
			"4000",
		}

		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}