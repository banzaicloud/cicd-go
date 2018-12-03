package secret

import (
	"context"

	"github.com/banzaicloud/cicd-go/cicd"
)

// V1 is version 1 of the secrets API
const V1 = "application/vnd.cicd.secret.v1+json"

type (
	// Request defines a secret request.
	Request struct {
		Name  string     `json:"name"`
		Repo  cicd.Repo  `json:"repo,omitempty"`
		Build cicd.Build `json:"build,omitempty"`
	}

	// Plugin responds to a secret request.
	Plugin interface {
		Find(context.Context, *Request) (*cicd.Secret, error)
	}
)
