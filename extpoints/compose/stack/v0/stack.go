//go:generate go run github.com/moby/moby/v2/internal/extensions/cmd/mobyextgen

// Package stackv0 defines the "compose stack" extension point: a service that
// receives the resolved Compose project definition (the equivalent of the
// fully-resolved compose.yaml) from the Compose client before "up" starts, so
// that an engine-side component can pre-populate caches and service maps.
package stackv0

import (
	"context"

	"github.com/moby/moby/v2/internal/extensions"
)

// Stack receives Compose project definitions ahead of deployment.
type Stack interface {
	// Publish delivers the resolved Compose project definition. It is called
	// by the Compose client at the start of "compose up", before any other
	// engine API call.
	Publish(ctx context.Context, req *PublishRequest) (*PublishReply, error)
}

// PublishRequest carries the resolved Compose project.
type PublishRequest struct {
	// Project is the Compose project name.
	Project string `pb:"1"`
	// ComposeYaml is the fully-resolved Compose model, YAML-encoded.
	ComposeYaml []byte `pb:"2"`
	// Complete is true when the payload is the whole project, false when it
	// was narrowed to a subset of services (compose up <service...>).
	Complete bool `pb:"3"`
}

// PublishReply is empty: the publication is fire-and-forget.
type PublishReply struct{}

//mobyextgen:service=Stack
var Point = extensions.DefinePoint[Stack]("com.docker.compose.stack.v0")
