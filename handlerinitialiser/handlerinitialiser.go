package handlerinitialiser

import (
	"context"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"
)

type HandlerInitialiser interface {
	Name() string                // Human-friendly name of the service
	Config() interface{} // Reference to config for this service.
}

type RestHandlerInitialiser interface {
	HandlerInitialiser
	WireRoutes(ctx context.Context, r chi.Router)
}

type GrpcHandlerInitialiser interface {
	HandlerInitialiser
	RegisterServer(ctx context.Context, server *grpc.Server)
}
