package appconfig

import (
	"github.com/deqodelabs/IaaC/appconfig/pb"
)

type App interface{
	ValidateAppConfig(app *pb.AppConfig) error
	Save(app *pb.AppConfig) (*pb.AppConfig, error)
	GetAppConfig(id string) (*pb.AppConfig, error)
	GetAppConfigForVersion(id string, version int32) (*pb.AppConfig, error)
}