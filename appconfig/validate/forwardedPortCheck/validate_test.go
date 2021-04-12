package forwardedPortCheck

import (
	"testing"

	"github.com/deqodelabs/IaaC/appconfig/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ForwardedPortTestSuite struct {
	suite.Suite
	forwardedPorts []*pb.ForwardedPort
}

func (suite *ForwardedPortTestSuite) SetupTest() {
	forwardedPort := pb.ForwardedPort{
		HostPort:      10399,
		ContainerPort: 10501,
	}
	suite.forwardedPorts = suite.forwardedPorts[:0]
	suite.forwardedPorts = append(suite.forwardedPorts, &forwardedPort)
}

func (suite *ForwardedPortTestSuite) TestForwardedPortsNoError() {
	err := CustomForwardedPortValidation(suite.forwardedPorts)
	assert.Equal(suite.T(), err, nil)
}

func (suite *ForwardedPortTestSuite) TestForwardedHostPort() {
	suite.forwardedPorts[0].HostPort = 10400
	err := CustomForwardedPortValidation(suite.forwardedPorts)
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "host port should not be in range [10400, 10500]")
	}
}

func (suite *ForwardedPortTestSuite) TestForwardedContainerPort() {
	suite.forwardedPorts[0].ContainerPort = 10500
	err := CustomForwardedPortValidation(suite.forwardedPorts)
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "container port should not be in range [10400, 10500]")
	}
}

func TestForwardedPortTestSuite(t *testing.T) {
	suite.Run(t, new(ForwardedPortTestSuite))
}
