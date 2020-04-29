package simple

import (
	"context"

	"github.com/anz-bank/pkg/log"
	"github.com/anz-bank/sysl-go/config"
	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/handlerinitialiser"
	"github.com/prometheus/client_golang/prometheus"
)

// GENERATED: SERVER

type Server struct {
	params           *core.ServerParams
	serviceInterface *ServiceInterface
	restGenCallback  core.RestGenCallback
	grpcGenCallback  core.GrpcGenCallback
}

func StartServer(file string, customConfig interface{}, build func(*config.DefaultConfig) (*Server, error)) error {
	cfg := NewDefaultConfig()
	err := config.LoadConfig(file, &cfg, &customConfig)
	if err != nil {
		return err
	}

	server, err := build(&cfg)
	if err != nil {
		return err
	}

	return server.Start()
}

func NewServer(ctx context.Context, config *config.DefaultConfig) *Server {
	return &Server{params: &core.ServerParams{Ctx: ctx, Name: "Foo", Config: config}}
}

func (s *Server) WithRest(serviceInterface *ServiceInterface, callback core.RestGenCallback) *Server {
	s.serviceInterface = serviceInterface
	s.restGenCallback = callback
	return s
}

func (s *Server) WithGrpc(serviceInterface *ServiceInterface, callback core.GrpcGenCallback) *Server {
	s.serviceInterface = serviceInterface
	s.grpcGenCallback = callback
	return s
}

func (s *Server) WithPrometheusRegistry(registry *prometheus.Registry) *Server {
	s.params.PrometheusRegistry = registry
	return s
}

func (s *Server) WithLogHook(hook log.Hook) *Server {
	s.params.LogHook = hook
	return s
}

func (s *Server) Start() error {
	clients, err := BuildDownstreamClients(s.params.Config)
	if err != nil {
		return err
	}

	if s.restGenCallback != nil {
		h := []handlerinitialiser.RestHandlerInitialiser{BuildRestHandlerInitialiser(*s.serviceInterface, s.restGenCallback, clients)}
		s.params.RestManager = core.NewRestManager(s.params.Config, h)
	}

	if s.grpcGenCallback != nil {
		h := []handlerinitialiser.GrpcHandlerInitialiser{BuildGrpcHandlerInitialiser(*s.serviceInterface, s.grpcGenCallback, clients)}
		s.params.GrpcManager = core.NewGrpcManager(s.params.Config, h)
	}

	return s.params.Start()
}

// GENERATED: NEW DEFAULT CONFIG

func NewDefaultConfig() config.DefaultConfig {
	return config.DefaultConfig{
		Library: config.LibraryConfig{},
		GenCode: config.GenCodeConfig{
			Downstream: &DownstreamConfig{}, // generated type
		},
	}
}
