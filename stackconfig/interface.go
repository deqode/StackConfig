package stackconfig

import (
	"github.com/deqodelabs/IaaC/stackconfig/pb"
)

type Stack interface {
	ValidateAppConfig(app *pb.StackConfig) error
	Save(app *pb.StackConfig) (*pb.StackConfig, error)
	GetAppConfig(id string) (*pb.StackConfig, error)
	GetAppConfigForVersion(id string, version int32) (*pb.StackConfig, error)
	//Delete()
	//Deprecate()
}
