package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/projectcalico/libcalico-go/lib/api"
)

func (h *Handlers) listNodes(w http.ResponseWriter, r *http.Request) {
	nodeMetadata := api.NodeMetadata{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &nodeMetadata); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	nHandler := h.client.Nodes()
	nodes, err := nHandler.List(nodeMetadata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(nodes.Items); err != nil {
		panic(err)
	}
}
