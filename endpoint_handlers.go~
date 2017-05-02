package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/projectcalico/libcalico-go/lib/api"
)

func (h *Handlers) listWorkloadEndpoints(w http.ResponseWriter, r *http.Request) {
	endpointMetadata := api.WorkloadEndpointMetadata{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &endpointMetadata); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	eHandler := h.client.WorkloadEndpoints()
	endpoints, err := eHandler.List(endpointMetadata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(endpoints.Items); err != nil {
		panic(err)
	}
}
