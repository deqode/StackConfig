package runtimeCheck

import (
	"testing"

	"github.com/deqodelabs/IaaC/stackconfig/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RuntimeTestSuite struct {
	suite.Suite
	ServiceConfig *pb.ServiceConfig
}

func (suite *RuntimeTestSuite) SetupTest() {
	entrypoint := pb.Entrypoint{
		CommandName: "node index.js",
	}
	nodeConfig := pb.NodeConfig{
		Entrypoint: &entrypoint,
	}
	runtimeConfig := pb.RuntimeConfig{
		RuntimeConfigs: &pb.RuntimeConfig_NodeConfig{NodeConfig: &nodeConfig},
	}
	service := pb.ServiceConfig{
		Version:       1,
		ServiceId:     "efga service",
		Runtime:       pb.Runtime_nodejs14,
		RuntimeConfig: &runtimeConfig,
	}
	suite.ServiceConfig = &service
}

func (suite *RuntimeTestSuite) TestRuntimeNoError() {
	err := CustomValidateRuntime(suite.ServiceConfig)
	assert.Equal(suite.T(), err, nil)
}

func (suite *RuntimeTestSuite) TestRuntimeError() {
	suite.ServiceConfig.Runtime = pb.Runtime_ruby25
	err := CustomValidateRuntime(suite.ServiceConfig)
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "runtime config not available")
	}
}

func TestRuntimeTestSuite(t *testing.T) {
	suite.Run(t, new(RuntimeTestSuite))
}
