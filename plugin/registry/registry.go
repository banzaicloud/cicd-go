package registry

import (
	"context"

	"github.com/banzaicloud/cicd-go/cicd"
)

// V1 is version 1 of the registry API
const V1 = "application/vnd.cicd.registry.v1+json"

type (
	// Request defines a registry request.
	Request struct {
		Repo  cicd.Repo  `json:"repo,omitempty"`
		Build cicd.Build `json:"build,omitempty"`
	}

	// Plugin responds to a registry request.
	Plugin interface {
		List(context.Context, *Request) ([]*cicd.Registry, error)
	}
)
