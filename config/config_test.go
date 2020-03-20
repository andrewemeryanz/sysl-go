package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type TestMyConfig struct {
	Server TestServer `yaml:"server"`
}

type TestServer struct {
	AdminServer TestAdminServer `yaml:"adminServer"`
}

type TestAdminServer struct {
	ContextTimeout time.Duration `yaml:"contextTimeout"`
	Http           TestHttp      `yaml:"http"`
}

type TestHttp struct {
	BasePath     string `yaml:"basePath"`
	ReadTimeout  string `yaml:"readTimeout"`
	WriteTimeout string `yaml:"writeTimeout"`
}

type TestDownstreamConfig struct {
	ContextTimeout time.Duration        `yaml:"contextTimeout"`
	Foo            CommonDownstreamData `yaml:"foo"`
	Bar            CommonDownstreamData `yaml:"bar"`
}

func TestSReadConfig(t *testing.T) {
	t.Parallel()

	lib := LibraryConfig{}
	gen := GenCodeConfig{}
	gen.Downstream = &TestDownstreamConfig{}
	myConfig := TestMyConfig{}
	ReadConfig("testdata/config.yaml", &lib, &gen, &myConfig)

	require.Equal(t, time.Duration(2*time.Second), myConfig.Server.AdminServer.ContextTimeout)
	require.Equal(t, "/admintest", myConfig.Server.AdminServer.Http.BasePath)

	require.False(t, lib.Log.ReportCaller)

	require.Equal(t, 8080, gen.Upstream.HTTP.Common.Port)
	require.Equal(t, 8081, gen.Upstream.GRPC.Port)
	require.Equal(t, time.Duration(120*time.Second), gen.Downstream.(*TestDownstreamConfig).Foo.ClientTimeout)
	require.Equal(t, "https://bar.example.com", gen.Downstream.(*TestDownstreamConfig).Bar.ServiceURL)
}
