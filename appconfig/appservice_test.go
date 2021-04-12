package appconfig

import (
	"testing"

	"github.com/deqodelabs/IaaC/appconfig/pb"
	"github.com/philippgille/gokv/leveldb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type AppServiceTestSuite struct {
	suite.Suite
	appConfig  *pb.AppConfig
	appService AppService
}

func (suite *AppServiceTestSuite) SetupTest() {
	// setup app service
	options := leveldb.DefaultOptions
	store, err := leveldb.NewStore(options)
	if err != nil {
		panic(err)
	}
	logger := zap.NewExample()
	appService := AppService{
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
	app := pb.AppConfig{
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
	suite.appService.Store.Delete(suite.appConfig.AppId)
	defer suite.appService.Store.Close()
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(AppServiceTestSuite))
}

//func TestValidateAppConfig(t *testing.T){
//	options := leveldb.DefaultOptions
//	store,err := leveldb.NewStore(options)
//	if err != nil{
//		panic(err)
//	}
//	defer store.Close()
//
//	logger := zap.NewExample()
//
//	appService := AppService{
//		Store: store,
//		Logger: logger,
//	}
//
//	entrypoint := pb.Entrypoint{
//		CommandName: "node index.js",
//	}
//
//	nodeConfig := pb.NodeConfig{
//		Entrypoint: &entrypoint,
//	}
//
//	runtimeConfig := pb.RuntimeConfig{
//		RuntimeConfigs: &pb.RuntimeConfig_NodeConfig{NodeConfig: &nodeConfig},
//	}
//
//	service := pb.ServiceConfig{
//		Version: 1,
//		ServiceId: "efga service",
//		Runtime: pb.Runtime_nodejs14,
//		RuntimeConfig: &runtimeConfig,
//	}
//	var services []*pb.ServiceConfig
//	services = append(services, &service)
//
//	app := pb.AppConfig{
//		AppId: "efga",
//		AppName: "app1",
//		Version: 1,
//		Services: services,
//	}
//
//	err = appService.ValidateAppConfig(&app)
//	if err != nil{
//		panic(err)
//	}else{
//		logger.Info("Success")
//	}
//
//	forwardedPort := pb.ForwardedPort{
//		HostPort: 10399,
//		ContainerPort: 10501,
//	}
//
//	var forwardedPorts []*pb.ForwardedPort
//	forwardedPorts = append(forwardedPorts, &forwardedPort)
//
//	network := pb.Network{
//		ForwardedPorts: forwardedPorts,
//	}
//
//	service.Network = &network
//	services = append(services, &service)
//	app.Services = services
//
//	err = appService.ValidateAppConfig(&app)
//	if err != nil{
//		panic(err)
//	}else{
//		logger.Info("Success")
//	}
//}
