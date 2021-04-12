package resourceCheck

import (
	"testing"

	"github.com/deqodelabs/IaaC/appconfig/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResourceTestSuite struct {
	suite.Suite
	resource *pb.Resource
}

func (suite *ResourceTestSuite) SetupTest() {
	suite.resource = &pb.Resource{
		Cpu: 1,
	}
}

func (suite *ResourceTestSuite) TestCpuNoError1() {
	err := CustomResourceValidation(suite.resource)
	assert.Equal(suite.T(), err, nil)
}

func (suite *ResourceTestSuite) TestCpuNoError2() {
	suite.resource.Cpu = 2
	err := CustomResourceValidation(suite.resource)
	assert.Equal(suite.T(), err, nil)
}

func (suite *ResourceTestSuite) TestCpuError() {
	suite.resource.Cpu = 3
	err := CustomResourceValidation(suite.resource)
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "cpu number not valid")
	}
}

func TestResourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceTestSuite))
}
