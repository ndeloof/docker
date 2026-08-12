// Package compose is a demo in-process extension exposing the
// "compose stack" point on the daemon socket. The provider is a no-op: it
// receives the resolved Compose project definition pushed by the Compose
// client at the start of "up" and only logs it, to validate the model.
package compose

import (
	"context"

	"github.com/containerd/log"
	stackv0 "github.com/moby/moby/v2/extpoints/compose/stack/v0"
	stackpb "github.com/moby/moby/v2/extpoints/compose/stack/v0/protogen"
	servicegrpcv0 "github.com/moby/moby/v2/extpoints/servicegrpc/v0"
	"github.com/moby/moby/v2/internal/extensions"
	"google.golang.org/grpc"
)

// ID is the extension id.
const ID = "com.docker.compose.v1"

type stack struct{}

func (stack) Publish(ctx context.Context, req *stackv0.PublishRequest) (*stackv0.PublishReply, error) {
	log.G(ctx).WithFields(log.Fields{
		"project":  req.Project,
		"complete": req.Complete,
		"bytes":    len(req.ComposeYaml),
	}).Info("compose extension: received compose stack definition")
	return &stackv0.PublishReply{}, nil
}

// expose registers the compose stack gRPC service for socket exposure.
type expose struct{}

func (expose) RegisterServices(r grpc.ServiceRegistrar) {
	stackpb.ServerPoint.Register(r, stack{})
}

// Extension implements the service.grpc point with a no-op compose stack
// receiver.
var Extension = extensions.New(extensions.Declaration{
	ID:        ID,
	Providers: []extensions.Provider{servicegrpcv0.Point.Provide(expose{})},
})
