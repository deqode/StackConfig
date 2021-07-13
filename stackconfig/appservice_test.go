package stackconfig

import (
	"testing"

	"github.com/deqodelabs/IaaC/stackconfig/pb"
	"github.com/philippgille/gokv/leveldb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type AppServiceTestSuite struct {
	suite.Suite
	appConfig  *pb.StackConfig
	appService StackService
}

func (suite *AppServiceTestSuite) SetupTest() {
	// setup app service
	options := leveldb.DefaultOptions
	store, err := leveldb.NewStore(options)
	if err != nil {
		panic(err)
	}
	logger := zap.NewExample()
	appService := StackService{
		Store:  store,
		Logger: logger,
	}
	suite.appService = appService

	//setup app config
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
	var services []*pb.ServiceConfig
	services = append(services, &service)
	app := pb.StackConfig{
		AppName:  "app1",
		Version:  1,
		Services: services,
	}
	suite.appConfig = &app
}

func (suite *AppServiceTestSuite) TestSave() {
	assert.Equal(suite.T(), suite.appConfig.AppId, "")
	config, err := suite.appService.Save(suite.appConfig)
	assert.Equal(suite.T(), err, nil)
	assert.IsType(suite.T(), "abc", config.AppId)
	config, err = suite.appService.GetAppConfig(config.AppId)
	assert.Equal(suite.T(), err, nil)
	config.Version = 2
	config, err = suite.appService.Save(config)
	assert.Equal(suite.T(), err, nil)
	config, err = suite.appService.GetAppConfig(config.AppId)
	assert.Equal(suite.T(), err, nil)
	assert.Equal(suite.T(), int32(2), config.Version)
	config, err = suite.appService.Save(config)
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "version already exist")
	}
	config.Version = 4
	_, err = suite.appService.Save(config)
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "version missed")
	}
	suite.appService.Store.Delete(suite.appConfig.AppId)
	defer suite.appService.Store.Close()
}

func (suite *AppServiceTestSuite) TestGetAppConfig() {
	_, err := suite.appService.GetAppConfig("abc")
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "id for app config not found")
	}
	config, err := suite.appService.Save(suite.appConfig)
	assert.Equal(suite.T(), err, nil)
	assert.IsType(suite.T(), "abc", config.AppId)
	_, err = suite.appService.GetAppConfig(config.AppId)
	assert.Equal(suite.T(), err, nil)
	defer suite.appService.Store.Close()
}

func (suite *AppServiceTestSuite) TestGetAppConfigForVersion() {
	_, err := suite.appService.GetAppConfig("abc")
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "id for app config not found")
	}
	config, err := suite.appService.Save(suite.appConfig)
	assert.Equal(suite.T(), err, nil)
	assert.IsType(suite.T(), "abc", config.AppId)
	_, err = suite.appService.GetAppConfigForVersion(config.AppId, 1)
	assert.Equal(suite.T(), err, nil)
	_, err = suite.appService.GetAppConfigForVersion(config.AppId, 2)
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "version does not exist")
	}
	_, err = suite.appService.GetAppConfigForVersion(config.AppId, 0)
	if assert.Error(suite.T(), err) {
		assert.EqualError(suite.T(), err, "version is deprecated")
	}
	defer suite.appService.Store.Close()
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(AppServiceTestSuite))
}
