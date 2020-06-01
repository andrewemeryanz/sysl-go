package core

// MARKED TO IGNORE COVERAGE

import (
	"context"
	"fmt"
	"github.com/anz-bank/pkg/log"
	"github.com/anz-bank/sysl-go/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type ServerParams struct {
	Ctx                context.Context
	Name               string
	logrusLogger       *logrus.Logger
	pkgLoggerConfigs   []log.Config
	restManager        Manager
	grpcManager        GrpcManager
	prometheusRegistry *prometheus.Registry
}

//nolint:gocognit // Long method names are okay because only generated code will call this, not humans.
func NewServerParams(ctx context.Context, name string, opts ...ServerOption) *ServerParams {
	params := &ServerParams{Ctx: ctx, Name: name}
	for _, o := range opts {
		o.apply(params)
	}
	return params
}

//nolint:gocognit // Long method are okay because only generated code will call this, not humans.
func (params *ServerParams) Start() error {
	ctx := params.Ctx

	// initialise the logger
	// sysl-go always uses a pkg logger internally. if custom code passes in a logrus logger, a
	// mechanism which is deprecated, then a hook is added to the internal pkg logger that forwards
	// logged events to the provided logrus logger.
	// sysl-go can be requested to log in a verbose manner. logger in a verbose manner logs additional
	// details within log events where appropriate. the mechanism to set this verbose manner is to
	// either have a sufficiently high logrus log level or the verbose mode set against the pkg logger.
	configs := params.pkgLoggerConfigs
	verboseLogging := false
	if params.logrusLogger != nil {
		configs = append(configs, log.AddHooks(&logrusHook{params.logrusLogger}))
		ctx = common.LoggerToContext(ctx, params.logrusLogger, nil)
		verboseLogging = params.logrusLogger.Level >= logrus.DebugLevel
	}
	ctx = log.WithConfigs(configs...).Onto(ctx)
	verboseMode := log.SetVerboseMode(true)
	for _, config := range configs {
		if config == verboseMode { // TODO: test this logic
			verboseLogging = true
			break
		}
	}

	// TODO: set the 'verbose logging' property in the context

	// prepare the middleware
	mWare := prepareMiddleware(ctx, params.Name, params.prometheusRegistry)

	var restIsRunning, grpcIsRunning bool

	// Run the REST server
	var listenAdmin func() error
	if params.restManager != nil && params.restManager.AdminServerConfig() != nil {
		var err error
		listenAdmin, err = configureAdminServerListener(ctx, params.restManager, params.prometheusRegistry, mWare.admin)
		if err != nil {
			return err
		}
	} else {
		// set up a dummy listener which will never exit if admin disabled
		listenAdmin = func() error { select {} }
	}

	var listenPublic func() error
	if params.restManager != nil && params.restManager.PublicServerConfig() != nil {
		var err error
		listenPublic, err = configurePublicServerListener(ctx, params.restManager, mWare.public)
		if err != nil {
			return err
		}
		restIsRunning = true
	} else {
		listenPublic = func() error { select {} }
	}

	// Run the gRPC server
	var listenPublicGrpc func() error
	if params.grpcManager != nil && params.grpcManager.GrpcPublicServerConfig() != nil {
		var err error
		listenPublicGrpc, err = configurePublicGrpcServerListener(ctx, params.grpcManager)
		if err != nil {
			return err
		}

		grpcIsRunning = true
	} else {
		listenPublicGrpc = func() error { select {} }
	}

	// Panic if REST&gRPC are not running
	if !restIsRunning && !grpcIsRunning {
		panic("Both servers are set to nil")
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- listenPublic()
	}()
	go func() {
		errChan <- listenAdmin()
	}()
	go func() {
		errChan <- listenPublicGrpc()
	}()

	return <-errChan
}

type ServerOption interface {
	apply(params *ServerParams)
}

type restManagerOption struct {
	restManager Manager
}

func (o *restManagerOption) apply(params *ServerParams) {
	params.restManager = o.restManager
}

func WithRestManager(manager Manager) ServerOption {
	return &restManagerOption{manager}
}

type logrusLoggerOption struct {
	logger *logrus.Logger
}

func (o *logrusLoggerOption) apply(params *ServerParams) {
	params.logrusLogger = o.logger
}

// Deprecated: Use WithPkgLogger instead
func WithLogrusLogger(logger *logrus.Logger) ServerOption {
	return &logrusLoggerOption{logger}
}

type pkgLoggerOption struct {
	configs []log.Config
}

func (o *pkgLoggerOption) apply(params *ServerParams) {
	params.pkgLoggerConfigs = o.configs
}

func WithPkgLogger(configs ...log.Config) ServerOption {
	return &pkgLoggerOption{configs}
}

func WithPrometheusRegistry(prometheusRegistry *prometheus.Registry) ServerOption {
	return &prometheusRegistryOption{prometheusRegistry}
}

type prometheusRegistryOption struct {
	prometheusRegistry *prometheus.Registry
}

func (o *prometheusRegistryOption) apply(params *ServerParams) {
	params.prometheusRegistry = o.prometheusRegistry
}

type grpcManagerOption struct {
	grpcManager GrpcManager
}

func (o *grpcManagerOption) apply(params *ServerParams) {
	params.grpcManager = o.grpcManager
}

func WithGrpcManager(manager GrpcManager) ServerOption {
	return &grpcManagerOption{manager}
}

// Deprecated: Use ServerParams instead
//nolint:gocognit // Long method names are okay because only generated code will call this, not humans.
func Server(ctx context.Context, name string, hl Manager, grpcHl GrpcManager, logger *logrus.Logger, promRegistry *prometheus.Registry) error {
	return NewServerParams(ctx, name,
		WithPkgLogger(),
		WithLogrusLogger(logger),
		WithRestManager(hl),
		WithGrpcManager(grpcHl),
		WithPrometheusRegistry(promRegistry)).Start()
}

// TODO: this (and other logging stuff) should probably be pulled out into another file
type logrusHook struct {
	logger *logrus.Logger
}

func (h *logrusHook) OnLogged(entry *log.LogEntry) error {
	e := pkgLogEntryToLogrusEntry(h.logger, entry)
	e.Log(e.Level)
	return nil
}

// Convert the given pkg entry to a logrus entry.
func pkgLogEntryToLogrusEntry(logger *logrus.Logger, entry *log.LogEntry) *logrus.Entry {
	return &logrus.Entry{
		Logger:  logger,
		Data:    pkgLogEntryToLogrusFields(entry),
		Time:    entry.Time,
		Level:   verboseToLogrusLevel(entry.Verbose),
		Message: entry.Message,
	}
}

// Convert the pkg log entry into appropriate logrus fields to log.
func pkgLogEntryToLogrusFields(entry *log.LogEntry) logrus.Fields {
	fields := make(map[string]interface{})
	iterator := entry.Data.Range()
	for iterator.Next() {
		fields[iterator.Key().(string)] = iterator.Value()
	}
	if entry.Caller.File != "" {
		fields["caller"] = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	}
	return fields
}

// Convert the pkg concept of verbosity to a logrus level.
func verboseToLogrusLevel(verbose bool) logrus.Level {
	if verbose {
		return logrus.DebugLevel
	}
	return logrus.InfoLevel
}
