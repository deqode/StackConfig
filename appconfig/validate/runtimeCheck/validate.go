package runtimeCheck

import(
	"errors"
	"github.com/deqodelabs/IaaC/appconfig/pb"
)

func CustomValidateRuntime(service *pb.ServiceConfig) error {
	switch  service.Runtime {
		case pb.Runtime_python27:
			if service.RuntimeConfig.GetPython27Config() == nil{
				return errors.New("runtime config not available")
			}
		case pb.Runtime_nodejs10:
			if service.RuntimeConfig.GetNodeConfig() == nil{
				return errors.New("runtime config not available")
			}
		case pb.Runtime_nodejs12:
			if service.RuntimeConfig.GetNodeConfig() == nil{
				return errors.New("runtime config not available")
			}
		case pb.Runtime_nodejs14:
			if service.RuntimeConfig.GetNodeConfig() == nil{
				return errors.New("runtime config not available")
			}
		case pb.Runtime_ruby25:
			if service.RuntimeConfig.GetRubyConfig() == nil{
				return errors.New("runtime config not available")
			}
		case pb.Runtime_ruby26:
			if service.RuntimeConfig.GetRubyConfig() == nil{
				return errors.New("runtime config not available")
			}
		case pb.Runtime_ruby27:
			if service.RuntimeConfig.GetRubyConfig() == nil{
				return errors.New("runtime config not available")
			}
	}
	return nil
}
