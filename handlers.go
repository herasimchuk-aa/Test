package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/projectcalico/libcalico-go/lib/client"
)

// Handlers implements the API handler to work with REST APIs.
type Handlers struct {
	client *client.Client
}

// NewCalicoAPIHandlers returns a Handlers implementation.
func NewCalicoAPIHandlers() *Handlers {
	var err error

	cfg, err := client.LoadClientConfig("")
	if err != nil {
		log.Errorf("Failed to load client config: %q", err)
		os.Exit(1)
	}

	c, err := client.New(*cfg)
	if err != nil {
		log.Errorf("Failed creating client: %q", err)
		os.Exit(1)
	}
	log.Infof("Client: %v", c)

	h := &Handlers{
		client: c,
	}

	return h
}
