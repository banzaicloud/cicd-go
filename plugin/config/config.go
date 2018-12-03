package config

import (
	"context"

	"github.com/banzaicloud/cicd-go/cicd"
)

// V1 is version 1 of the configuration API
const V1 = "application/vnd.cicd.config.v1+json"

type (
	// Request defines a configuration request.
	Request struct {
		Build cicd.Build `json:"build,omitempty"`
		Repo  cicd.Repo  `json:"repo,omitempty"`
	}

	// Plugin responds to a configuration request.
	Plugin interface {
		Find(context.Context, *Request) (*cicd.Config, error)
	}
)
